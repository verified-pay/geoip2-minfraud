package model

type Score struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses/#score-body

	FundsRemaining   float32 `json:"funds_remaining"`
	QueriesRemaining uint64  `json:"queries_remaining"`
	ID               string  `json:"id"`

	RiskScore float32 `json:"risk_score"`

	IPAddress struct {
		Risk float32 `json:"risk"`
	} `json:"ip_address"`

	Continent struct {
		Code      string            `json:"code"`
		GeoNameID uint              `json:"geoname_id"`
		Names     map[string]string `json:"names"`
	} `json:"continent"`

	Warnings []Warning `json:"warnings,omitempty"`
}
