package netboxapi

type NameID struct {
	Name string `json:"name,omitempty"`
	ID   int    `json:"id,omitempty"`
}

type ValueLabel struct {
	Value int    `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
}

type Short struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
	Slug string `json:"slug,omitempty"`
}
