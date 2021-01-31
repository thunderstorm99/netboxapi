package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/thunderstorm99/apihandler"
)

// NewNetboxConnection creates a new connection to a netbox via url and token
func NewNetboxConnection(url string, token string) NetboxConnection {
	return NetboxConnection{Token: token, BaseURL: url}
}

// GetTenants gets all tenants from this NetboxConnection
func (n *NetboxConnection) GetTenants() ([]Tenant, error) {
	url := "/api/tenancy/tenants"

	var tenants []Tenant
	err := n.getAPI(url, &tenants)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}

// GetTenantGroups gets Tenant Groups from this NetboxConnection
func (n *NetboxConnection) GetTenantGroups() ([]TenantGroup, error) {
	url := "/api/tenancy/tenant-groups/"

	var tenantGroups []TenantGroup
	err := n.getAPI(url, &tenantGroups)
	if err != nil {
		return nil, err
	}

	return tenantGroups, nil
}

// GetVLANs gets VLANs from this NetboxConnection,
// if tenantID is specified it will only get VLANs for that specific tenant
func (n *NetboxConnection) GetVLANs(tenantID ...int) ([]VLAN, error) {
	url := "/api/ipam/vlans/"

	if tenantID != nil {
		// only get VLANs for specific tenant, change url
		url += fmt.Sprintf("?tenant_id=%d", tenantID[0])
	}

	var v []VLAN
	err := n.getAPI(url, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetIPAdresses gets all IP addresses from this NetboxConnection
func (n *NetboxConnection) GetIPAdresses(config ...ipconfig) ([]IPAddress, error) {
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

// getAPI ...
func (n *NetboxConnection) getAPI(url string, output interface{}) error {
	a, err := n.getAPIRaw(url)
	if err != nil {
		return err
	}

	err = resultToOtherStructure(a, &output)
	if err != nil {
		return err
	}

	return nil
}

// getAPIRaw performs a GET request onto any netbox API endpoint specified with url
func (n *NetboxConnection) getAPIRaw(url string) (APIAnswer, error) {
	api := apihandler.APICall{URL: n.BaseURL + url, Method: http.MethodGet, Header: map[string]string{"Authorization": "Token " + n.Token}, Insecure: n.Insecure}

	// create new variable that holds the answer of the API
	var answer APIAnswer

	// calling API
	log.Println("calling URL", api.URL)
	err := api.Exec(&answer)
	if err != nil {
		return APIAnswer{}, err
	}

	// check for pagination
	if answer.Next != "" {
		// check if BaseURL is https
		if n.BaseURL[0:8] == "https://" {
			// check if Next URL is http (which it should not be), but if so upgrade to https
			if answer.Next[0:7] == "http://" {
				// upgrade http to https
				answer.Next = strings.Replace(answer.Next, "http://", "https://", 1)
			}
		}

		// trim the base off of the next url
		nextURL := strings.TrimPrefix(answer.Next, n.BaseURL)

		next, err := n.getAPIRaw(nextURL)
		if err != nil {
			return APIAnswer{}, err
		}
		// append all results of next to these results
		answer.Results = append(answer.Results, next.Results...)
	}

	return answer, nil
}
