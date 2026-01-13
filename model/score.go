package model

type Score struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses/#score-body

	Disposition Disposition `json:"disposition"`

	FundsRemaining   float32 `json:"funds_remaining"`
	QueriesRemaining uint64  `json:"queries_remaining"`
	ID               string  `json:"id"`

	RiskScore float32 `json:"risk_score"`

	IPAddress ScoreIPAddress `json:"ip_address"`

	Warnings []Warning `json:"warnings,omitempty"`
}

// ScoreIPAddress contains the IP address risk for the Score response.
type ScoreIPAddress struct {
	Risk float32 `json:"risk"`
}

// Disposition contains the disposition set by custom rules.
type Disposition struct {
	// Action is the action to take on the transaction as determined by your custom rules.
	// Possible values: "accept", "manual_review", "reject", "test"
	Action string `json:"action"`
	// Reason is the reason for the action.
	// Possible values: "custom_rule", "default"
	Reason string `json:"reason"`
	// RuleLabel is the label of the custom rule that triggered the action.
	RuleLabel string `json:"rule_label"`
}
