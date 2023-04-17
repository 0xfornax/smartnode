package node

import (
	"fmt"

	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/rocket-pool/smartnode/shared/services/rocketpool"
	cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
)

const signatureVersion = 1

type PersonalSignature struct {
	Address   common.Address `json:"address"`
	Message   string         `json:"msg"`
	Signature string         `json:"sig"`
	Version   string         `json:"version"` // beaconcha.in expects a string
}

func SignArbitraryMessage(c *cli.Context, message string) ([]byte, error) {
	// Get RP client
	rp, err := rocketpool.NewClientFromCtx(c)
	if err != nil {
		return []byte{}, err
	}
	defer rp.Close()

	// Get & check wallet status
	status, err := rp.WalletStatus()
	if err != nil {
		return []byte{}, err
	}

	if !status.WalletInitialized {
		fmt.Println("The node wallet is not initialized.")
		return []byte{}, nil
	}

	if message == "" {
		return []byte{}, fmt.Errorf("signed message can't be empty")
	}

	response, err := rp.SignMessage(message)
	if err != nil {
		return []byte{}, err
	}

	// Print the signature
	formattedSignature := PersonalSignature{
		Address:   status.AccountAddress,
		Message:   message,
		Signature: response.SignedData,
		Version:   fmt.Sprint(signatureVersion),
	}
	bytes, err := json.MarshalIndent(formattedSignature, "", "    ")
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func signMessage(c *cli.Context) error {
	message := c.String("message")
	for message == "" {
		message = cliutils.Prompt("Please enter the message you want to sign: (EIP-191 personal_sign)", "^.+$", "Please enter the message you want to sign: (EIP-191 personal_sign)")
	}
	bytes, err := SignArbitraryMessage(c, c.String("message"))
	if err != nil {
		return err
	}

	fmt.Printf("Signed Message:\n\n%s\n", string(bytes))

	return nil

}
