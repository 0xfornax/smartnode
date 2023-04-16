package rescue_node

import (
	"fmt"

	"github.com/rocket-pool/smartnode/shared/types/addons"
	cfgtypes "github.com/rocket-pool/smartnode/shared/types/config"
)

const (
	ContainerID_RescueNodeAddOn  cfgtypes.ContainerID = "rn"
	RescueNodeAddOnContainerName string               = "addon_rn"
)

type RescueNodeAddOn struct {
	cfg *RescueNodeAddOnConfig `yaml:"config,omitempty"`
}

func NewRescueNodeAddOn() addons.SmartnodeAddon {
	return &RescueNodeAddOn{
		cfg: NewConfig(),
	}
}

func (rn *RescueNodeAddOn) GetName() string {
	return "Rescue Node Add-on"
}

func (rn *RescueNodeAddOn) GetDescription() string {
	return "This addon adds support for connecting to the Rocket Pool Rescue Node"
}

func (rn *RescueNodeAddOn) GetConfig() cfgtypes.Config {
	return rn.cfg
}

func (rn *RescueNodeAddOn) GetContainerName() string {
	return fmt.Sprint(ContainerID_RescueNodeAddOn)
}

func (rn *RescueNodeAddOn) GetEnabledParameter() *cfgtypes.Parameter {
	return &rn.cfg.Enabled
}

func (rn *RescueNodeAddOn) GetContainerTag() string {
	return containerTag
}

func (rn *RescueNodeAddOn) UpdateEnvVars(envVars map[string]string) error {
	if rn.cfg.Enabled.Value == true {
		cfgtypes.AddParametersToEnvVars(rn.cfg.GetParameters(), envVars)
	}
	return nil
}
