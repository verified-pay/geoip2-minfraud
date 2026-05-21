package model

// The Enterprise struct corresponds to the data in the GeoIP2 Enterprise
// database.
type Enterprise struct {
	City struct {
		Confidence uint8             `maxminddb:"confidence" example:"25"`
		GeoNameID  uint              `maxminddb:"geoname_id" example:"5375480"`
		Names      map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Continent struct {
		Code      string            `maxminddb:"code" example:"NA"`
		GeoNameID uint              `maxminddb:"geoname_id" example:"6255149"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"continent"`
	Country struct {
		GeoNameID         uint              `maxminddb:"geoname_id" example:"6252001"`
		IsoCode           string            `maxminddb:"iso_code" example:"US"`
		Names             map[string]string `maxminddb:"names"`
		Confidence        uint8             `maxminddb:"confidence" example:"99"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union" example:"false"`
	} `maxminddb:"country"`
	Location struct {
		AccuracyRadius uint16  `maxminddb:"accuracy_radius" example:"20"`
		Latitude       float64 `maxminddb:"latitude" example:"37.751"`
		Longitude      float64 `maxminddb:"longitude" example:"-97.822"`
		MetroCode      uint    `maxminddb:"metro_code" example:"807"`
		TimeZone       string  `maxminddb:"time_zone" example:"America/Los_Angeles"`
	} `maxminddb:"location"`
	Postal struct {
		Code       string `maxminddb:"code" example:"94103"`
		Confidence uint8  `maxminddb:"confidence" example:"40"`
	} `maxminddb:"postal"`
	RegisteredCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id" example:"6252001"`
		IsoCode           string            `maxminddb:"iso_code" example:"US"`
		Names             map[string]string `maxminddb:"names"`
		Confidence        uint8             `maxminddb:"confidence" example:"99"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union" example:"false"`
	} `maxminddb:"registered_country"`
	RepresentedCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id" example:"6252001"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union" example:"false"`
		IsoCode           string            `maxminddb:"iso_code" example:"US"`
		Names             map[string]string `maxminddb:"names"`
		Type              string            `maxminddb:"type" example:"military"`
	} `maxminddb:"represented_country"`
	Subdivisions []struct {
		Confidence uint8             `maxminddb:"confidence" example:"50"`
		GeoNameID  uint              `maxminddb:"geoname_id" example:"5332921"`
		IsoCode    string            `maxminddb:"iso_code" example:"CA"`
		Names      map[string]string `maxminddb:"names"`
	} `maxminddb:"subdivisions"`
	Traits struct {
		AutonomousSystemNumber       uint    `maxminddb:"autonomous_system_number" example:"13335"`
		AutonomousSystemOrganization string  `maxminddb:"autonomous_system_organization" example:"CLOUDFLARENET"`
		ConnectionType               string  `maxminddb:"connection_type" example:"Corporate"`
		Domain                       string  `maxminddb:"domain" example:"cloudflare.com"`
		IsAnonymousProxy             bool    `maxminddb:"is_anonymous_proxy" example:"false"`
		IsLegitimateProxy            bool    `maxminddb:"is_legitimate_proxy" example:"false"`
		IsSatelliteProvider          bool    `maxminddb:"is_satellite_provider" example:"false"`
		ISP                          string  `maxminddb:"isp" example:"Cloudflare"`
		MobileCountryCode            string  `maxminddb:"mobile_country_code" example:"310"`
		MobileNetworkCode            string  `maxminddb:"mobile_network_code" example:"260"`
		Organization                 string  `maxminddb:"organization" example:"Cloudflare"`
		StaticIPScore                float64 `maxminddb:"static_ip_score" example:"1.3"`
		UserType                     string  `maxminddb:"user_type" example:"hosting"`
	} `maxminddb:"traits"`
}

// The City struct corresponds to the data in the GeoIP2/GeoLite2 City
// databases.
type City struct {
	City struct {
		GeoNameID uint              `maxminddb:"geoname_id" example:"5375480"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Continent struct {
		Code      string            `maxminddb:"code" example:"NA"`
		GeoNameID uint              `maxminddb:"geoname_id" example:"6255149"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"continent"`
	Country struct {
		GeoNameID         uint              `maxminddb:"geoname_id" example:"6252001"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union" example:"false"`
		IsoCode           string            `maxminddb:"iso_code" example:"US"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	Location struct {
		AccuracyRadius uint16  `maxminddb:"accuracy_radius" example:"20"`
		Latitude       float64 `maxminddb:"latitude" example:"37.751"`
		Longitude      float64 `maxminddb:"longitude" example:"-97.822"`
		MetroCode      uint    `maxminddb:"metro_code" example:"807"`
		TimeZone       string  `maxminddb:"time_zone" example:"America/Los_Angeles"`
	} `maxminddb:"location"`
	Postal struct {
		Code string `maxminddb:"code" example:"94103"`
	} `maxminddb:"postal"`
	RegisteredCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id" example:"6252001"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union" example:"false"`
		IsoCode           string            `maxminddb:"iso_code" example:"US"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"registered_country"`
	RepresentedCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id" example:"6252001"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union" example:"false"`
		IsoCode           string            `maxminddb:"iso_code" example:"US"`
		Names             map[string]string `maxminddb:"names"`
		Type              string            `maxminddb:"type" example:"military"`
	} `maxminddb:"represented_country"`
	Subdivisions []struct {
		GeoNameID uint              `maxminddb:"geoname_id" example:"5332921"`
		IsoCode   string            `maxminddb:"iso_code" example:"CA"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"subdivisions"`
	Traits struct {
		IsAnonymousProxy    bool `maxminddb:"is_anonymous_proxy" example:"false"`
		IsSatelliteProvider bool `maxminddb:"is_satellite_provider" example:"false"`
	} `maxminddb:"traits"`
}

// The Country struct corresponds to the data in the GeoIP2/GeoLite2
// Country databases.
type Country struct {
	Continent struct {
		Code      string            `maxminddb:"code" example:"NA"`
		GeoNameID uint              `maxminddb:"geoname_id" example:"6255149"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"continent"`
	Country struct {
		GeoNameID         uint              `maxminddb:"geoname_id" example:"6252001"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union" example:"false"`
		IsoCode           string            `maxminddb:"iso_code" example:"US"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	RegisteredCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id" example:"6252001"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union" example:"false"`
		IsoCode           string            `maxminddb:"iso_code" example:"US"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"registered_country"`
	RepresentedCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id" example:"6252001"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union" example:"false"`
		IsoCode           string            `maxminddb:"iso_code" example:"US"`
		Names             map[string]string `maxminddb:"names"`
		Type              string            `maxminddb:"type" example:"military"`
	} `maxminddb:"represented_country"`
	Traits struct {
		IsAnonymousProxy    bool `maxminddb:"is_anonymous_proxy" example:"false"`
		IsSatelliteProvider bool `maxminddb:"is_satellite_provider" example:"false"`
	} `maxminddb:"traits"`
}

// The AnonymousIP struct corresponds to the data in the GeoIP2
// Anonymous IP database.
type AnonymousIP struct {
	IsAnonymous        bool `maxminddb:"is_anonymous" example:"true"`
	IsAnonymousVPN     bool `maxminddb:"is_anonymous_vpn" example:"true"`
	IsHostingProvider  bool `maxminddb:"is_hosting_provider" example:"false"`
	IsPublicProxy      bool `maxminddb:"is_public_proxy" example:"false"`
	IsResidentialProxy bool `maxminddb:"is_residential_proxy" example:"false"`
	IsTorExitNode      bool `maxminddb:"is_tor_exit_node" example:"false"`
}

// The ASN struct corresponds to the data in the GeoLite2 ASN database.
type ASN struct {
	AutonomousSystemNumber       uint   `maxminddb:"autonomous_system_number" example:"13335"`
	AutonomousSystemOrganization string `maxminddb:"autonomous_system_organization" example:"CLOUDFLARENET"`
}

// The ConnectionType struct corresponds to the data in the GeoIP2
// Connection-Type database.
type ConnectionType struct {
	ConnectionType string `maxminddb:"connection_type" example:"Corporate"`
}

// The Domain struct corresponds to the data in the GeoIP2 Domain database.
type Domain struct {
	Domain string `maxminddb:"domain" example:"cloudflare.com"`
}

// The ISP struct corresponds to the data in the GeoIP2 ISP database.
type ISP struct {
	AutonomousSystemNumber       uint   `maxminddb:"autonomous_system_number" example:"13335"`
	AutonomousSystemOrganization string `maxminddb:"autonomous_system_organization" example:"CLOUDFLARENET"`
	ISP                          string `maxminddb:"isp" example:"Cloudflare"`
	MobileCountryCode            string `maxminddb:"mobile_country_code" example:"310"`
	MobileNetworkCode            string `maxminddb:"mobile_network_code" example:"260"`
	Organization                 string `maxminddb:"organization" example:"Cloudflare"`
}

type DatabaseType int

const (
	IsAnonymousIP DatabaseType = 1 << iota
	IsASN
	IsCity
	IsConnectionType
	IsCountry
	IsDomain
	IsEnterprise
	IsISP
)
