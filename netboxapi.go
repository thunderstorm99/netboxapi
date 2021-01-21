package netboxapi

import (
	"net/http"

	"github.com/thunderstorm99/apihandler"
)

// NewNetboxConnection creates a new connection to a netbox via url and token
func NewNetboxConnection(url string, token string) NetboxConnection {
	return NetboxConnection{Token: token, BaseURL: url}
}

// GetTenants gets all tenants from this NetboxConnection
func (n *NetboxConnection) GetTenants() interface{} {
	return n.GetAPIRaw("/api/tenancy/tenants")
}

// GetTenantGroups gets Tenant Groups from this NetboxConnection
func (n *NetboxConnection) GetTenantGroups() interface{} {
	return n.GetAPIRaw("/api/tenancy/tenant-groups/")
}

// GetAPIRaw performs a GET request onto any netbox API endpoint specified with url
func (n *NetboxConnection) GetAPIRaw(url string) interface{} {
	a := apihandler.APICall{URL: n.BaseURL + url, Method: http.MethodGet, Header: map[string]string{"Authorization": n.Token}}
	return a.Exec()
}
