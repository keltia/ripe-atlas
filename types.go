// types.go

package atlas

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
		Since string `json:"since"`
		ID    int    `json:"id"`
		Name  string `json:"name"`
	} `json:"status"`
	StatusSince int `json:"status_since"`
	Tags        []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"tags"`
	Type string `json:"type"`
}

// Measurement is what we are working with
type Measurement struct {
	Af                    int                    `json:"af"`
	CreationTime          int                    `json:"creation_time"`
	Description           string                 `json:"description"`
	Group                 string                 `json:"group"`
	GroupID               int                    `json:"group_id"`
	ID                    int                    `json:"id"`
	InWifiGroup           bool                   `json:"in_wifi_group"`
	Interval              int                    `json:"interval"`
	IsAllScheduled        bool                   `json:"is_all_scheduled"`
	IsOneoff              bool                   `json:"is_oneoff"`
	IsPublic              bool                   `json:"is_public"`
	PacketInterval        interface{}            `json:"packet_interval"`
	Packets               int                    `json:"packets"`
	ParticipantCount      int                    `json:"participant_count"`
	ParticipationRequests []ParticipationRequest `json:"participation_requests"`
	ProbesRequested       int                    `json:"probes_requested"`
	ProbesScheduled       int                    `json:"probes_scheduled"`
	ResolveOnProbe        bool                   `json:"resolve_on_probe"`
	ResolvedIPs           []string               `json:"resolved_ips"`
	Result                string                 `json:"result"`
	Size                  int                    `json:"size"`
	Spread                interface{}            `json:"spread"`
	StartTime             int                    `json:"start_time"`
	Status                struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"status"`
	StopTime  int    `json:"stop_time"`
	Target    string `json:"target"`
	TargetASN int    `json:"target_asn"`
	TargetIP  string `json:"target_ip"`
	Type      string `json:"type"`
}

// ParticipationRequest allow you to add or remove probes from a measurement that
// was already created
type ParticipationRequest struct {
	Action        string `json:"action"`
	CreatedAt     int    `json:"created_at"`
	ID            int    `json:"id"`
	Self          string `json:"self"`
	Measurement   string `json:"measurement"`
	MeasurementID int    `json:"measurement_id"`
	Requested     int    `json:"requested"`
	Type          string `json:"type"`
	Value         string `json:"value"`
	Logs          string `json:"logs"`
}

// Definition is used to create measurements
type Definition struct {
	// Required fields
	Description string
	Type string
	AF int
	// Required for all but "dns"
	Target string
}