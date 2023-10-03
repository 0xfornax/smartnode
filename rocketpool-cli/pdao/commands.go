package pdao

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/rocket-pool/rocketpool-go/settings/protocol"
	cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
)

const (
	boolUsage             string = "specify 'true', 'false', 'yes', or 'no'"
	floatEthUsage         string = "specify an amount of ETH (e.g., '16.0')"
	floatRplUsage         string = "specify an amount of RPL (e.g., '16.0')"
	blockCountUsage       string = "specify a number, in blocks (e.g., '40000')"
	percentUsage          string = "specify a percentage between 0 and 1 (e.g., '0.51' for 51%)"
	unboundedPercentUsage string = "specify a percentage that can go over 100% (e.g., '1.5' for 150%)"
	uintUsage             string = "specify an integer (e.g., '50')"
	durationUsage         string = "specify a duration using hours, minutes, and seconds (e.g., '20m' or '72h0m0s')"
)

// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Manage the Rocket Pool Protocol DAO",
		Subcommands: []cli.Command{

			{
				Name:      "settings",
				Aliases:   []string{"s"},
				Usage:     "Show all of the current Protocol DAO settings and values",
				UsageText: "rocketpool pdao settings",
				Action: func(c *cli.Context) error {

					// Run
					return getSettings(c)

				},
			},

			{
				Name:    "propose",
				Aliases: []string{"p"},
				Usage:   "Make a Protocol DAO proposal",
				Subcommands: []cli.Command{

					{
						Name:    "setting",
						Aliases: []string{"s"},
						Usage:   "Make a Protocol DAO setting proposal",
						Subcommands: []cli.Command{

							{
								Name:    "auction",
								Aliases: []string{"a"},
								Usage:   "Auction settings",
								Subcommands: []cli.Command{

									{
										Name:      "is-create-lot-enabled",
										Aliases:   []string{"icle"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.CreateLotEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting auction is-create-lot-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingAuctionIsCreateLotEnabled(c, value)

										},
									},

									{
										Name:      "is-bid-on-lot-enabled",
										Aliases:   []string{"ibole"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.BidOnLotEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting auction is-bid-on-lot-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingAuctionIsBidOnLotEnabled(c, value)

										},
									},

									{
										Name:      "lot-minimum-eth-value",
										Aliases:   []string{"lminev"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.LotMinimumEthValueSettingPath, floatEthUsage),
										UsageText: "rocketpool pdao propose setting auction lot-minimum-eth-value value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingAuctionLotMinimumEthValue(c, value)

										},
									},

									{
										Name:      "lot-maximum-eth-value",
										Aliases:   []string{"lmaxev"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.LotMaximumEthValueSettingPath, floatEthUsage),
										UsageText: "rocketpool pdao propose setting auction lot-maximum-eth-value value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingAuctionLotMaximumEthValue(c, value)

										},
									},

									{
										Name:      "lot-duration",
										Aliases:   []string{"ld"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.LotDurationSettingPath, blockCountUsage),
										UsageText: "rocketpool pdao propose setting auction lot-duration value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidatePositiveUint("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingAuctionLotDuration(c, value)

										},
									},

									{
										Name:      "lot-starting-price-ratio",
										Aliases:   []string{"lspr"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.LotStartingPriceRatioSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting auction lot-starting-price-ratio value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingAuctionLotStartingPriceRatio(c, value)

										},
									},

									{
										Name:      "lot-reserve-price-ratio",
										Aliases:   []string{"lrpr"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.LotReservePriceRatioSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting auction lot-reserve-price-ratio value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingAuctionLotReservePriceRatio(c, value)

										},
									},
								},
							},

							{
								Name:    "deposit",
								Aliases: []string{"d"},
								Usage:   "Deposit pool settings",
								Subcommands: []cli.Command{

									{
										Name:      "is-depositing-enabled",
										Aliases:   []string{"ide"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.DepositEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting deposit is-depositing-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingDepositIsDepositingEnabled(c, value)

										},
									},

									{
										Name:      "are-deposit-assignments-enabled",
										Aliases:   []string{"adae"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.AssignDepositsEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting deposit are-deposit-assignments-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingDepositAreDepositAssignmentsEnabled(c, value)

										},
									},

									{
										Name:      "minimum-deposit",
										Aliases:   []string{"md"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MinimumDepositSettingPath, floatEthUsage),
										UsageText: "rocketpool pdao propose setting deposit minimum-deposit value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingDepositMinimumDeposit(c, value)

										},
									},

									{
										Name:      "maximum-deposit-pool-size",
										Aliases:   []string{"mdps"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MaximumDepositPoolSizeSettingPath, floatEthUsage),
										UsageText: "rocketpool pdao propose setting deposit maximum-deposit-pool-size value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingDepositMaximumDepositPoolSize(c, value)

										},
									},

									{
										Name:      "maximum-assignments-per-deposit",
										Aliases:   []string{"mapd"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MaximumDepositAssignmentsSettingPath, uintUsage),
										UsageText: "rocketpool pdao propose setting deposit maximum-assignments-per-deposit value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidatePositiveUint("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingDepositMaximumAssignmentsPerDeposit(c, value)

										},
									},

									{
										Name:      "maximum-socialised-assignments-per-deposit",
										Aliases:   []string{"msapd"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MaximumSocializedDepositAssignmentsSettingPath, uintUsage),
										UsageText: "rocketpool pdao propose setting deposit maximum-socialised-assignments-per-deposit value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidatePositiveUint("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingDepositMaximumSocialisedAssignmentsPerDeposit(c, value)

										},
									},

									{
										Name:      "deposit-fee",
										Aliases:   []string{"df"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.DepositFeeSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting deposit deposit-fee value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingDepositDepositFee(c, value)

										},
									},
								},
							},

							{
								Name:    "minipool",
								Aliases: []string{"m"},
								Usage:   "Minipool settings",
								Subcommands: []cli.Command{

									{
										Name:      "is-submit-withdrawable-enabled",
										Aliases:   []string{"iswe"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MinipoolSubmitWithdrawableEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting minipool is-submit-withdrawable-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingMinipoolIsSubmitWithdrawableEnabled(c, value)

										},
									},

									{
										Name:      "launch-timeout",
										Aliases:   []string{"lt"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MinipoolLaunchTimeoutSettingPath, durationUsage),
										UsageText: "rocketpool pdao propose setting minipool launch-timeout value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateDuration("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingMinipoolLaunchTimeout(c, value)

										},
									},

									{
										Name:      "is-bond-reduction-enabled",
										Aliases:   []string{"ibre"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.BondReductionEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting minipool is-bond-reduction-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingMinipoolIsBondReductionEnabled(c, value)

										},
									},

									{
										Name:      "max-count",
										Aliases:   []string{"mc"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MaximumMinipoolCountSettingPath, uintUsage),
										UsageText: "rocketpool pdao propose setting minipool max-count value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidatePositiveUint("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingMinipoolMaximumCount(c, value)

										},
									},

									{
										Name:      "user-distribute-window-start",
										Aliases:   []string{"udws"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MinipoolUserDistributeWindowStartSettingPath, durationUsage),
										UsageText: "rocketpool pdao propose setting minipool user-distribute-window-start value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateDuration("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingMinipoolUserDistributeWindowStart(c, value)

										},
									},

									{
										Name:      "user-distribute-window-length",
										Aliases:   []string{"udwl"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MinipoolUserDistributeWindowLengthSettingPath, durationUsage),
										UsageText: "rocketpool pdao propose setting minipool user-distribute-window-length value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateDuration("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingMinipoolUserDistributeWindowLength(c, value)

										},
									},
								},
							},

							{
								Name:    "network",
								Aliases: []string{"ne"},
								Usage:   "Network settings",
								Subcommands: []cli.Command{

									{
										Name:      "oracle-dao-consensus-threshold",
										Aliases:   []string{"odct"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.NodeConsensusThresholdSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting network oracle-dao-consensus-threshold value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkOracleDaoConsensusThreshold(c, value)

										},
									},

									{
										Name:      "node-penalty-threshold",
										Aliases:   []string{"npt"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.NetworkPenaltyThresholdSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting network node-penalty-threshold value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkNodePenaltyThreshold(c, value)

										},
									},

									{
										Name:      "per-penalty-rate",
										Aliases:   []string{"ppr"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.NetworkPenaltyPerRateSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting network per-penalty-rate value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkPerPenaltyRate(c, value)

										},
									},

									{
										Name:      "is-submit-balances-enabled",
										Aliases:   []string{"isbe"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.SubmitBalancesEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting network is-submit-balances-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkIsSubmitBalancesEnabled(c, value)

										},
									},

									{
										Name:      "submit-balances-frequency",
										Aliases:   []string{"sbf"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.SubmitBalancesFrequencySettingPath, blockCountUsage),
										UsageText: "rocketpool pdao propose setting network submit-balances-frequency value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidatePositiveUint("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkSubmitBalancesFrequency(c, value)

										},
									},

									{
										Name:      "is-submit-prices-enabled",
										Aliases:   []string{"ispe"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.SubmitPricesEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting network is-submit-prices-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkIsSubmitPricesEnabled(c, value)

										},
									},

									{
										Name:      "submit-prices-frequency",
										Aliases:   []string{"spf"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.SubmitPricesFrequencySettingPath, blockCountUsage),
										UsageText: "rocketpool pdao propose setting network submit-prices-frequency value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidatePositiveUint("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkSubmitPricesFrequency(c, value)

										},
									},

									{
										Name:      "minimum-node-fee",
										Aliases:   []string{"minnf"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MinimumNodeFeeSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting network minimum-node-fee value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkMinimumNodeFee(c, value)

										},
									},

									{
										Name:      "target-node-fee",
										Aliases:   []string{"tnf"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.TargetNodeFeeSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting network target-node-fee value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkTargetNodeFee(c, value)

										},
									},

									{
										Name:      "maximum-node-fee",
										Aliases:   []string{"maxnf"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.TargetNodeFeeSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting network maximum-node-fee value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkMaximumNodeFee(c, value)

										},
									},

									{
										Name:      "node-fee-demand-range",
										Aliases:   []string{"nfdr"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.NodeFeeDemandRangeSettingPath, floatEthUsage),
										UsageText: "rocketpool pdao propose setting network node-fee-demand-range value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkNodeFeeDemandRange(c, value)

										},
									},

									{
										Name:      "target-reth-collateral-rate",
										Aliases:   []string{"trcr"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.TargetRethCollateralRateSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting network target-reth-collateral-rate value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkTargetRethCollateralRate(c, value)

										},
									},

									{
										Name:      "is-submit-rewards-enabled",
										Aliases:   []string{"isre"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.SubmitRewardsEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting network is-submit-rewards-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNetworkIsSubmitRewardsEnabled(c, value)

										},
									},
								},
							},

							{
								Name:    "node",
								Aliases: []string{"no"},
								Usage:   "Node settings",
								Subcommands: []cli.Command{

									{
										Name:      "is-registration-enabled",
										Aliases:   []string{"ire"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.NodeRegistrationEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting node is-registration-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNodeIsRegistrationEnabled(c, value)

										},
									},

									{
										Name:      "is-smoothing-pool-registration-enabled",
										Aliases:   []string{"ispre"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.SmoothingPoolRegistrationEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting node is-smoothing-pool-registration-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNodeIsSmoothingPoolRegistrationEnabled(c, value)

										},
									},

									{
										Name:      "is-depositing-enabled",
										Aliases:   []string{"ide"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.NodeDepositEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting node is-depositing-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNodeIsDepositingEnabled(c, value)

										},
									},

									{
										Name:      "are-vacant-minipools-enabled",
										Aliases:   []string{"avme"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.VacantMinipoolsEnabledSettingPath, boolUsage),
										UsageText: "rocketpool pdao propose setting node are-vacant-minipools-enabled value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateBool("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNodeAreVacantMinipoolsEnabled(c, value)

										},
									},

									{
										Name:      "minimum-per-minipool-stake",
										Aliases:   []string{"minpms"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MinimumPerMinipoolStakeSettingPath, unboundedPercentUsage),
										UsageText: "rocketpool pdao propose setting node minimum-per-minipool-stake value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNodeMinimumPerMinipoolStake(c, value)

										},
									},

									{
										Name:      "maximum-per-minipool-stake",
										Aliases:   []string{"minpms"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.MaximumPerMinipoolStakeSettingPath, unboundedPercentUsage),
										UsageText: "rocketpool pdao propose setting node maximum-per-minipool-stake value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingNodeMaximumPerMinipoolStake(c, value)

										},
									},
								},
							},

							{
								Name:    "proposals",
								Aliases: []string{"p"},
								Usage:   "Proposal settings",
								Subcommands: []cli.Command{

									{
										Name:      "vote-time",
										Aliases:   []string{"vt"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.VoteTimeSettingPath, durationUsage),
										UsageText: "rocketpool pdao propose setting proposals vote-time value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateDuration("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingProposalsVoteTime(c, value)

										},
									},

									{
										Name:      "vote-delay-time",
										Aliases:   []string{"vdt"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.VoteDelayTimeSettingPath, durationUsage),
										UsageText: "rocketpool pdao propose setting proposals vote-dalay-time value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateDuration("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingProposalsVoteDelayTime(c, value)

										},
									},

									{
										Name:      "execute-time",
										Aliases:   []string{"et"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.ExecuteTimeSettingPath, durationUsage),
										UsageText: "rocketpool pdao propose setting proposals execute-time value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateDuration("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingProposalsExecuteTime(c, value)

										},
									},

									{
										Name:      "proposal-bond",
										Aliases:   []string{"pb"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.ProposalBondSettingPath, floatRplUsage),
										UsageText: "rocketpool pdao propose setting proposals proposal-bond value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingProposalsProposalBond(c, value)

										},
									},

									{
										Name:      "challenge-bond",
										Aliases:   []string{"cb"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.ChallengeBondSettingPath, floatRplUsage),
										UsageText: "rocketpool pdao propose setting proposals challenge-bond value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingProposalsChallengeBond(c, value)

										},
									},

									{
										Name:      "challenge-period",
										Aliases:   []string{"cp"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.ChallengePeriodSettingPath, durationUsage),
										UsageText: "rocketpool pdao propose setting proposals challenge-period value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateDuration("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingProposalsChallengePeriod(c, value)

										},
									},

									{
										Name:      "quorum",
										Aliases:   []string{"q"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.ProposalQuorumSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting proposals quorum value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingProposalsQuorum(c, value)

										},
									},

									{
										Name:      "veto-quorum",
										Aliases:   []string{"vq"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.ProposalVetoQuorumSettingPath, percentUsage),
										UsageText: "rocketpool pdao propose setting proposals veto-quorum value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := parseFloat(c, c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingProposalsVetoQuorum(c, value)

										},
									},

									{
										Name:      "max-block-age",
										Aliases:   []string{"mba"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.ProposalMaxBlockAgeSettingPath, blockCountUsage),
										UsageText: "rocketpool pdao propose setting proposals max-block-age value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidatePositiveUint("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingProposalsMaxBlockAge(c, value)

										},
									},
								},
							},

							{
								Name:    "rewards",
								Aliases: []string{"r"},
								Usage:   "Rewards settings",
								Subcommands: []cli.Command{

									{
										Name:      "interval-time",
										Aliases:   []string{"it"},
										Usage:     fmt.Sprintf("Propose updating the %s setting; %s", protocol.RewardsClaimIntervalTimeSettingPath, durationUsage),
										UsageText: "rocketpool pdao propose setting rewards interval-time value",
										Action: func(c *cli.Context) error {

											// Validate args
											if err := cliutils.ValidateArgCount(c, 1); err != nil {
												return err
											}
											value, err := cliutils.ValidateDuration("value", c.Args().Get(0))
											if err != nil {
												return err
											}

											// Run
											return proposeSettingRewardsIntervalTime(c, value)

										},
									},
								},
							},
						},
					},
				},
			},

			{
				Name:    "proposals",
				Aliases: []string{"o"},
				Usage:   "Manage oracle DAO proposals",
				Subcommands: []cli.Command{

					{
						Name:      "list",
						Aliases:   []string{"l"},
						Usage:     "List the oracle DAO proposals",
						UsageText: "rocketpool pdao proposals list",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "states, s",
								Usage: "Comma separated list of states to filter ('pending', 'active', 'succeeded', 'executed', 'cancelled', 'defeated', or 'expired')",
								Value: "",
							},
						},
						Action: func(c *cli.Context) error {

							// Validate args
							if err := cliutils.ValidateArgCount(c, 0); err != nil {
								return err
							}

							// Run
							return getProposals(c, c.String("states"))

						},
					},

					{
						Name:      "details",
						Aliases:   []string{"d"},
						Usage:     "View proposal details",
						UsageText: "rocketpool pdao proposals details proposal-id",
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
							return getProposal(c, id)

						},
					},

					{
						Name:      "vote",
						Aliases:   []string{"v"},
						Usage:     "Vote on a proposal",
						UsageText: "rocketpool pdao proposals vote [options]",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "proposal, p",
								Usage: "The ID of the proposal to vote on",
							},
							cli.StringFlag{
								Name:  "support, s",
								Usage: "Whether to support the proposal ('yes' or 'no')",
							},
							cli.BoolFlag{
								Name:  "yes, y",
								Usage: "Automatically confirm vote",
							},
						},
						Action: func(c *cli.Context) error {

							// Validate args
							if err := cliutils.ValidateArgCount(c, 0); err != nil {
								return err
							}

							// Validate flags
							if c.String("proposal") != "" {
								if _, err := cliutils.ValidatePositiveUint("proposal ID", c.String("proposal")); err != nil {
									return err
								}
							}
							if c.String("support") != "" {
								if _, err := cliutils.ValidateBool("support", c.String("support")); err != nil {
									return err
								}
							}

							// Run
							return voteOnProposal(c)

						},
					},

					{
						Name:      "execute",
						Aliases:   []string{"x"},
						Usage:     "Execute a proposal",
						UsageText: "rocketpool pdao proposals execute [options]",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "proposal, p",
								Usage: "The ID of the proposal to execute (or 'all')",
							},
						},
						Action: func(c *cli.Context) error {

							// Validate args
							if err := cliutils.ValidateArgCount(c, 0); err != nil {
								return err
							}

							// Validate flags
							if c.String("proposal") != "" && c.String("proposal") != "all" {
								if _, err := cliutils.ValidatePositiveUint("proposal ID", c.String("proposal")); err != nil {
									return err
								}
							}

							// Run
							return executeProposal(c)

						},
					},
				},
			},
		},
	})
}
