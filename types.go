// types.go

// This file contains the various types used by the API

package atlas

import (
	"log"
	"net/http"
)

// Client is the main struct holding state in an API client
type Client struct {
	config Config
	client *http.Client
	log    *log.Logger
	opts   map[string]string // Default, optional options
}

// Config is the main object when creating an API Client
type Config struct {
	endpoint     string
	APIKey       string
	DefaultProbe int
	AreaType     string
	AreaValue    string
	IsOneOff     bool
	PoolSize     int
	WantAF       string
	ProxyAuth    string
	Verbose      bool
	Tags         string
	Log          *log.Logger
}

// APIError is for errors returned by the RIPE API.
type APIError struct {
	Error struct {
		Status int    `json:"status"`
		Code   int    `json:"code"`
		Detail string `json:"detail"`
		Title  string `json:"title"`
		Errors []struct {
			Source struct {
				Pointer string
			} `json:"source"`
			Detail string
		} `json:"errors"`
	} `json:"error"`
}

// Key is holding the API key parameters
type Key struct {
	UUID      string `json:"uuid"`
	ValidFrom string `json:"valid_from"`
	ValidTo   string `json:"valid_to"`
	Enabled   bool
	IsActive  bool    `json:"is_active"`
	CreatedAt string  `json:"created_at"`
	Label     string  `json:"label"`
	Grants    []Grant `json:"grants"`
	Type      string  `json:"type"`
}

// Grant is the permission(s) associated with a key
type Grant struct {
	Permission string `json:"permission"`
	Target     struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"target"`
}

// Credits is holding credits data
type Credits struct {
	CurrentBalance            int    `json:"current_balance"`
	EstimatedDailyIncome      int    `json:"estimated_daily_income"`
	EstimatedDailyExpenditure int    `json:"estimated_daily_expenditure"`
	EstimatedDailyBalance     int    `json:"estimated_daily_balance"`
	CalculationTime           string `json:"calculation_time"`
	EstimatedRunoutSeconds    int    `json:"estimated_runout_seconds"`
	PastDayMeasurementResults int    `json:"past_day_measurement_results"`
	PastDayCreditsSpent       int    `json:"past_day_credits_spent"`
	IncomeItems               string `json:"income_items"`
	ExpenseItems              string `json:"expense_items"`
	Transactions              string `json:"transactions"`
}

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
	DestinationOptionSize interface{}            `json:"destination_option_size"`
	DontFragment          interface{}            `json:"dont_fragment"`
	DuplicateTimeout      interface{}            `json:"duplicate_timeout"`
	FirstHop              int                    `json:"first_hop"`
	Group                 string                 `json:"group"`
	GroupID               int                    `json:"group_id"`
	HopByHopOptionSize    interface{}            `json:"hop_by_hop_option_size"`
	ID                    int                    `json:"id"`
	InWifiGroup           bool                   `json:"in_wifi_group"`
	Interval              int                    `json:"interval"`
	IsAllScheduled        bool                   `json:"is_all_scheduled"`
	IsOneoff              bool                   `json:"is_oneoff"`
	IsPublic              bool                   `json:"is_public"`
	MaxHops               int                    `json:"max_hops"`
	PacketInterval        interface{}            `json:"packet_interval"`
	Packets               int                    `json:"packets"`
	Paris                 int                    `json:"paris"`
	ParticipantCount      int                    `json:"participant_count"`
	ParticipationRequests []ParticipationRequest `json:"participation_requests"`
	Port                  interface{}            `json:"port"`
	ProbesRequested       int                    `json:"probes_requested"`
	ProbesScheduled       int                    `json:"probes_scheduled"`
	Protocol              string                 `json:"protocol"`
	ResolveOnProbe        bool                   `json:"resolve_on_probe"`
	ResolvedIPs           []string               `json:"resolved_ips"`
	ResponseTimeout       int                    `json:"response_timeout"`
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
	CreatedAt     int    `json:"created_at,omitempty"`
	ID            int    `json:"id,omitempty"`
	Self          string `json:"self,omitempty"`
	Measurement   string `json:"measurement,omitempty"`
	MeasurementID int    `json:"measurement_id,omitempty"`
	Requested     int    `json:"requested,omitempty"`
	Type          string `json:"type,omitempty"`
	Value         string `json:"value,omitempty"`
	Logs          string `json:"logs,omitempty"`
}

var (
	// ProbeTypes should be obvious
	ProbeTypes = []string{"area", "country", "prefix", "asn", "probes", "msm"}
	// AreaTypes should also be obvious
	AreaTypes = []string{"WW", "West", "North-Central", "South-Central", "North-East", "South-East"}
)

// MeasurementRequest contains the different measurement to create/view
type MeasurementRequest struct {
	// see below for definition
	Definitions []Definition `json:"definitions"`

	// requested set of probes
	Probes []ProbeSet `json:"probes"`
	//
	BillTo       int  `json:"bill_to,omitempty"`
	IsOneoff     bool `json:"is_oneoff,omitempty"`
	SkipDNSCheck bool `json:"skip_dns_check,omitempty"`
	Times        int  `json:"times,omitempty"`
	StartTime    int  `json:"start_time,omitempty"`
	StopTime     int  `json:"stop_time,omitempty"`
}

// ProbeSet is a set of probes obviously
type ProbeSet struct {
	Requested   int    `json:"requested"` // number of probes
	Type        string `json:"type"`      // area, country, prefix, asn, probes, msm
	Value       string `json:"value"`     // can be numeric or string
	TagsInclude string `json:"tags_include,omitempty"`
	TagsExclude string `json:"tags_exclude,omitempty"`
}

// Definition is used to create measurements
type Definition struct {
	// Required fields
	Description string `json:"description"`
	Type        string `json:"type"`
	AF          int    `json:"af"`

	// Required for all but "dns"
	Target string `json:"target,omitempty"`

	GroupID        int    `json:"group_id,omitempty"`
	Group          string `json:"group,omitempty"`
	InWifiGroup    bool   `json:"in_wifi_group,omitempty"`
	Spread         int    `json:"spread,omitempty"`
	Packets        int    `json:"packets,omitempty"`
	PacketInterval int    `json:"packet_interval,omitempty"`
	Tags           string `json:"tags"`

	// Common parameters
	ExtraWait      int  `json:"extra_wait,omitempty"`
	IsOneoff       bool `json:"is_oneoff,omitempty"`
	IsPublic       bool `json:"is_public,omitempty"`
	ResolveOnProbe bool `json:"resolve_on_probe,omitempty"`

	// Default depends on type
	Interval int `json:"interval,omitempty"`

	// dns & traceroute parameters
	Protocol string `json:"protocol,omitempty"`

	// dns parameters
	QueryClass       string `json:"query_class,omitempty"`
	QueryType        string `json:"query_type,omitempty"`
	QueryArgument    string `json:"query_argument,omitempty"`
	Retry            int    `json:"retry,omitempty"`
	SetCDBit         bool   `json:"set_cd_bit,omitempty"`
	SetDOBit         bool   `json:"set_do_bit,omitempty"`
	SetNSIDBit       bool   `json:"set_nsid_bit,omitempty"`
	SetRDBit         bool   `json:"set_rd_bit,omitempty"`
	UDPPayloadSize   int    `json:"udp_payload_size,omitempty"`
	UseProbeResolver bool   `json:"use_probe_resolver"`

	// ping parameters
	//   none (see target)

	// traceroute parameters
	DestinationOptionSize int  `json:"destination_option_size,omitempty"`
	DontFragment          bool `json:"dont_fragment,omitempty"`
	DuplicateTimeout      int  `json:"duplicate_timeout,omitempty"`
	FirstHop              int  `json:"first_hop,omitempty"`
	HopByHopOptionSize    int  `json:"hop_by_hop_option_size,omitempty"`
	MaxHops               int  `json:"max_hops,omitempty"`
	Paris                 int  `json:"paris,omitempty"`

	// ntp parameters
	//   none (see target)

	// http parameters
	ExtendedTiming     bool   `json:"extended_timing,omitempty"`
	HeaderBytes        int    `json:"header_bytes,omitempty"`
	Method             string `json:"method,omitempty"`
	MoreExtendedTiming bool   `json:"more_extended_timing,omitempty"`
	Path               string `json:"path,omitempty"`
	QueryOptions       string `json:"query_options,omitempty"`
	UserAgent          string `json:"user_agent,omitempty"`
	Version            string `json:"version,omitempty"`

	// sslcert parameters
	//   none (see target)

	// sslcert & traceroute & http parameters
	Port int `json:"port,omitempty"`

	// ping & traceroute parameters
	Size int `json:"size,omitempty"`

	// wifi parameters
	AnonymousIdentity string `json:"anonymous_identity,omitempty"`
	Cert              string `json:"cert,omitempty"`
	EAP               string `json:"eap,omitempty"`
}
