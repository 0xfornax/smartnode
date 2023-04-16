package rescue_node

import (
	"github.com/rocket-pool/smartnode/shared/types/config"
)

// Constants
const (
	containerTag string = "rocketpool/rescue-node-addon:v1.0.0"
)

// Configuration for the Graffiti Wall Writer
type RescueNodeAddOnConfig struct {
	Title string `yaml:"-"`

	Enabled config.Parameter `yaml:"enabled,omitempty"`

	InputURL config.Parameter `yaml:"inputUrl,omitempty"`

	// The Docker Hub tag
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`

	// Custom command line flags
	AdditionalFlags config.Parameter `yaml:"additionalFlags,omitempty"`
}

// Creates a new configuration instance
func NewConfig() *RescueNodeAddOnConfig {
	return &RescueNodeAddOnConfig{
		Title: "Rescue Node Add-On Settings",

		Enabled: config.Parameter{
			ID:                   "enabled",
			Name:                 "Enabled",
			Description:          "Enable the Rescue Node Add-On",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: false},
			AffectsContainers:    []config.ContainerID{ContainerID_RescueNodeAddOn, config.ContainerID_Validator},
			EnvironmentVariables: []string{"ADDON_RN_ENABLED"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		InputURL: config.Parameter{
			ID:                   "inputUrl",
			Name:                 "Input URL",
			Description:          "URL or filepath for the input JSON file that contains the graffiti image to write to the wall. By default, this is the Rocket Pool logo.\n\nSee https://gist.github.com/RomiRand/dfa1b5286af3e926deff0be2746db2df for info on making your own images.\n\nNOTE: for local files, you must manually put the file into the `addons/gww` folder of your `rocketpool` directory, and then enter the name of it as `/gww/<filename>` here.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: "https://api.rescuenode.com"},
			AffectsContainers:    []config.ContainerID{ContainerID_RescueNodeAddOn},
			EnvironmentVariables: []string{"ADDON_GWW_INPUT_URL"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:                   "containerTag",
			Name:                 "Container Tag",
			Description:          "The tag name of the container you want to use on Docker Hub.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: containerTag},
			AffectsContainers:    []config.ContainerID{ContainerID_RescueNodeAddOn},
			EnvironmentVariables: []string{"ADDON_RN_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		AdditionalFlags: config.Parameter{
			ID:                   "additionalFlags",
			Name:                 "Additional Flags",
			Description:          "Additional custom command line flags you want to pass to the addon, to take advantage of other settings that the Smartnode's configuration doesn't cover.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{ContainerID_RescueNodeAddOn},
			EnvironmentVariables: []string{"ADDON_RN_ADDITIONAL_FLAGS"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Get the parameters for this config
func (cfg *RescueNodeAddOnConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Enabled,
		&cfg.InputURL,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// The the title for the config
func (cfg *RescueNodeAddOnConfig) GetConfigTitle() string {
	return cfg.Title
}
