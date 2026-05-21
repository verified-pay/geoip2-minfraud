package model

type Score struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses/#score-body

	Disposition Disposition `json:"disposition"`

	FundsRemaining   float32 `json:"funds_remaining" example:"15.99"`
	QueriesRemaining uint64  `json:"queries_remaining" example:"1000"`
	ID               string  `json:"id" example:"e6e7e9f3-6f5a-4f9a-9b1a-3c2d1e0f4a5b"`

	RiskScore float32 `json:"risk_score" example:"0.01"`

	IPAddress ScoreIPAddress `json:"ip_address"`

	Warnings []Warning `json:"warnings,omitempty"`
}

// ScoreIPAddress contains the IP address risk for the Score response.
type ScoreIPAddress struct {
	Risk float32 `json:"risk" example:"0.01"`
}

// Disposition contains the disposition set by custom rules.
type Disposition struct {
	// Action is the action to take on the transaction as determined by your custom rules.
	// Possible values: "accept", "manual_review", "reject", "test"
	Action string `json:"action" example:"accept"`
	// Reason is the reason for the action.
	// Possible values: "custom_rule", "default"
	Reason string `json:"reason" example:"default"`
	// RuleLabel is the label of the custom rule that triggered the action.
	RuleLabel string `json:"rule_label" example:"custom rule"`
}
