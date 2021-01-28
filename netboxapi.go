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
func (n *NetboxConnection) GetTenants() interface{} {
	var tenantsNew []Tenant
	// resultToArray(n.GetAPIRaw("/api/tenancy/tenants"), &tenantsNew)
	return tenantsNew
}

// GetTenantGroups gets Tenant Groups from this NetboxConnection
func (n *NetboxConnection) GetTenantGroups() interface{} {
	return nil
	// return n.GetAPIRaw("/api/tenancy/tenant-groups/")
}

// GetVLANs gets VLANs from this NetboxConnection,
// if tenantID is specified it will only get VLANs for that specific tenant
func (n *NetboxConnection) GetVLANs(tenantID ...int) ([]VLAN, error) {
	// var vlans []interface{}
	url := "/api/ipam/vlans/"

	if tenantID != nil {
		// only get VLANs for specific tenant, change url
		url += fmt.Sprintf("?tenant_id=%d", tenantID[0])
	}

	// call url
	var a, err = n.GetAPIRaw(url)
	if err != nil {
		return nil, err
	}

	// cast result to VLAN array
	var vlans []VLAN
	err = resultToOtherStructure(a, &vlans)
	if err != nil {
		return nil, err
	}

	return vlans, nil
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

	a, err := n.GetAPIRaw(url)
	if err != nil {
		return []IPAddress{}, err
	}

	var i []IPAddress
	err = resultToOtherStructure(a, &i)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetAPIRaw performs a GET request onto any netbox API endpoint specified with url
func (n *NetboxConnection) GetAPIRaw(url string) (APIAnswer, error) {
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

		next, err := n.GetAPIRaw(nextURL)
		if err != nil {
			return APIAnswer{}, err
		}
		// append all results of next to these results
		answer.Results = append(answer.Results, next.Results...)
	}

	return answer, nil
}
