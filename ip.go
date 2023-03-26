package netboxapi

import (
	"fmt"
	"time"
)

// IPAddress is the structure that holds info on IPs
type IPAddress struct {
	Address      string    `json:"address"`
	Created      string    `json:"created"`
	Description  string    `json:"description"`
	Family       int       `json:"family"`
	ID           int       `json:"id"`
	CustomFields any       `json:"custom_fields"`
	LastUpdated  time.Time `json:"last_updated"`
	NATInside    string    `json:"nat_inside"`
	NATOutside   string    `json:"nat_outside"`
	Role         string    `json:"role"`
	Tags         []string  `json:"tags"`
	VRF          string    `json:"vrf"`
	Status       struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	} `json:"status"`
	Tenant struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
		URL  string `json:"url"`
	} `json:"tenant"`
	Interface struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		URL    string `json:"url"`
		VM     string `json:"virtual_machine"`
		Device struct {
			DisplayName string `json:"display_name"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			URL         string `json:"url"`
		} `json:"device"`
	} `json:"interface"`
}

type IPconfig struct {
	TenantID int
	Family   int
}

// GetIPAdresses gets all IP addresses from this NetboxConnection
func (n *NetboxConnection) GetIPAdresses(config ...IPconfig) ([]IPAddress, error) {
	url := "/api/ipam/ip-addresses/?limit=500"

	// check if any config was provided
	if len(config) != 0 {
		// IP family
		if config[0].Family == 4 {
			url += "&family=4"
		} else if config[0].Family == 6 {
			url += "&family=6"
		}

		if config[0].TenantID != 0 {
			url += fmt.Sprintf("&tenant_id=%d", config[0].TenantID)
		}
	}

	var i []IPAddress
	err := n.getAPI(url, i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
