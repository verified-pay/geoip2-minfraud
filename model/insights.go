package model

import "time"

type Insights struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses#insights-body

	Disposition      Disposition `json:"disposition"`
	FundsRemaining   float32     `json:"funds_remaining"`
	QueriesRemaining uint64      `json:"queries_remaining"`
	ID               string      `json:"id"`

	RiskScore float32 `json:"risk_score"`

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
	Risk float32 `json:"risk"`
	City struct {
		Confidence uint8             `json:"confidence"`
		GeoNameID  uint              `json:"geoname_id"`
		Names      map[string]string `json:"names"`
	} `json:"city"`
	Continent struct {
		Code      string            `json:"code"`
		GeoNameID uint              `json:"geoname_id"`
		Names     map[string]string `json:"names"`
	} `json:"continent"`
	Country struct {
		GeoNameID         uint              `json:"geoname_id"`
		IsoCode           string            `json:"iso_code"`
		Names             map[string]string `json:"names"`
		Confidence        uint8             `json:"confidence"`
		IsInEuropeanUnion bool              `json:"is_in_european_union"`
	} `json:"country"`
	Location struct {
		AccuracyRadius    uint16    `json:"accuracy_radius"`
		AverageIncome     float32   `json:"average_income"`
		LocalTime         time.Time `json:"local_time"`
		PopulationDensity float32   `json:"population_density"`
		Latitude          float64   `json:"latitude"`
		Longitude         float64   `json:"longitude"`
		MetroCode         uint      `json:"metro_code"`
		TimeZone          string    `json:"time_zone"`
	} `json:"location"`
	Postal struct {
		Code       string `json:"code"`
		Confidence uint8  `json:"confidence"`
	} `json:"postal"`
	RegisteredCountry struct {
		GeoNameID         uint              `json:"geoname_id"`
		IsoCode           string            `json:"iso_code"`
		Names             map[string]string `json:"names"`
		Confidence        uint8             `json:"confidence"`
		IsInEuropeanUnion bool              `json:"is_in_european_union"`
	} `json:"registered_country"`
	RepresentedCountry struct {
		GeoNameID         uint              `json:"geoname_id"`
		IsInEuropeanUnion bool              `json:"is_in_european_union"`
		IsoCode           string            `json:"iso_code"`
		Names             map[string]string `json:"names"`
		Type              string            `json:"type"`
	} `json:"represented_country"`
	Subdivisions []struct {
		Confidence uint8             `json:"confidence"`
		GeoNameID  uint              `json:"geoname_id"`
		IsoCode    string            `json:"iso_code"`
		Names      map[string]string `json:"names"`
	} `json:"subdivisions"`
	Traits      IPTraits       `json:"traits"`
	RiskReasons []IPRiskReason `json:"risk_reasons"`
}

// IPTraits contains data for the traits of the requested IP address.
type IPTraits struct {
	AutonomousSystemNumber       uint   `json:"autonomous_system_number"`
	AutonomousSystemOrganization string `json:"autonomous_system_organization"`
	Domain                       string `json:"domain"`
	IPAddress                    string `json:"ip_address"`

	IsAnonymous         bool `json:"is_anonymous"`
	IsAnonymousProxy    bool `json:"is_anonymous_proxy"`
	IsAnonymousVpn      bool `json:"is_anonymous_vpn"`
	IsHostingProvider   bool `json:"is_hosting_provider"`
	IsPublicProxy       bool `json:"is_public_proxy"`
	IsResidentialProxy  bool `json:"is_residential_proxy"`
	IsSatelliteProvider bool `json:"is_satellite_provider"`
	IsTorExitNode       bool `json:"is_tor_exit_node"`

	ISP               string  `json:"isp"`
	MobileCountryCode string  `json:"mobile_country_code"`
	MobileNetworkCode string  `json:"mobile_network_code"`
	Network           string  `json:"network"`
	Organization      string  `json:"organization"`
	StaticIPScore     float64 `json:"static_ip_score"`
	UserCount         uint    `json:"user_count"`
	UserType          string  `json:"user_type"`
}

// IPRiskReason provides a machine-readable code and human-readable explanation
// for the IP risk score.
type IPRiskReason struct {
	// Code is a machine-readable code identifying the reason.
	// Possible values: ANONYMOUS_IP, BILLING_POSTAL_VELOCITY, EMAIL_VELOCITY,
	// HIGH_RISK_DEVICE, HIGH_RISK_EMAIL, ISSUER_ID_NUMBER_VELOCITY, MINFRAUD_NETWORK_ACTIVITY
	Code string `json:"code"`
	// Reason is a human-readable explanation of the reason.
	Reason string `json:"reason"`
}

// Phone contains information about the billing or shipping phone number.
type Phone struct {
	// Country is the two-character ISO 3166-1 country code for the country
	// associated with the phone number.
	Country string `json:"country,omitempty"`
	// IsVoip is true if the phone number is a Voice over Internet Protocol (VoIP)
	// number allocated by a regulator.
	IsVoip bool `json:"is_voip,omitempty"`
	// MatchesPostal is true if the phone number's prefix is commonly associated
	// with the postal code.
	MatchesPostal bool `json:"matches_postal,omitempty"`
	// NetworkOperator is the name of the original network operator associated
	// with the phone number.
	NetworkOperator string `json:"network_operator,omitempty"`
	// NumberType is the type of phone number. Possible values: "fixed", "mobile".
	NumberType string `json:"number_type,omitempty"`
}

// BillingAddress contains minFraud data related to the billing address.
type BillingAddress struct {
	DistanceToIPLocation int     `json:"distance_to_ip_location"`
	IsInIPCountry        bool    `json:"is_in_ip_country"`
	IsPostalInCity       bool    `json:"is_postal_in_city"`
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
}

// ShippingAddress contains minFraud data related to the shipping address.
type ShippingAddress struct {
	DistanceToBillingAddress int     `json:"distance_to_billing_address"`
	DistanceToIPLocation     int     `json:"distance_to_ip_location"`
	IsHighRisk               bool    `json:"is_high_risk"`
	IsInIPCountry            bool    `json:"is_in_ip_country"`
	IsPostalInCity           bool    `json:"is_postal_in_city"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
}

// CreditCard contains minFraud data about the credit card used in the transaction.
type CreditCard struct {
	Brand                           string       `json:"brand"`
	Country                         string       `json:"country"`
	IsBusiness                      bool         `json:"is_business"`
	IsIssuedInBillingAddressCountry bool         `json:"is_issued_in_billing_address_country"`
	IsPrepaid                       bool         `json:"is_prepaid"`
	IsVirtual                       bool         `json:"is_virtual"`
	Issuer                          CreditIssuer `json:"issuer"`
	Type                            string       `json:"type"` // "credit", "debit", "charge"
}

// CreditIssuer contains information about the credit card issuer.
type CreditIssuer struct {
	MatchesProvidedName        bool   `json:"matches_provided_name"`
	MatchesProvidedPhoneNumber bool   `json:"matches_provided_phone_number"`
	Name                       string `json:"name"`
	PhoneNumber                string `json:"phone_number"`
}

// Device contains information about the device associated with the IP address.
type Device struct {
	// Confidence is a number from 0.01 to 99 representing the confidence that
	// the device_id refers to a unique device.
	Confidence float32 `json:"confidence"`
	// ID is a UUID that MaxMind uses for the device.
	ID string `json:"id"`
	// LastSeen is the date and time of the last sighting of the device (RFC 3339).
	LastSeen time.Time `json:"last_seen"`
	// LocalTime is the local date and time of the transaction in the device's time zone (RFC 3339).
	LocalTime time.Time `json:"local_time"`
}

// Email contains information about the email address.
type Email struct {
	Domain       EmailDomain `json:"domain"`
	FirstSeen    string      `json:"first_seen"` // YYYY-MM-DD
	IsDisposable bool        `json:"is_disposable"`
	IsFree       bool        `json:"is_free"`
	IsHighRisk   bool        `json:"is_high_risk"`
}

// EmailDomain contains information about the email domain.
type EmailDomain struct {
	// Classification of the email domain.
	// Possible values: "business", "education", "government", "isp_email"
	Classification string `json:"classification,omitempty"`
	// FirstSeen is a date string (YYYY-MM-DD) when the domain was first seen.
	FirstSeen string `json:"first_seen"`
	// Risk is the risk associated with the email domain (0.01 to 99).
	Risk float32 `json:"risk,omitempty"`
	// Visit contains information about an automated visit to the email domain.
	Visit EmailDomainVisit `json:"visit,omitempty"`
	// Volume indicates activity on the email domain across minFraud network
	// (sightings per million, 0.001 to 1,000,000).
	Volume float32 `json:"volume,omitempty"`
}

// EmailDomainVisit contains information about an automated visit to the email domain.
type EmailDomainVisit struct {
	// HasRedirect is true if the domain redirects to another URL.
	HasRedirect bool `json:"has_redirect,omitempty"`
	// LastVisitedOn is the date when MaxMind last checked the domain (YYYY-MM-DD).
	LastVisitedOn string `json:"last_visited_on,omitempty"`
	// Status is the status of the domain based on the most recent automated check.
	// Possible values: "live", "dns_error", "network_error", "http_error", "parked", "pre_development"
	Status string `json:"status,omitempty"`
}
