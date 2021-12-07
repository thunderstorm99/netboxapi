package netboxapi

type Cable struct {
	TermAID    int         `json:"termination_a_id"`
	TermAType  string      `json:"termination_a_type"`
	TermA      termination `json:"termination_a"`
	TermBID    int         `json:"termination_b_id"`
	TermB      termination `json:"termination_b"`
	Type       int         `json:"type"`
	Status     status      `json:"status"`
	Label      string      `json:"label"`
	Color      string      `json:"color"`
	Length     int         `json:"length"`
	LengthUnit int         `json:"length_unit"`
}

type status struct {
	Value bool   `json:"value"`
	Label string `json:"label"`
}
type termination struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Device NameID `json:"device"`
	Name   string `json:"name"`
	Cable  int    `json:"cable"`
	Status status `json:"connection_status"`
}

func (n *NetboxConnection) GetCables() ([]Cable, error) {
	var cables []Cable
	err := n.getAPI("/api/dcim/cables/", &cables)

	return cables, err
}
