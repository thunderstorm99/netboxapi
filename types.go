package main

import (
	"time"
)

// NetboxConnection is a type that holds all necessary info
type NetboxConnection struct {
	Token    string
	BaseURL  string
	Insecure bool
}

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

// APIAnswer is the answer from the netbox API
type APIAnswer struct {
	Count   int           `json:"count"`
	Next    string        `json:"next"`
	Results []interface{} `json:"results"`
}

// VLAN is the structure that holds info on VLANs
type VLAN struct {
	Created     string      `json:"created"`
	Description string      `json:"description"`
	Group       interface{} `json:"group"`
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Role        interface{} `json:"role"`
	Site        interface{} `json:"site"`
	Tags        []string    `json:"tags"`
	Tenant      interface{} `json:"tenant"`
	VID         int         `json:"vid"`
	Status      struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	} `json:"status"`
}

// IPAddress is the structure that holds info on IPs
type IPAddress struct {
	Address      string      `json:"address"`
	Created      string      `json:"created"`
	Description  string      `json:"description"`
	Family       int         `json:"family"`
	ID           int         `json:"id"`
	CustomFields interface{} `json:"custom_fields"`
	LastUpdated  time.Time   `json:"last_updated"`
	NATInside    string      `json:"nat_inside"`
	NATOutside   string      `json:"nat_outside"`
	Role         string      `json:"role"`
	Tags         []string    `json:"tags"`
	VRF          string      `json:"vrf"`
	Status       struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	} `json:"status"`
	Tenant struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
		URL  string `json:"url"`
	} `json:"tenant"`
	Interface struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		URL    string `json:"url"`
		VM     string `json:"virtual_machine"`
		Device struct {
			DisplayName string `json:"display_name"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			URL         string `json:"url"`
		} `json:"device"`
	} `json:"interface"`
}

type ipconfig struct {
	TenantID int
	Family   int
}

// TenantGroup is a struct that holds info for a tenant group from the Netbox API
type TenantGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
