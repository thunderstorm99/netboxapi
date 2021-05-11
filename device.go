package netboxapi

import (
	"fmt"
)

type DeviceConfig struct {
	Name   string
	Tenant string
}

type Device struct {
	Config      interface{} `json:"local_context_data,omitempty"`
	Name        string      `json:"name"`
	ID          int         `json:"id"`
	PrimaryIP   DeviceIP    `json:"primary_ip,omitempty"`
	PrimaryIPv4 DeviceIP    `json:"primary_ip4,omitempty"`
	PrimaryIPv6 DeviceIP    `json:"primary_ip6,omitempty"`
}

type DeviceIP struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Familiy int    `json:"family"`
	Address string `json:"address"`
}

// GetDevice retreaves info for one specific device identified by ID
func (n *NetboxConnection) GetDevice(ID int) Device {
	url := fmt.Sprintf("/api/dcim/devices/%d", ID)
	var d Device
	err := n.getAPISingle(url, &d)
	if err != nil {
		panic(err)
	}

	return d
}

// GetDevices will retrieve all devices listed in netbox
func (n *NetboxConnection) GetDevices(config ...DeviceConfig) []Device {
	url := "/api/dcim/devices?limit=500"

	// check to see if deviceConfig was provided
	if len(config) != 0 {
		if config[0].Name != "" {
			url += fmt.Sprintf("&name=%s", config[0].Name)
		}
		if config[0].Tenant != "" {
			url += fmt.Sprintf("&tenant=%s", config[0].Tenant)
		}
	}

	var d []Device
	err := n.getAPI(url, &d)
	if err != nil {
		panic(err)
	}
	return d
}
