package netboxapi

type Rack struct {
	ID              int        `json:"id,omitempty"`
	Name            string     `json:"name,omitempty"`
	FacilityID      string     `json:"facility_id,omitempty"`
	DisplayName     string     `json:"display_name,omitempty"`
	Site            Short      `json:"site,omitempty"`
	Group           Short      `json:"group,omitempty"`
	Tenant          Short      `json:"tenant,omitempty"`
	Status          ValueLabel `json:"status,omitempty"`
	Role            int        `json:"role,omitempty"`
	Serial          string     `json:"string,omitempty"`
	AssetTag        string     `json:"asset_tag,omitempty"`
	Type            ValueLabel `json:"type,omitempty"`
	Width           ValueLabel `json:"width,omitempty"`
	HeightUnits     int        `json:"u_height,omitempty"`
	DescendingUnits bool       `json:"desc_units,omitempty"`
	OuterWidth      int        `json:"outer_width,omitempty"`
	OuterDepth      int        `json:"outer_depth,omitempty"`
	OuterUnit       ValueLabel `json:"outer_unit,omitempty"`
	Comments        string     `json:"comments,omitempty"`
	Tags            []string   `json:"tags,omitempty"`
	CustomFields    any        `json:"custom_fields,omitempty"`
	Created         string     `json:"created,omitempty"`
	LastUpdated     string     `json:"last_updated,omitempty"`
}

func (n *NetboxConnection) GetRacks() ([]Rack, error) {
	var racks []Rack
	err := n.getAPI("/api/dcim/racks/", &racks)

	return racks, err
}
