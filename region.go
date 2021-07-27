package netboxapi

type Region struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Parent Short  `json:"parent"`
}

func (n *NetboxConnection) GetRegions() ([]Region, error) {
	var regions []Region
	err := n.getAPI("/api/dcim/regions/", &regions)

	return regions, err
}
