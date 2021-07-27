package netboxapi

import (
	"fmt"
	"strings"
)

type DeviceConfig struct {
	Name   string
	Tenant string
	Role   string
	Rack   int
}

type Device struct {
	Config      interface{} `json:"local_context_data,omitempty"`
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	DisplayName string      `json:"display_name"`
	PrimaryIP   DeviceIP    `json:"primary_ip,omitempty"`
	PrimaryIPv4 DeviceIP    `json:"primary_ip4,omitempty"`
	PrimaryIPv6 DeviceIP    `json:"primary_ip6,omitempty"`
	DeviceType  DeviceType  `json:"device_type"`
	DeviceRole  Short       `json:"device_role"`
	Serial      string      `json:"serial"`
	Tenant      Short       `json:"tenant"`
	Platform    Short       `json:"platform,omitempty"`
	AssetTag    string      `json:"asset_tag"`
	Site        Short       `json:"site"`
	Rack        struct {
		ID          int    `json:"id"`
		URL         string `json:"url"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
	} `json:"rack"`
	Position     int        `json:"position"`
	Face         ValueLabel `json:"face"`
	ParentDevice Short      `json:"parent_device"`
	Status       Short      `json:"status"`
	Comments     string     `json:"comments"`
}

type DeviceType struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`
	Model        string `json:"model"`
	Slug         string `json:"slug"`
	Manufacturer Short  `json:"manufacturer"`
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
func (n *NetboxConnection) GetDevices(config ...DeviceConfig) ([]Device, error) {
	url := "/api/dcim/devices?limit=500"

	// check to see if deviceConfig was provided
	if len(config) != 0 {
		// check if name was provided
		if config[0].Name != "" {
			url += fmt.Sprintf("&name=%s", config[0].Name)
		}

		// check if tenant was provided
		if config[0].Tenant != "" {
			url += fmt.Sprintf("&tenant=%s", config[0].Tenant)
		}

		// check if role was provided
		if config[0].Role != "" {

			// Split multiple roles
			roleSplit := strings.Split(config[0].Role, ",")

			for _, role := range roleSplit {
				// append for earch role
				url += fmt.Sprintf("&role=%s", role)
			}
		}

		// check if rack was provided
		if config[0].Rack != 0 {
			url += fmt.Sprintf("&rack_id=%d", config[0].Rack)
		}
	}

	var d []Device
	err := n.getAPI(url, &d)

	return d, err
}
