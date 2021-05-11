package netboxapi

type NameID struct {
	Name string `json:"name,omitempty"`
	ID   int    `json:"id,omitempty"`
}
