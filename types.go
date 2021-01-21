package netboxapi

// NetboxConnection is a type that holds all necessary info
type NetboxConnection struct {
	Token    string
	BaseURL  string
	Insecure bool
}
