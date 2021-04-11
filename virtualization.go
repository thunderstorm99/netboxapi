package netboxapi

import (
	"fmt"
)

type VMConfig struct {
	Name string
}

type IP struct {
	Address string `json:"address"`
	Family  int    `json:"family"`
}

type VM struct {
	Config      interface{} `json:"local_context_data,omitempty"`
	Name        string      `json:"name"`
	ID          int         `json:"id"`
	PrimaryIP   IP          `json:"primary_ip,omitempty"`
	PrimaryIPv4 IP          `json:"primary_ip4,omitempty"`
	PrimaryIPv6 IP          `json:"primary_ip6,omitempty"`
}

func (n *NetboxConnection) GetVMs(config ...VMConfig) []VM {
	url := "/api/virtualization/virtual-machines/"

	// check if any config was provided
	if len(config) != 0 {
		// check if name was provided
		if config[0].Name != "" {
			url += fmt.Sprintf("?name=%s", config[0].Name)
		}
	}

	var vms []VM
	err := n.getAPI(url, &vms)
	if err != nil {
		panic(err)
	}

	return vms
}
