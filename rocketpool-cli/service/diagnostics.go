package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/rocket-pool/smartnode/rocketpool-cli/node"
	"github.com/rocket-pool/smartnode/shared/services/config"
	"github.com/rocket-pool/smartnode/shared/services/rocketpool"
	cfgtypes "github.com/rocket-pool/smartnode/shared/types/config"
	cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"
)

type RecommendedVersions struct {
	Rp            string         `json:"rp"`
	RpClients     ClientVersions `json:"rp_clients"`
	RpBeta        string         `json:"rp_beta"`
	RpBetaClients ClientVersions `json:"rp_beta_clients"`
}

type ClientVersions struct {
	Geth       string `json:"geth"`
	Besu       string `json:"besu"`
	Nethermind string `json:"nethermind"`
	Lighthouse string `json:"lighthouse"`
	Nimbus     string `json:"nimbus"`
	Teku       string `json:"Teku"`
	Prysm      string `json:"prysm"`
	Lodestar   string `json:"lodestar"`
}

type DiagnosticsResponse struct {
	Status       string               `json:"status"`
	Error        string               `json:"error"`
	Architecture string               `json:"arch"`
	ECPort       uint16               `json:"ec_port"`
	CCPort       uint16               `json:"cc_port"`
	ExternalIP   string               `json:"ip"`
	IPV6         bool                 `json:"json:ipv6"`
	ECPortOpen   bool                 `json:"ec_port_open"`
	CCPortOpen   bool                 `json:"cc_port_open"`
	FreeDisk     uint64               `json:"free_disk"`
	TotalRAM     uint64               `json:"total_ram"`
	Chronyd      bool                 `json:"chronyd"`
	RecVersions  *RecommendedVersions `json:"recVersions"`
}

func runDiagnostics(c *cli.Context) error {

	// Get RP client
	rp, err := rocketpool.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer rp.Close()

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(rp)
	if err != nil {
		return err
	}

	cfg, _, err := rp.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading old configuration: %w", err)
	}

	// Get the container prefix
	prefix, err := getContainerPrefix(rp)
	if err != nil {
		return fmt.Errorf("Error getting container prefix: %w", err)
	}

	// Get RP service version
	servVersion, err := rp.GetServiceVersion()
	if err != nil {
		return err
	}

	// Response
	response := DiagnosticsResponse{}

	fmt.Print("Running system diagnostics:\n\n")

	node.GetSyncProgress(c)

	checkCpuFeatures()

	// // Get node account
	// nodeAccount, err := w.GetNodeAccount()
	// if err != nil {
	// 	return nil, err
	// }

	// Data
	var wg errgroup.Group

	// Get system arch
	wg.Go(func() error {
		var err error
		response.Architecture = getArchitecture()
		return err
	})

	wg.Go(func() error {
		var err error
		response.TotalRAM = uint64(readMemoryStats().MemTotal)
		return err
	})

	// Check if the EC P2P port is open
	wg.Go(func() error {
		var err error
		port := cfg.ExecutionCommon.P2pPort.Value.(uint16)
		ip, err := getExternalIP()
		if err != nil {
			return err
		}

		response.IPV6 = strings.Contains(ip, ":")
		response.ECPort = port
		response.ECPortOpen = isPortOpen(ip, port)
		return err
	})

	// Check if the CC P2P port is opened
	wg.Go(func() error {
		var err error
		port := cfg.ConsensusCommon.P2pPort.Value.(uint16)

		ip, err := getExternalIP()
		if err != nil {
			return err
		}
		response.CCPort = port
		response.CCPortOpen = isPortOpen(ip, port)
		return err
	})

	// Check free disk space
	wg.Go(func() error {
		var err error
		freeDisk, err := checkDiskSpace(prefix, rp)
		if err != nil {
			return err
		}
		response.FreeDisk = freeDisk
		return err
	})

	wg.Go(func() error {
		var err error
		recommendedVersions, err := fetchRecommendedVersions()
		if err != nil {
			return err
		}
		response.RecVersions = recommendedVersions
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return err
	}

	getClientVersions(c, cfg, servVersion, response.RecVersions)

	// Print diagnostics & return

	fmt.Printf("\nArchitecture: %s\n\n", response.Architecture)

	if response.IPV6 {
		printYellow("Using an external IPv6 address\n")
	} else {
		printGreen("Using an external IPv4 address\n")
	}
	if response.ECPortOpen {
		printGreen(fmt.Sprintf("EC P2P Port %d is open\n", response.ECPort))
	} else {
		printRed(fmt.Sprintf("EC P2P Port %d is closed!\n", response.ECPort))
	}

	if response.CCPortOpen {
		printGreen(fmt.Sprintf("CC P2P Port %d is open\n\n", response.CCPort))
	} else {
		printRed(fmt.Sprintf("CC P2P Port %d is closed!\n\n", response.CCPort))
	}

	if response.FreeDisk < 100*1024*1024*1024 {
		printYellow(fmt.Sprintf("Low free disk space: %s\n\n", humanize.IBytes(response.FreeDisk)))
	} else {
		printGreen(fmt.Sprintf("Free disk space: %s\n\n", humanize.IBytes(response.FreeDisk)))
	}

	if response.TotalRAM >= 31*1024*1024*1024 {
		printGreen(fmt.Sprintf("Total RAM: %s - good for any client combination\n", humanize.IBytes(response.TotalRAM)))
	} else if response.TotalRAM >= 15*1024*1024*1024 {
		printYellow(fmt.Sprintf("Total RAM: %s - a few clients might have issues\n", humanize.IBytes(response.TotalRAM)))
	} else {
		printRed(fmt.Sprintf("Total RAM: %s - is very low, only specific clients are recommended\n", humanize.IBytes(response.TotalRAM)))
	}

	return nil

}

func fetchRecommendedVersions() (*RecommendedVersions, error) {
	//url := "https://raw.githubusercontent.com/rocket-pool/smartnode/master/recommended_versions.json"
	url := "https://raw.githubusercontent.com/0xfornax/smartnode/run-diagnostics/recommended_versions.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	recVersions := RecommendedVersions{}

	err = json.Unmarshal(body, &recVersions)
	if err != nil {
		return nil, err
	}

	return &recVersions, nil
}

type Memory struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
}

func readMemoryStats() Memory {
	res := Memory{0, 0, 0}
	memInfoRaw, err := exec.Command("cat", "/proc/meminfo").Output()
	if err != nil {
		fmt.Errorf(err.Error())
		return res
	}

	memInfo := strings.Split(string(memInfoRaw), "\n")

	for _, info := range memInfo {
		if len(info) > 0 {
			key, value := parseLine(info)
			// Multiplying by 1024 as the results com in kb
			switch key {
			case "MemTotal":
				res.MemTotal = value * 1024
			case "MemFree":
				res.MemFree = value * 1024
			case "MemAvailable":
				res.MemAvailable = value * 1024
			}
		}
	}
	return res
}

func parseLine(raw string) (key string, value int) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
	if raw == "" {
		return 0
	}
	res, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return res
}

func getClientVersions(c *cli.Context, cfg *config.RocketPoolConfig, servVersion string, recVersions *RecommendedVersions) {
	rpVersion := ""
	var cv ClientVersions
	if strings.Contains(c.App.Version, "b") {
		rpVersion = recVersions.RpBeta
		cv = recVersions.RpClients
	} else {
		rpVersion = recVersions.Rp
		cv = recVersions.RpBetaClients
	}
	// Get the execution client string
	var eth1ClientString string
	eth1ClientMode := cfg.ExecutionClientMode.Value.(cfgtypes.Mode)
	switch eth1ClientMode {
	case cfgtypes.Mode_Local:
		eth1Client := cfg.ExecutionClient.Value.(cfgtypes.ExecutionClient)
		format := "%s (Locally managed)\n\tImage: %s Recommended: %s"
		switch eth1Client {
		case cfgtypes.ExecutionClient_Geth:
			eth1ClientString = fmt.Sprintf(format, "Geth", cfg.Geth.ContainerTag.Value.(string), cv.Geth)
		case cfgtypes.ExecutionClient_Nethermind:
			eth1ClientString = fmt.Sprintf(format, "Nethermind", cfg.Nethermind.ContainerTag.Value.(string), cv.Nethermind)
		case cfgtypes.ExecutionClient_Besu:
			eth1ClientString = fmt.Sprintf(format, "Besu", cfg.Besu.ContainerTag.Value.(string), cv.Besu)
		default:
			fmt.Errorf("unknown local execution client [%v]", eth1Client)
		}

	case cfgtypes.Mode_External:
		eth1ClientString = "Externally managed"

	default:
		fmt.Errorf("unknown execution client mode [%v]", eth1ClientMode)
	}

	// Get the consensus client string
	var eth2ClientString string
	var validatorClientString string
	eth2ClientMode := cfg.ConsensusClientMode.Value.(cfgtypes.Mode)
	switch eth2ClientMode {
	case cfgtypes.Mode_Local:
		eth2Client := cfg.ConsensusClient.Value.(cfgtypes.ConsensusClient)
		format := "%s (Locally managed)\n\tImage: %s - Recommended: %s"
		switch eth2Client {
		case cfgtypes.ConsensusClient_Lighthouse:
			eth2ClientString = fmt.Sprintf(format, "Lighthouse", cfg.Lighthouse.ContainerTag.Value.(string), cv.Lighthouse)
		case cfgtypes.ConsensusClient_Nimbus:
			eth2ClientString = fmt.Sprintf(format, "Nimbus", cfg.Nimbus.ContainerTag.Value.(string), cv.Nimbus)
		case cfgtypes.ConsensusClient_Prysm:
			// Prysm is a special case, as the BN and VC image versions may differ
			eth2ClientString = fmt.Sprintf(format, "Prysm", cfg.Prysm.BnContainerTag.Value.(string), cv.Prysm)
			validatorClientString = cfg.Prysm.VcContainerTag.Value.(string)
		case cfgtypes.ConsensusClient_Teku:
			eth2ClientString = fmt.Sprintf(format, "Teku", cfg.Teku.ContainerTag.Value.(string), cv.Teku)
		default:
			fmt.Errorf("unknown local consensus client [%v]", eth2Client)
		}

	case cfgtypes.Mode_External:
		eth2Client := cfg.ExternalConsensusClient.Value.(cfgtypes.ConsensusClient)
		format := "%s (Externally managed)\n\tVC Image: %s"
		switch eth2Client {
		case cfgtypes.ConsensusClient_Lighthouse:
			eth2ClientString = fmt.Sprintf(format, "Lighthouse", cfg.ExternalLighthouse.ContainerTag.Value.(string))
		case cfgtypes.ConsensusClient_Prysm:
			eth2ClientString = fmt.Sprintf(format, "Prysm", cfg.ExternalPrysm.ContainerTag.Value.(string))
		case cfgtypes.ConsensusClient_Teku:
			eth2ClientString = fmt.Sprintf(format, "Teku", cfg.ExternalTeku.ContainerTag.Value.(string))
		default:
			fmt.Errorf("unknown external consensus client [%v]", eth2Client)
		}

	default:
		fmt.Errorf("unknown consensus client mode [%v]", eth2ClientMode)
	}

	// Print version info
	if c.App.Version == rpVersion {
		printGreen(fmt.Sprintf("\nRocket Pool client version: %s - The recommended version is %s\n", c.App.Version, rpVersion))
	} else {
		printYellow(fmt.Sprintf("\nRocket Pool client version: %s - The recommended version is %s\n", c.App.Version, rpVersion))
	}
	if c.App.Version == servVersion {
		printGreen(fmt.Sprintf("Rocket Pool service version: %s\n", servVersion))
	} else {
		printRed(fmt.Sprintf("Rocket Pool service version: %s doesn't match the client version\n", servVersion))
	}
	fmt.Printf("Clients:\n%s\n%s", eth1ClientString, eth2ClientString)
}

func printGreen(str string) {
	fmt.Printf("%s%s%s", colorGreen, str, colorReset)
}

func printYellow(str string) {
	fmt.Printf("%s%s%s", colorYellow, str, colorReset)
}

func printRed(str string) {
	fmt.Printf("%s%s%s", colorRed, str, colorReset)
}

func getArchitecture() string {
	return runtime.GOARCH
}

func isPortOpen(ip string, port uint16) bool {
	address := ""
	proto := "tcp"
	if strings.Contains(ip, ":") {
		proto = "tcp6"
		address = fmt.Sprintf("[%s]:%d", ip, port)
	} else {
		address = fmt.Sprintf("%s:%d", ip, port)
	}

	conn, err := net.DialTimeout(proto, address, 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func getExternalIP() (string, error) {
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	resp, err := client.Get("https://icanhazip.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body and print the client's IP address
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(ip), "\n"), nil
}

func checkDiskSpace(prefix string, rp *rocketpool.Client) (uint64, error) {
	// Check for enough free space
	executionContainerName := prefix + ExecutionContainerSuffix
	volumePath, err := rp.GetClientVolumeSource(executionContainerName, clientDataVolumeName)
	if err != nil {
		return 0, fmt.Errorf("Error getting execution volume source path: %w", err)
	}
	partitions, err := disk.Partitions(true)
	if err != nil {
		return 0, fmt.Errorf("Error getting partition list: %w", err)
	}

	longestPath := 0
	bestPartition := disk.PartitionStat{}
	for _, partition := range partitions {
		if strings.HasPrefix(volumePath, partition.Mountpoint) && len(partition.Mountpoint) > longestPath {
			bestPartition = partition
			longestPath = len(partition.Mountpoint)
		}
	}

	diskUsage, err := disk.Usage(bestPartition.Mountpoint)

	if err != nil {
		return 0, fmt.Errorf("Error getting free disk space available: %w", err)
	}

	return diskUsage.Free, nil
}
