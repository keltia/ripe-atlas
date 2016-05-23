// types.go

package atlas

import "time"

// Probe is holding probe's data
type Probe struct {
	AddressV4      string `json:"address_v4"`
	AddressV6      string `json:"address_v6"`
	AsnV4          int    `json:"asn_v4"`
	AsnV6          int    `json:"asn_v6"`
	CountryCode    string `json:"country_code"`
	Description    string `json:"description"`
	FirstConnected int    `json:"first_connected"`
	Geometry       struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	ID            int    `json:"id"`
	IsAnchor      bool   `json:"is_anchor"`
	IsPublic      bool   `json:"is_public"`
	LastConnected int    `json:"last_connected"`
	PrefixV4      string `json:"prefix_v4"`
	PrefixV6      string `json:"prefix_v6"`
	Status        struct {
		Since time.Time `json:"since"`
		ID    int       `json:"id"`
		Name  string    `json:"name"`
	} `json:"status"`
	StatusSince int `json:"status_since"`
	Tags        []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"tags"`
	Type string `json:"type"`
}
