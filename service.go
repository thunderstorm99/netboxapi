package netboxapi

import (
	"fmt"
)

type Protocol struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type Service struct {
	Description string     `json:"description"`
	ID          int        `json:"id"`
	IPAddresses []DeviceIP `json:"ipaddresses"`
	Name        string     `json:"name"`
	Port        int        `json:"port"`
	Protocol    Protocol   `json:"protocol"`
	Device      NameID     `json:"device,omitempty"`
	VM          NameID     `json:"virtual_machine,omitempty"`
}

type ServiceConfig struct {
	DeviceID int
	Name     string
	VMID     int
}

func (n *NetboxConnection) GetServices(s ...ServiceConfig) []Service {
	url := "/api/ipam/services/"

	// check if any config was provided
	if len(s) != 0 {
		url += "?"
		// check if DeviceID was provided
		if s[0].DeviceID != 0 {
			url += fmt.Sprintf("device_id=%d&", s[0].DeviceID)
		}
		// check if VMID was provided
		if s[0].VMID != 0 {
			url += fmt.Sprintf("virtual_machine_id=%d&", s[0].VMID)
		}
		// check if Name was provided
		if s[0].Name != "" {
			url += fmt.Sprintf("name=%s", s[0].Name)
		}
	}

	var services []Service
	err := n.getAPI(url, &services)
	if err != nil {
		panic(err)
	}
	return services
}
