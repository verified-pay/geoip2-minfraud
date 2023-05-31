package model

import "time"

type Insights struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses#insights-body

	Disposition struct {
		Action    string `json:"action"`
		Reason    string `json:"reason"`
		RuleLabel string `json:"rule_label"`
	} `json:"disposition"`
	FundsRemaining   float32 `json:"funds_remaining"`
	QueriesRemaining uint64  `json:"queries_remaining"`
	ID               string  `json:"id"`

	RiskScore float32 `json:"risk_score"`

	IPAddress struct {
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
		Type         string `json:"type"` // TODO enum? also string official PHP lib
		Subdivisions []struct {
			Confidence uint8             `json:"confidence"`
			GeoNameID  uint              `json:"geoname_id"`
			IsoCode    string            `json:"iso_code"`
			Names      map[string]string `json:"names"`
		} `json:"subdivisions"`

		Traits struct {
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
			//IsLegitimateProxy bool    `json:"is_legitimate_proxy"`

			ISP               string  `json:"isp"`
			MobileCountryCode string  `json:"mobile_country_code"`
			MobileNetworkCode string  `json:"mobile_network_code"`
			Network           string  `json:"network"`
			Organization      string  `json:"organization"`
			StaticIPScore     float64 `json:"static_ip_score"`
			UserCount         uint    `json:"user_count"`
			UserType          string  `json:"user_type"`
			//ConnectionType string `json:"connection_type"`
		} `json:"traits"`
	} `json:"ip_address"`

	RiskReasons []RiskReason `json:"risk_reasons"`

	BillingAddress struct {
		DistanceToIPLocation float32 `json:"distance_to_ip_location"`
		IsInIPCountry        bool    `json:"is_in_ip_country"`
		IsPostalInCity       bool    `json:"is_postal_in_city"`
		Latitude             float64 `json:"latitude"`
		Longitude            float64 `json:"longitude"`
	} `json:"billing_address"`

	ShippingAddress struct {
		DistanceToBillingAddress float32 `json:"distance_to_billing_address"`
		DistanceToIPLocation     float32 `json:"distance_to_ip_location"`
		IsHighRisk               bool    `json:"is_high_risk"`
		IsInIPCountry            bool    `json:"is_in_ip_country"`
		IsPostalInCity           bool    `json:"is_postal_in_city"`
		Latitude                 float64 `json:"latitude"`
		Longitude                float64 `json:"longitude"`
	} `json:"shipping_address"`

	CreditCard struct {
		Brand                           string `json:"brand"`
		Country                         string `json:"country"`
		IsBusiness                      bool   `json:"is_business"`
		IsIssuedInBillingAddressCountry bool   `json:"is_issued_in_billing_address_country"`
		IsPrepaid                       bool   `json:"is_prepaid"`
		IsVirtual                       bool   `json:"is_virtual"`
		Issuer                          struct {
			MatchesProvidedName        bool   `json:"matches_provided_name"`
			MatchesProvidedPhoneNumber bool   `json:"matches_provided_phone_number"`
			Name                       string `json:"name"`
			PhoneNumber                string `json:"phone_number"`
		}
		Type string `json:"type"` // credit|debit|prepaid ... ?
	} `json:"credit_card"`

	Device struct {
		Confidence uint8     `json:"confidence"`
		ID         string    `json:"id"`
		LastSeen   time.Time `json:"last_seen"`
		LocalTime  time.Time `json:"local_time"`
	} `json:"device"`

	Email struct {
		Domain struct {
			FirstSeen string `json:"first_seen"` // YYYY-MM-DD
		} `json:"domain"`
		FirstSeen    string `json:"first_seen"` // YYYY-MM-DD
		IsDisposable bool   `json:"is_disposable"`
		IsFree       bool   `json:"is_free"`
		IsHighRisk   bool   `json:"is_high_risk"`
	} `json:"email"`

	Warnings []Warning `json:"warnings,omitempty"`
}

type RiskReasonCode string

const (
	RiskReasonCodeAnonymousIP        RiskReasonCode = "ANONYMOUS_IP"
	RiskReasonCodeBillPostalVelocity RiskReasonCode = "BILLING_POSTAL_VELOCITY"
	RiskReasonCodeEmailVelocity      RiskReasonCode = "EMAIL_VELOCITY"
	RiskReasonCodeHighRiskDevice     RiskReasonCode = "HIGH_RISK_DEVICE"
	RiskReasonCodeHighRiskEmail      RiskReasonCode = "HIGH_RISK_EMAIL"
	RiskReasonCodeIssuerIDVelocity   RiskReasonCode = "ISSUER_ID_NUMBER_VELOCITY"
	RiskReasonCodeNetworkActivity    RiskReasonCode = "MINFRAUD_NETWORK_ACTIVITY"
)

type RiskReason struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses/#schema--response--ip-address--risk-reason
	Code   RiskReasonCode `json:"code"`
	Reason string         `json:"reason"`
}
