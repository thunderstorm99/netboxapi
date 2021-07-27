package netboxapi

type Site struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	Slug            string        `json:"slug"`
	Status          ValueLabel    `json:"status"`
	Region          Short         `json:"region"`
	Tenant          Short         `json:"tenant"`
	Facility        string        `json:"facility"`
	ASN             int           `json:"asn"`
	TimeZone        string        `json:"time_zone"`
	Description     string        `json:"description"`
	PhysicalAddress string        `json:"physical_address"`
	ShippingAddress string        `json:"shipping_address"`
	Latitude        float64       `json:"latitude"`
	Longitude       float64       `json:"longitude"`
	ContactName     string        `json:"contact_name"`
	ContactPhone    string        `json:"contact_phone"`
	ContactEmail    string        `json:"contact_email"`
	Comments        string        `json:"comments"`
	Tags            []string      `json:"tags"`
	CustomFields    []interface{} `json:"custom_fields"`
	Created         string        `json:"created"`
	LastUpdated     string        `json:"last_updated"`
	CountPrefixes   int           `json:"count_prefixes"`
	CountVlans      int           `json:"count_vlans"`
	CountRacks      int           `json:"count_racks"`
	CountDevices    int           `json:"count_devices"`
	CountCircuits   int           `json:"count_circuits"`
}
