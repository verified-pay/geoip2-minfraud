package model

import "time"

type Insights struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses#insights-body

	Disposition      Disposition `json:"disposition"`
	FundsRemaining   float32     `json:"funds_remaining" example:"15.99"`
	QueriesRemaining uint64      `json:"queries_remaining" example:"1000"`
	ID               string      `json:"id" example:"e6e7e9f3-6f5a-4f9a-9b1a-3c2d1e0f4a5b"`

	RiskScore float32 `json:"risk_score" example:"0.01"`

	IPAddress InsightsIPAddress `json:"ip_address"`

	BillingAddress  BillingAddress  `json:"billing_address"`
	BillingPhone    Phone           `json:"billing_phone"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
	ShippingPhone   Phone           `json:"shipping_phone"`

	CreditCard CreditCard `json:"credit_card"`
	Device     Device     `json:"device"`
	Email      Email      `json:"email"`

	Warnings []Warning `json:"warnings,omitempty"`
}

// InsightsIPAddress contains GeoIP2 and minFraud data about the IP address.
type InsightsIPAddress struct {
	Risk float32 `json:"risk" example:"0.01"`
	City struct {
		Confidence uint8             `json:"confidence" example:"25"`
		GeoNameID  uint              `json:"geoname_id" example:"5375480"`
		Names      map[string]string `json:"names"`
	} `json:"city"`
	Continent struct {
		Code      string            `json:"code" example:"NA"`
		GeoNameID uint              `json:"geoname_id" example:"6255149"`
		Names     map[string]string `json:"names"`
	} `json:"continent"`
	Country struct {
		GeoNameID         uint              `json:"geoname_id" example:"6252001"`
		IsoCode           string            `json:"iso_code" example:"US"`
		Names             map[string]string `json:"names"`
		Confidence        uint8             `json:"confidence" example:"99"`
		IsInEuropeanUnion bool              `json:"is_in_european_union" example:"false"`
	} `json:"country"`
	Location struct {
		AccuracyRadius    uint16    `json:"accuracy_radius" example:"20"`
		AverageIncome     float32   `json:"average_income" example:"50000"`
		LocalTime         time.Time `json:"local_time" example:"2026-05-21T04:34:44-07:00"`
		PopulationDensity float32   `json:"population_density" example:"1500"`
		Latitude          float64   `json:"latitude" example:"37.751"`
		Longitude         float64   `json:"longitude" example:"-97.822"`
		MetroCode         uint      `json:"metro_code" example:"807"`
		TimeZone          string    `json:"time_zone" example:"America/Los_Angeles"`
	} `json:"location"`
	Postal struct {
		Code       string `json:"code" example:"94103"`
		Confidence uint8  `json:"confidence" example:"40"`
	} `json:"postal"`
	RegisteredCountry struct {
		GeoNameID         uint              `json:"geoname_id" example:"6252001"`
		IsoCode           string            `json:"iso_code" example:"US"`
		Names             map[string]string `json:"names"`
		Confidence        uint8             `json:"confidence" example:"99"`
		IsInEuropeanUnion bool              `json:"is_in_european_union" example:"false"`
	} `json:"registered_country"`
	RepresentedCountry struct {
		GeoNameID         uint              `json:"geoname_id" example:"6252001"`
		IsInEuropeanUnion bool              `json:"is_in_european_union" example:"false"`
		IsoCode           string            `json:"iso_code" example:"US"`
		Names             map[string]string `json:"names"`
		Type              string            `json:"type" example:"military"`
	} `json:"represented_country"`
	Subdivisions []struct {
		Confidence uint8             `json:"confidence" example:"50"`
		GeoNameID  uint              `json:"geoname_id" example:"5332921"`
		IsoCode    string            `json:"iso_code" example:"CA"`
		Names      map[string]string `json:"names"`
	} `json:"subdivisions"`
	Traits      IPTraits       `json:"traits"`
	RiskReasons []IPRiskReason `json:"risk_reasons"`
}

// IPTraits contains data for the traits of the requested IP address.
type IPTraits struct {
	AutonomousSystemNumber       uint   `json:"autonomous_system_number" example:"13335"`
	AutonomousSystemOrganization string `json:"autonomous_system_organization" example:"CLOUDFLARENET"`
	Domain                       string `json:"domain" example:"cloudflare.com"`
	IPAddress                    string `json:"ip_address" example:"104.21.44.66"`

	IsAnonymous         bool `json:"is_anonymous" example:"false"`
	IsAnonymousProxy    bool `json:"is_anonymous_proxy" example:"false"`
	IsAnonymousVpn      bool `json:"is_anonymous_vpn" example:"false"`
	IsHostingProvider   bool `json:"is_hosting_provider" example:"true"`
	IsPublicProxy       bool `json:"is_public_proxy" example:"false"`
	IsResidentialProxy  bool `json:"is_residential_proxy" example:"false"`
	IsSatelliteProvider bool `json:"is_satellite_provider" example:"false"`
	IsTorExitNode       bool `json:"is_tor_exit_node" example:"false"`

	ISP               string  `json:"isp" example:"Cloudflare"`
	MobileCountryCode string  `json:"mobile_country_code" example:"310"`
	MobileNetworkCode string  `json:"mobile_network_code" example:"260"`
	Network           string  `json:"network" example:"104.21.44.0/24"`
	Organization      string  `json:"organization" example:"Cloudflare"`
	StaticIPScore     float64 `json:"static_ip_score" example:"1.3"`
	UserCount         uint    `json:"user_count" example:"1"`
	UserType          string  `json:"user_type" example:"hosting"`
}

// IPRiskReason provides a machine-readable code and human-readable explanation
// for the IP risk score.
type IPRiskReason struct {
	// Code is a machine-readable code identifying the reason.
	// Possible values: ANONYMOUS_IP, BILLING_POSTAL_VELOCITY, EMAIL_VELOCITY,
	// HIGH_RISK_DEVICE, HIGH_RISK_EMAIL, ISSUER_ID_NUMBER_VELOCITY, MINFRAUD_NETWORK_ACTIVITY
	Code string `json:"code" example:"ANONYMOUS_IP"`
	// Reason is a human-readable explanation of the reason.
	Reason string `json:"reason" example:"The IP address is registered to an anonymous network."`
}

// Phone contains information about the billing or shipping phone number.
type Phone struct {
	// Country is the two-character ISO 3166-1 country code for the country
	// associated with the phone number.
	Country string `json:"country,omitempty" example:"US"`
	// IsVoip is true if the phone number is a Voice over Internet Protocol (VoIP)
	// number allocated by a regulator.
	IsVoip bool `json:"is_voip,omitempty" example:"false"`
	// MatchesPostal is true if the phone number's prefix is commonly associated
	// with the postal code.
	MatchesPostal bool `json:"matches_postal,omitempty" example:"true"`
	// NetworkOperator is the name of the original network operator associated
	// with the phone number.
	NetworkOperator string `json:"network_operator,omitempty" example:"Verizon"`
	// NumberType is the type of phone number. Possible values: "fixed", "mobile".
	NumberType string `json:"number_type,omitempty" example:"mobile"`
}

// BillingAddress contains minFraud data related to the billing address.
type BillingAddress struct {
	DistanceToIPLocation int     `json:"distance_to_ip_location" example:"100"`
	IsInIPCountry        bool    `json:"is_in_ip_country" example:"true"`
	IsPostalInCity       bool    `json:"is_postal_in_city" example:"true"`
	Latitude             float64 `json:"latitude" example:"37.751"`
	Longitude            float64 `json:"longitude" example:"-97.822"`
}

// ShippingAddress contains minFraud data related to the shipping address.
type ShippingAddress struct {
	DistanceToBillingAddress int     `json:"distance_to_billing_address" example:"15"`
	DistanceToIPLocation     int     `json:"distance_to_ip_location" example:"100"`
	IsHighRisk               bool    `json:"is_high_risk" example:"false"`
	IsInIPCountry            bool    `json:"is_in_ip_country" example:"true"`
	IsPostalInCity           bool    `json:"is_postal_in_city" example:"true"`
	Latitude                 float64 `json:"latitude" example:"37.751"`
	Longitude                float64 `json:"longitude" example:"-97.822"`
}

// CreditCard contains minFraud data about the credit card used in the transaction.
type CreditCard struct {
	Brand                           string       `json:"brand" example:"Visa"`
	Country                         string       `json:"country" example:"US"`
	IsBusiness                      bool         `json:"is_business" example:"false"`
	IsIssuedInBillingAddressCountry bool         `json:"is_issued_in_billing_address_country" example:"true"`
	IsPrepaid                       bool         `json:"is_prepaid" example:"false"`
	IsVirtual                       bool         `json:"is_virtual" example:"false"`
	Issuer                          CreditIssuer `json:"issuer"`
	Type                            string       `json:"type" example:"credit"` // "credit", "debit", "charge"
}

// CreditIssuer contains information about the credit card issuer.
type CreditIssuer struct {
	MatchesProvidedName        bool   `json:"matches_provided_name" example:"true"`
	MatchesProvidedPhoneNumber bool   `json:"matches_provided_phone_number" example:"true"`
	Name                       string `json:"name" example:"Bank of America"`
	PhoneNumber                string `json:"phone_number" example:"+1-800-123-4567"`
}

// Device contains information about the device associated with the IP address.
type Device struct {
	// Confidence is a number from 0.01 to 99 representing the confidence that
	// the device_id refers to a unique device.
	Confidence float32 `json:"confidence" example:"99"`
	// ID is a UUID that MaxMind uses for the device.
	ID string `json:"id" example:"7835b099-d385-4e5b-969e-7df26181d73b"`
	// LastSeen is the date and time of the last sighting of the device (RFC 3339).
	LastSeen time.Time `json:"last_seen" example:"2026-05-20T08:30:00Z"`
	// LocalTime is the local date and time of the transaction in the device's time zone (RFC 3339).
	LocalTime time.Time `json:"local_time" example:"2026-05-21T04:34:44-07:00"`
}

// Email contains information about the email address.
type Email struct {
	Domain       EmailDomain `json:"domain"`
	FirstSeen    string      `json:"first_seen" example:"2017-01-02"` // YYYY-MM-DD
	IsDisposable bool        `json:"is_disposable" example:"false"`
	IsFree       bool        `json:"is_free" example:"true"`
	IsHighRisk   bool        `json:"is_high_risk" example:"false"`
}

// EmailDomain contains information about the email domain.
type EmailDomain struct {
	// Classification of the email domain.
	// Possible values: "business", "education", "government", "isp_email"
	Classification string `json:"classification,omitempty" example:"isp_email"`
	// FirstSeen is a date string (YYYY-MM-DD) when the domain was first seen.
	FirstSeen string `json:"first_seen" example:"2017-01-02"`
	// Risk is the risk associated with the email domain (0.01 to 99).
	Risk float32 `json:"risk,omitempty" example:"0.1"`
	// Visit contains information about an automated visit to the email domain.
	Visit EmailDomainVisit `json:"visit,omitempty"`
	// Volume indicates activity on the email domain across minFraud network
	// (sightings per million, 0.001 to 1,000,000).
	Volume float32 `json:"volume,omitempty" example:"1234.5"`
}

// EmailDomainVisit contains information about an automated visit to the email domain.
type EmailDomainVisit struct {
	// HasRedirect is true if the domain redirects to another URL.
	HasRedirect bool `json:"has_redirect,omitempty" example:"false"`
	// LastVisitedOn is the date when MaxMind last checked the domain (YYYY-MM-DD).
	LastVisitedOn string `json:"last_visited_on,omitempty" example:"2026-05-01"`
	// Status is the status of the domain based on the most recent automated check.
	// Possible values: "live", "dns_error", "network_error", "http_error", "parked", "pre_development"
	Status string `json:"status,omitempty" example:"live"`
}
