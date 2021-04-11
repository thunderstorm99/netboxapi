package netboxapi

import "time"

// Tenant is a type that holds all info on a tenant from the netbox api
type Tenant struct {
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	ID           int         `json:"id"`
	Created      string      `json:"created"`
	Comments     string      `json:"comments"`
	Description  string      `json:"description"`
	Tags         []string    `json:"tags"`
	LastUpdated  time.Time   `json:"last_updated"`
	CustomFields interface{} `json:"custom_fields"`
	Group        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
		URL  string `json:"url"`
	} `json:"group,omitempty"`
}

// TenantGroup is a struct that holds info for a tenant group from the Netbox API
type TenantGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
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
