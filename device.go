package netboxapi

import (
	"fmt"
)

type deviceConfig struct {
	Name string
}

type device struct {
	Config      interface{} `json:"local_context_data,omitempty"`
	Name        string      `json:"name"`
	ID          int         `json:"id"`
	PrimaryIP   string      `json:"primary_ip,omitempty"`
	PrimaryIPv4 string      `json:"primary_ip4,omitempty"`
	PrimaryIPv6 string      `json:"primary_ip6,omitempty"`
}

// GetDevice retreaves info for one specific device identified by ID
func (n *NetboxConnection) GetDevice(ID int) device {
	url := fmt.Sprintf("/api/dcim/devices/%d", ID)
	var d device
	err := n.getAPISingle(url, &d)
	if err != nil {
		panic(err)
	}

	return d
}

// GetDevices will retrieve all devices listed in netbox
func (n *NetboxConnection) GetDevices(config ...deviceConfig) []device {
	url := "/api/dcim/devices?limit=500"

	// check to see if deviceConfig was provided
	if len(config) != 0 {
		if config[0].Name != "" {
			url += fmt.Sprintf("&name=%s", config[0].Name)
		}
	}

	var d []device
	err := n.getAPI(url, &d)
	if err != nil {
		panic(err)
	}
	return d
}
