package api

import (
	"math/big"
)

type DiagnosticsResponse struct {
	Status       string   `json:"status"`
	Error        string   `json:"error"`
	Architecture string   `json:"arch"`
	ECPort       uint16   `json:"ec_port"`
	CCPort       uint16   `json:"cc_port"`
	ExternalIP   string   `json:"ip"`
	IPV6         bool     `json:"json:ipv6"`
	ECPortOpened bool     `json:"ec_port_opened"`
	CCPortOpened bool     `json:"cc_port_opened"`
	FreeDisk     uint64   `json:"free_disk"`
	TotalRam     *big.Int `json:"total_ram"`
	Chronyd      bool     `json:"chronyd"`
}
