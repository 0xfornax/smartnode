package service

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"

	"github.com/rocket-pool/smartnode/shared/services"
	"github.com/rocket-pool/smartnode/shared/services/rocketpool"
	"github.com/rocket-pool/smartnode/shared/types/api"
	cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
)

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

	cfg, err := services.GetConfig(c)
	if err != nil {
		return err
	}

	client, err := rocketpool.NewClientFromCtx(c)
	if err != nil {
		return err
	}

	// Get the container prefix
	prefix, err := getContainerPrefix(client)
	if err != nil {
		return fmt.Errorf("Error getting container prefix: %w", err)
	}

	checkCpuFeatures()

	// Response
	response := api.DiagnosticsResponse{}

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

	// Check if the EC P2P port is opened
	wg.Go(func() error {
		var err error
		port := cfg.ExecutionCommon.P2pPort.Value.(uint16)
		ip, err := getExternalIP()
		if err != nil {
			return err
		}

		response.IPV6 = strings.Contains(ip, ":")
		response.ECPort = port
		response.ECPortOpened, err = isPortOpen(ip, port)
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
		response.CCPortOpened, err = isPortOpen(ip, port)
		return err
	})

	// Check free disk space
	wg.Go(func() error {
		var err error
		freeDisk, err := checkDiskSpace(prefix, client)
		if err != nil {
			return err
		}
		response.FreeDisk = freeDisk
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return err
	}

	// Print diagnostics & return
	fmt.Printf("Running system diagnostics:\n")
	fmt.Printf("Architecture: %s\n\n", response.Architecture)

	if response.IPV6 {
		printYellow("External IPv6 address\n")
	} else {
		printGreen("External IPv4 address\n")
	}
	if response.ECPortOpened {
		printGreen(fmt.Sprintf("EC P2P Port: %d opened\n", response.ECPort))
	} else {
		printRed(fmt.Sprintf("EC P2P Port: %d closed!\n\n", response.ECPort))
	}

	if response.CCPortOpened {
		printGreen(fmt.Sprintf("CC P2P Port: %d opened\n", response.CCPort))
	} else {
		printRed(fmt.Sprintf("CC P2P Port: %d closed!\n\n", response.CCPort))
	}

	if response.FreeDisk < 100*1024*1024*1024 {
		printYellow(fmt.Sprintf("Low free disk space: %s\n", humanize.IBytes(response.FreeDisk)))
	} else {
		printGreen(fmt.Sprintf("Free disk space: %s\n", humanize.IBytes(response.FreeDisk)))
	}

	return nil

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

func isPortOpen(ip string, port uint16) (bool, error) {
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
		return false, err
	}
	defer conn.Close()
	return true, nil
}

func getExternalIP() (string, error) {
	resp, err := http.Get("https://icanhazip.com")
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
