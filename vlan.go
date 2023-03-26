package netboxapi

import "fmt"

// VLAN is the structure that holds info on VLANs
type VLAN struct {
	Created     string   `json:"created"`
	Description string   `json:"description"`
	Group       any      `json:"group"`
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Role        any      `json:"role"`
	Site        any      `json:"site"`
	Tags        []string `json:"tags"`
	Tenant      any      `json:"tenant"`
	VID         int      `json:"vid"`
	Status      struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	} `json:"status"`
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
