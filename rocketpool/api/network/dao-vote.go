package network

import (
	"encoding/hex"

	"github.com/rocket-pool/smartnode/rocketpool/api/node"
	"github.com/rocket-pool/smartnode/shared/services"
	"github.com/rocket-pool/smartnode/shared/types/api"
	hexutils "github.com/rocket-pool/smartnode/shared/utils/hex"
	"github.com/urfave/cli"
)

func submitDAOVote(c *cli.Context, message string) (*api.NetworkDAOVoteResponse, error) {

	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	signedBytes, err := w.SignMessage(message)
	if err != nil {
		return nil, err
	}

	// Get node account
	response := api.NetworkDAOVoteResponse{}
	apiDomain := cfg.Smartnode.GetSnapshotApiDomain()
	//snapshotId := cfg.Smartnode.GetSnapshotID()

	_, err = node.SendDAOVoteToSnapshot(apiDomain, message, hexutils.AddPrefix(hex.EncodeToString(signedBytes)))
	if err != nil {
		return nil, err
	}
	return &response, nil
}
