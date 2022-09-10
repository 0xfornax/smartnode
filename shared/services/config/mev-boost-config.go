package config

import (
	"github.com/rocket-pool/smartnode/shared/types/config"
)

// Constants
const (
	//mevBoostTag       string = "rocketpool/mev-boost:v0.7.10-portable"
	mevBoostTag          string = "flashbots/mev-boost:v0.8.2"
	mevBoostUrlEnvVar    string = "MEV_BOOST_URL"
	mevBoostRelaysEnvVar string = "MEV_BOOST_RELAYS"
)

// Configuration for MEV-Boost
type MevBoostConfig struct {
	Title string `yaml:"-"`

	// Ownership mode
	Mode config.Parameter `yaml:"mode,omitempty"`

	// Flashbots relay
	FlashbotsRelay config.Parameter `yaml:"flashbotsEnabled,omitempty"`

	// bloXroute ethical relay
	BloxRouteEthicalRelay config.Parameter `yaml:"bloxRouteEthicalEnabled,omitempty"`

	// bloXroute max profit relay
	BloxRouteMaxProfitRelay config.Parameter `yaml:"bloxRouteMaxProfitEnabled,omitempty"`

	// bloXroute regulated relay
	BloxRouteRegulatedRelay config.Parameter `yaml:"bloxRouteRegulatedEnabled,omitempty"`

	// The RPC port
	Port config.Parameter `yaml:"port,omitempty"`

	// Toggle for forwarding the HTTP port outside of Docker
	OpenRpcPort config.Parameter `yaml:"openRpcPort,omitempty"`

	// The Docker Hub tag for MEV-Boost
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`

	// Custom command line flags
	AdditionalFlags config.Parameter `yaml:"additionalFlags,omitempty"`

	// The URL of an external MEV-Boost client
	ExternalUrl config.Parameter `yaml:"externalUrl"`

	///////////////////////////
	// Non-editable settings //
	///////////////////////////

	flashbotsUrls          map[config.Network]string `yaml:"-"`
	bloxRouteEthicalUrls   map[config.Network]string `yaml:"-"`
	bloxRouteMaxProfitUrls map[config.Network]string `yaml:"-"`
	bloxRouteRegulatedUrls map[config.Network]string `yaml:"-"`
}

// Generates a new MEV-Boost configuration
func NewMevBoostConfig(cfg *RocketPoolConfig) *MevBoostConfig {
	return &MevBoostConfig{
		Title: "MEV-Boost Settings",

		Mode: config.Parameter{
			ID:                   "mode",
			Name:                 "MEV-Boost Mode",
			Description:          "Choose whether to let the Smartnode manage your MEV-Boost instance (Locally Managed), or if you manage your own outside of the Smartnode stack (Externally Managed).",
			Type:                 config.ParameterType_Choice,
			Default:              map[config.Network]interface{}{config.Network_All: config.Mode_Local},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2, config.ContainerID_MevBoost},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
			Options: []config.ParameterOption{{
				Name:        "Locally Managed",
				Description: "Allow the Smartnode to manage the MEV-Boost client for you",
				Value:       config.Mode_Local,
			}, {
				Name:        "Externally Managed",
				Description: "Use an existing MEV-Boost client that you manage on your own",
				Value:       config.Mode_External,
			}},
		},

		FlashbotsRelay: config.Parameter{
			ID:                   "flashbotsEnabled",
			Name:                 "Use Flashbots Relay",
			Description:          "Select this to enable the official Flashbots relay. You can enable multiple relays.\n\nFlashbots is the developer of MEV-Boost, and one of the best-known and most trusted relays in the space. It does not filter on MEV type, so it includes sandwiching and front-running bundles.\n\nNote that this relay obeys some government sanctions lists (e.g., OFAC compliance), and will not include transactions from blacklisted addresses.\n\nUses Address Blacklist: YES\nIncludes Frontrunning: YES",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: false},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		BloxRouteEthicalRelay: config.Parameter{
			ID:                   "bloxRouteEthicalEnabled",
			Name:                 "Use bloXroute Ethical Relay",
			Description:          "Select this to enable the \"ethical\" relay from bloXroute. You can enable multiple relays.\n\nThis relay does not include a blacklist, and ignores bundles that extract value from Ethereum users by frontrunning their transactions (\"sandwich attacks\").\n\nUses Address Blacklist: NO\nIncludes Frontrunning: NO",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: false},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		BloxRouteMaxProfitRelay: config.Parameter{
			ID:                   "bloxRouteMaxProfitEnabled",
			Name:                 "Use bloXroute Max Profit Relay",
			Description:          "Select this to enable the \"max profit\" relay from bloXroute. You can enable multiple relays.\n\nThis relay does not include a blacklist, and allows for all types of MEV which includes sandwiching and front-running bundles.\n\nUses Address Blacklist: NO\nIncludes Frontrunning: YES",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: false},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		BloxRouteRegulatedRelay: config.Parameter{
			ID:                   "bloxRouteRegulatedEnabled",
			Name:                 "Use bloXroute Regulated Relay",
			Description:          "Select this to enable the \"regulated\" relay from bloXroute. You can enable multiple relays.\n\nThis relay allows for all types of MEV which includes sandwiching and front-running bundles.\n\nNote that this relay obeys some government sanctions lists (e.g., OFAC compliance), and will not include transactions from blacklisted addresses.\n\nUses Address Blacklist: YES\nIncludes Frontrunning: YES",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: false},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		Port: config.Parameter{
			ID:                   "port",
			Name:                 "Port",
			Description:          "The port that MEV-Boost should serve its API on.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: uint16(18550)},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2, config.ContainerID_MevBoost},
			EnvironmentVariables: []string{"MEV_BOOST_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		OpenRpcPort: config.Parameter{
			ID:                   "openRpcPort",
			Name:                 "Expose API Port",
			Description:          "Expose the API port to your local network, so other local machines can access MEV-Boost's API.",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: false},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:                   "containerTag",
			Name:                 "Container Tag",
			Description:          "The tag name of the MEV-Boost container you want to use on Docker Hub.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: mevBoostTag},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{"MEV_BOOST_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		AdditionalFlags: config.Parameter{
			ID:                   "additionalFlags",
			Name:                 "Additional Flags",
			Description:          "Additional custom command line flags you want to pass to MEV-Boost, to take advantage of other settings that the Smartnode's configuration doesn't cover.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{"MEV_BOOST_ADDITIONAL_FLAGS"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		ExternalUrl: config.Parameter{
			ID:                   "externalUrl",
			Name:                 "External URL",
			Description:          "The URL of the external MEV-Boost client or provider",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{mevBoostUrlEnvVar},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		flashbotsUrls: map[config.Network]string{
			config.Network_Mainnet: "https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net?id=rocketpool",
			config.Network_Prater:  "https://0xafa4c6985aa049fb79dd37010438cfebeb0f2bd42b115b89dd678dab0670c1de38da0c4e9138c9290a398ecd9a0b3110@builder-relay-goerli.flashbots.net?id=rocketpool",
			config.Network_Kiln:    "https://0xb5246e299aeb782fbc7c91b41b3284245b1ed5206134b0028b81dfb974e5900616c67847c2354479934fc4bb75519ee1@builder-relay-kiln.flashbots.net?id=rocketpool",
			config.Network_Ropsten: "https://0xb124d80a00b80815397b4e7f1f05377ccc83aeeceb6be87963ba3649f1e6efa32ca870a88845917ec3f26a8e2aa25c77@builder-relay-ropsten.flashbots.net?id=rocketpool",
		},

		bloxRouteEthicalUrls: map[config.Network]string{
			config.Network_Mainnet: "https://0xad0a8bb54565c2211cee576363f3a347089d2f07cf72679d16911d740262694cadb62d7fd7483f27afd714ca0f1b9118@bloxroute.ethical.blxrbdn.com?id=rocketpool",
			config.Network_Prater:  "",
			config.Network_Kiln:    "",
			config.Network_Ropsten: "",
		},

		bloxRouteMaxProfitUrls: map[config.Network]string{
			config.Network_Mainnet: "https://0x8b5d2e73e2a3a55c6c87b8b6eb92e0149a125c852751db1422fa951e42a09b82c142c3ea98d0d9930b056a3bc9896b8f@bloxroute.max-profit.blxrbdn.com?id=rocketpool",
			config.Network_Prater:  "https://0x821f2a65afb70e7f2e820a925a9b4c80a159620582c1766b1b09729fec178b11ea22abb3a51f07b288be815a1a2ff516@bloxroute.max-profit.builder.goerli.blxrbdn.com?id=rocketpool",
			config.Network_Kiln:    "",
			config.Network_Ropsten: "https://0xb8a0bad3f3a4f0b35418c03357c6d42017582437924a1e1ca6aee2072d5c38d321d1f8b22cd36c50b0c29187b6543b6e@builder-relay.virginia.ropsten.blxrbdn.com?id=rocketpool",
		},

		bloxRouteRegulatedUrls: map[config.Network]string{
			config.Network_Mainnet: "https://0xb0b07cd0abef743db4260b0ed50619cf6ad4d82064cb4fbec9d3ec530f7c5e6793d9f286c4e082c0244ffb9f2658fe88@bloxroute.regulated.blxrbdn.com?id=rocketpool",
			config.Network_Prater:  "",
			config.Network_Kiln:    "",
			config.Network_Ropsten: "",
		},
	}
}

// Get the config.Parameters for this config
func (cfg *MevBoostConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Mode,
		&cfg.FlashbotsRelay,
		&cfg.BloxRouteEthicalRelay,
		&cfg.BloxRouteMaxProfitRelay,
		&cfg.BloxRouteRegulatedRelay,
		&cfg.Port,
		&cfg.OpenRpcPort,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
		&cfg.ExternalUrl,
	}
}

// The the title for the config
func (config *MevBoostConfig) GetConfigTitle() string {
	return config.Title
}
