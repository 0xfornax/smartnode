package pdao

import (
	"errors"
	"time"

	"github.com/urfave/cli"

	"github.com/rocket-pool/smartnode/shared/utils/api"
	cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
)

// Register subcommands
func RegisterSubcommands(command *cli.Command, name string, aliases []string) {
	command.Subcommands = append(command.Subcommands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Manage the Rocket Pool protocol DAO",
		Subcommands: []cli.Command{

			{
				Name:      "proposals",
				Aliases:   []string{"p"},
				Usage:     "Get the protocol DAO proposals",
				UsageText: "rocketpool api pdao proposals",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(getProposals(c))
					return nil

				},
			},

			{
				Name:      "proposal-details",
				Aliases:   []string{"d"},
				Usage:     "Get details of a proposal",
				UsageText: "rocketpool api pdao proposal-details proposal-id",
				Action: func(c *cli.Context) error {

					// Validate args
					var err error
					if err = cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					id, err := cliutils.ValidateUint("proposal-id", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(getProposal(c, id))
					return nil

				},
			},

			{
				Name:      "can-vote-proposal",
				Usage:     "Check whether the node can vote on a proposal",
				UsageText: "rocketpool api pdao can-vote-proposal proposal-id vote-direction",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					proposalId, err := cliutils.ValidatePositiveUint("proposal ID", c.Args().Get(0))
					if err != nil {
						return err
					}
					voteDir, err := cliutils.ValidateVoteDirection("vote direction", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canVoteOnProposal(c, proposalId, voteDir))
					return nil

				},
			},
			{
				Name:      "vote-proposal",
				Aliases:   []string{"v"},
				Usage:     "Vote on a proposal",
				UsageText: "rocketpool api pdao vote-proposal proposal-id vote-direction",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					proposalId, err := cliutils.ValidatePositiveUint("proposal ID", c.Args().Get(0))
					if err != nil {
						return err
					}
					voteDir, err := cliutils.ValidateVoteDirection("vote direction", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(voteOnProposal(c, proposalId, voteDir))
					return nil

				},
			},

			{
				Name:      "can-execute-proposal",
				Usage:     "Check whether the node can execute a proposal",
				UsageText: "rocketpool api pdao can-execute-proposal proposal-id",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					proposalId, err := cliutils.ValidatePositiveUint("proposal ID", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canExecuteProposal(c, proposalId))
					return nil

				},
			},
			{
				Name:      "execute-proposal",
				Aliases:   []string{"x"},
				Usage:     "Execute a proposal",
				UsageText: "rocketpool api pdao execute-proposal proposal-id",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					proposalId, err := cliutils.ValidatePositiveUint("proposal ID", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(executeProposal(c, proposalId))
					return nil

				},
			},

			{
				Name:      "get-settings",
				Usage:     "Get the Protocol DAO settings",
				UsageText: "rocketpool api pdao get-member-settings",
				Action: func(c *cli.Context) error {

					// Run
					api.PrintResponse(getSettings(c))
					return nil

				},
			},

			{
				Name:      "can-propose-setting",
				Usage:     "Check whether the node can propose a PDAO setting",
				UsageText: "rocketpool api pdao can-propose-setting setting-name value",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					settingName := c.Args().Get(0)
					value := c.Args().Get(1)

					// Run
					api.PrintResponse(canProposeSetting(c, settingName, value))
					return nil

				},
			},
			{
				Name:      "propose-setting",
				Usage:     "Propose updating a PDAO setting (use can-propose-setting to get the pollard)",
				UsageText: "rocketpool api pdao propose-setting setting-name value",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					settingName := c.Args().Get(0)
					value := c.Args().Get(1)
					blockNumber, err := cliutils.ValidatePositiveUint32("block-number", c.Args().Get(2))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(proposeSetting(c, settingName, value, blockNumber))
					return nil

				},
			},

			{
				Name:      "get-rewards-percentages",
				Usage:     "Get the allocation percentages of RPL rewards for the Oracle DAO, the Protocol DAO, and the node operators",
				UsageText: "rocketpool api pdao get-rewards-percentages",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(getRewardsPercentages(c))
					return nil

				},
			},
			{
				Name:      "can-propose-rewards-percentages",
				Usage:     "Check whether the node can propose new RPL rewards allocation percentages for the Oracle DAO, the Protocol DAO, and the node operators",
				UsageText: "rocketpool api pdao can-propose-rewards-percentages node odao pdao",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					node, err1 := cliutils.ValidateBigInt("node", c.Args().Get(0))
					odao, err2 := cliutils.ValidateBigInt("odao", c.Args().Get(1))
					pdao, err3 := cliutils.ValidateBigInt("pdao", c.Args().Get(2))
					err := errors.Join(err1, err2, err3)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canProposeRewardsPercentages(c, node, odao, pdao))
					return nil

				},
			},
			{
				Name:      "propose-rewards-percentages",
				Usage:     "Propose new RPL rewards allocation percentages for the Oracle DAO, the Protocol DAO, and the node operators",
				UsageText: "rocketpool api pdao propose-rewards-percentages node odao pdao block-number",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 4); err != nil {
						return err
					}
					node, err1 := cliutils.ValidateBigInt("node", c.Args().Get(0))
					odao, err2 := cliutils.ValidateBigInt("odao", c.Args().Get(1))
					pdao, err3 := cliutils.ValidateBigInt("pdao", c.Args().Get(2))
					blockNumber, err4 := cliutils.ValidateUint32("blockNumber", c.Args().Get(3))
					err := errors.Join(err1, err2, err3, err4)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(proposeRewardsPercentages(c, node, odao, pdao, blockNumber))
					return nil

				},
			},

			{
				Name:      "can-propose-one-time-spend",
				Usage:     "Check whether the node can propose a one-time spend of the Protocol DAO's treasury",
				UsageText: "rocketpool api pdao can-propose-one-time-spend invoice-id recipient amount",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					invoiceID := c.Args().Get(0)
					recipient, err1 := cliutils.ValidateAddress("recipient", c.Args().Get(1))
					amount, err2 := cliutils.ValidateBigInt("amount", c.Args().Get(2))
					err := errors.Join(err1, err2)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canProposeOneTimeSpend(c, invoiceID, recipient, amount))
					return nil

				},
			},
			{
				Name:      "propose-one-time-spend",
				Usage:     "Propose a one-time spend of the Protocol DAO's treasury",
				UsageText: "rocketpool api pdao propose-one-time-spend invoice-id recipient amount block-number",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 4); err != nil {
						return err
					}
					invoiceID := c.Args().Get(0)
					recipient, err1 := cliutils.ValidateAddress("recipient", c.Args().Get(1))
					amount, err2 := cliutils.ValidateBigInt("amount", c.Args().Get(2))
					blockNumber, err3 := cliutils.ValidateUint32("blockNumber", c.Args().Get(3))
					err := errors.Join(err1, err2, err3)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(proposeOneTimeSpend(c, invoiceID, recipient, amount, blockNumber))
					return nil

				},
			},

			{
				Name:      "can-propose-recurring-spend",
				Usage:     "Check whether the node can propose a recurring spend of the Protocol DAO's treasury",
				UsageText: "rocketpool api pdao can-propose-recurring-spend contract-name recipient amount-per-period period-length start-time number-of-periods",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 6); err != nil {
						return err
					}
					contractName := c.Args().Get(0)
					recipient, err1 := cliutils.ValidateAddress("recipient", c.Args().Get(1))
					amountPerPeriod, err2 := cliutils.ValidateBigInt("amount-per-period", c.Args().Get(2))
					periodLength, err3 := cliutils.ValidateDuration("period-length", c.Args().Get(3))
					startTime, err4 := cliutils.ValidatePositiveUint("start-time", c.Args().Get(4))
					numberOfPeriods, err5 := cliutils.ValidatePositiveUint("number-of-periods", c.Args().Get(5))
					err := errors.Join(err1, err2, err3, err4, err5)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canProposeRecurringSpend(c, contractName, recipient, amountPerPeriod, periodLength, time.Unix(int64(startTime), 0), numberOfPeriods))
					return nil

				},
			},
			{
				Name:      "propose-recurring-spend",
				Usage:     "Propose a recurring spend of the Protocol DAO's treasury",
				UsageText: "rocketpool api pdao propose-recurring-spend contract-name recipient amount-per-period period-length start-time number-of-periods block-number",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 7); err != nil {
						return err
					}
					contractName := c.Args().Get(0)
					recipient, err1 := cliutils.ValidateAddress("recipient", c.Args().Get(1))
					amountPerPeriod, err2 := cliutils.ValidateBigInt("amount-per-period", c.Args().Get(2))
					periodLength, err3 := cliutils.ValidateDuration("period-length", c.Args().Get(3))
					startTime, err4 := cliutils.ValidatePositiveUint("start-time", c.Args().Get(4))
					numberOfPeriods, err5 := cliutils.ValidatePositiveUint("number-of-periods", c.Args().Get(5))
					blockNumber, err6 := cliutils.ValidateUint32("blockNumber", c.Args().Get(6))
					err := errors.Join(err1, err2, err3, err4, err5, err6)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(proposeRecurringSpend(c, contractName, recipient, amountPerPeriod, periodLength, time.Unix(int64(startTime), 0), numberOfPeriods, blockNumber))
					return nil

				},
			},

			{
				Name:      "can-propose-recurring-spend-update",
				Usage:     "Check whether the node can propose an update to an existing recurring spend plan",
				UsageText: "rocketpool api pdao can-propose-recurring-spend-update contract-name recipient amount-per-period period-length number-of-periods",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 5); err != nil {
						return err
					}
					contractName := c.Args().Get(0)
					recipient, err1 := cliutils.ValidateAddress("recipient", c.Args().Get(1))
					amountPerPeriod, err2 := cliutils.ValidateBigInt("amount-per-period", c.Args().Get(2))
					periodLength, err3 := cliutils.ValidateDuration("period-length", c.Args().Get(3))
					numberOfPeriods, err4 := cliutils.ValidatePositiveUint("number-of-periods", c.Args().Get(4))
					err := errors.Join(err1, err2, err3, err4)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canProposeRecurringSpendUpdate(c, contractName, recipient, amountPerPeriod, periodLength, numberOfPeriods))
					return nil

				},
			},
			{
				Name:      "propose-recurring-spend-update",
				Usage:     "Propose an update to an existing recurring spend plan",
				UsageText: "rocketpool api pdao propose-recurring-spend-update contract-name recipient amount-per-period period-length number-of-periods block-number",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 6); err != nil {
						return err
					}
					contractName := c.Args().Get(0)
					recipient, err1 := cliutils.ValidateAddress("recipient", c.Args().Get(1))
					amountPerPeriod, err2 := cliutils.ValidateBigInt("amount-per-period", c.Args().Get(2))
					periodLength, err3 := cliutils.ValidateDuration("period-length", c.Args().Get(3))
					numberOfPeriods, err4 := cliutils.ValidatePositiveUint("number-of-periods", c.Args().Get(4))
					blockNumber, err5 := cliutils.ValidateUint32("blockNumber", c.Args().Get(5))
					err := errors.Join(err1, err2, err3, err4, err5)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(proposeRecurringSpendUpdate(c, contractName, recipient, amountPerPeriod, periodLength, numberOfPeriods, blockNumber))
					return nil

				},
			},

			{
				Name:      "can-propose-invite-to-security-council",
				Usage:     "Check whether the node can invite someone to the security council",
				UsageText: "rocketpool api pdao can-propose-invite-to-security-council id address",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					id := c.Args().Get(0)
					address, err1 := cliutils.ValidateAddress("address", c.Args().Get(1))
					err := errors.Join(err1)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canProposeInviteToSecurityCouncil(c, id, address))
					return nil

				},
			},
			{
				Name:      "propose-invite-to-security-council",
				Usage:     "Propose inviting someone to the security council",
				UsageText: "rocketpool api pdao propose-invite-to-security-council id address block-number",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					id := c.Args().Get(0)
					address, err1 := cliutils.ValidateAddress("address", c.Args().Get(1))
					blockNumber, err2 := cliutils.ValidateUint32("blockNumber", c.Args().Get(2))
					err := errors.Join(err1, err2)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(proposeInviteToSecurityCouncil(c, id, address, blockNumber))
					return nil

				},
			},

			{
				Name:      "can-propose-kick-from-security-council",
				Usage:     "Check whether the node can kick someone from the security council",
				UsageText: "rocketpool api pdao can-propose-kick-from-security-council address",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					address, err1 := cliutils.ValidateAddress("address", c.Args().Get(0))
					err := errors.Join(err1)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canProposeKickFromSecurityCouncil(c, address))
					return nil

				},
			},
			{
				Name:      "propose-kick-from-security-council",
				Usage:     "Propose kicking someone from the security council",
				UsageText: "rocketpool api pdao propose-kick-from-security-council address block-number",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					address, err1 := cliutils.ValidateAddress("address", c.Args().Get(0))
					blockNumber, err2 := cliutils.ValidateUint32("blockNumber", c.Args().Get(1))
					err := errors.Join(err1, err2)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(proposeKickFromSecurityCouncil(c, address, blockNumber))
					return nil

				},
			},

			{
				Name:      "can-propose-kick-multi-from-security-council",
				Usage:     "Check whether the node can kick multiple members from the security council",
				UsageText: "rocketpool api pdao can-propose-kick-multi-from-security-council addresses",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					addresses, err1 := cliutils.ValidateAddresses("address", c.Args().Get(0))
					err := errors.Join(err1)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canProposeKickMultiFromSecurityCouncil(c, addresses))
					return nil

				},
			},
			{
				Name:      "propose-kick-multi-from-security-council",
				Usage:     "Propose kicking multiple members from the security council",
				UsageText: "rocketpool api pdao propose-kick-multi-from-security-council addresses block-number",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					addresses, err1 := cliutils.ValidateAddresses("addresess", c.Args().Get(0))
					blockNumber, err2 := cliutils.ValidateUint32("blockNumber", c.Args().Get(1))
					err := errors.Join(err1, err2)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(proposeKickMultiFromSecurityCouncil(c, addresses, blockNumber))
					return nil

				},
			},

			{
				Name:      "can-propose-replace-member-of-security-council",
				Usage:     "Check whether the node can propose replacing someone on the security council with another member",
				UsageText: "rocketpool api pdao can-propose-replace-member-of-security-council existing-address new-id new-address",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					existingAddress, err1 := cliutils.ValidateAddress("existingAddress", c.Args().Get(0))
					newID := c.Args().Get(1)
					newAddress, err2 := cliutils.ValidateAddress("newAddress", c.Args().Get(2))
					err := errors.Join(err1, err2)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canProposeReplaceMemberOfSecurityCouncil(c, existingAddress, newID, newAddress))
					return nil

				},
			},
			{
				Name:      "propose-replace-member-of-security-council",
				Usage:     "Propose replacing someone on the security council with another member",
				UsageText: "rocketpool api pdao propose-replace-member-of-security-council existing-address new-id new-address block-number",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 4); err != nil {
						return err
					}
					existingAddress, err1 := cliutils.ValidateAddress("existingAddress", c.Args().Get(0))
					newID := c.Args().Get(1)
					newAddress, err2 := cliutils.ValidateAddress("newAddress", c.Args().Get(2))
					blockNumber, err3 := cliutils.ValidateUint32("blockNumber", c.Args().Get(3))
					err := errors.Join(err1, err2, err3)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(proposeReplaceMemberOfSecurityCouncil(c, existingAddress, newID, newAddress, blockNumber))
					return nil

				},
			},
		},
	})
}
