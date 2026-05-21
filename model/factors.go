package model

// Factors represents the response from the minFraud Factors API.
// It contains all the data from the Insights response plus risk score reasons and subscores.
type Factors struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses/#factors-body

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

	// RiskScoreReasons contains risk score reasons for a given transaction that
	// change the risk score significantly. Risk score reasons are usually only
	// returned for medium to high risk transactions.
	RiskScoreReasons []RiskScoreReason `json:"risk_score_reasons,omitempty"`

	// Subscores contains scores for many of the individual risk factors that are
	// used to calculate the overall risk score.
	// Deprecated: use RiskScoreReasons instead.
	Subscores Subscores `json:"subscores"`

	Warnings []Warning `json:"warnings,omitempty"`
}

// RiskScoreReason describes a risk score multiplier and the reasons for it.
type RiskScoreReason struct {
	// Multiplier is the factor by which the risk score is increased (if > 1)
	// or decreased (if < 1). Multipliers > 1.5 or < 0.66 are considered significant.
	Multiplier float32 `json:"multiplier,omitempty" example:"1.5"`
	// Reasons is an array describing the reasons for the multiplier.
	Reasons []Reason `json:"reasons,omitempty"`
}

// Reason provides a machine-readable code and human-readable explanation for
// a risk score reason.
//
// Possible codes include:
//   - BROWSER_LANGUAGE - Riskiness of the browser user-agent and language
//   - BUSINESS_ACTIVITY - Riskiness of business activity
//   - COUNTRY - Riskiness of the country
//   - CUSTOMER_ID - Riskiness of a customer's activity
//   - EMAIL_DOMAIN - Riskiness of email domain
//   - EMAIL_DOMAIN_NEW - Riskiness of newly-sighted email domain
//   - EMAIL_ADDRESS_NEW - Riskiness of newly-sighted email address
//   - EMAIL_LOCAL_PART - Riskiness of the local part of the email address
//   - EMAIL_VELOCITY - Velocity on email
//   - ISSUER_ID_NUMBER_COUNTRY_MISMATCH - Country mismatch between IP, billing, shipping and IIN
//   - ISSUER_ID_NUMBER_ON_SHOP_ID - Risk of IIN for the shop ID
//   - ISSUER_ID_NUMBER_LAST_DIGITS_ACTIVITY - Riskiness of IIN and last digits activity
//   - ISSUER_ID_NUMBER_SHOP_ID_VELOCITY - Risk of recent IIN activity for shop ID
//   - INTRACOUNTRY_DISTANCE - Risk of distance between IP, billing, and shipping
//   - ANONYMOUS_IP - Risk due to IP being anonymous
//   - IP_BILLING_POSTAL_VELOCITY - Velocity of distinct billing postal code on IP
//   - IP_EMAIL_VELOCITY - Velocity of distinct email address on IP
//   - IP_HIGH_RISK_DEVICE - High-risk device sighted on IP
//   - IP_ISSUER_ID_NUMBER_VELOCITY - Velocity of distinct IIN on IP
//   - IP_ACTIVITY - Riskiness of IP based on minFraud network activity
//   - LANGUAGE - Riskiness of browser language
//   - MAX_RECENT_EMAIL - Riskiness of email based on past minFraud risk scores
//   - MAX_RECENT_PHONE - Riskiness of phone based on past minFraud risk scores
//   - MAX_RECENT_SHIP - Riskiness of ship address based on past minFraud risk scores
//   - MULTIPLE_CUSTOMER_ID_ON_EMAIL - Riskiness of email having many customer IDs
//   - ORDER_AMOUNT - Riskiness of the order amount
//   - ORG_DISTANCE_RISK - Risk of ISP and distance between billing address and IP
//   - PHONE - Riskiness of the phone number
//   - CART - Riskiness of shopping cart contents
//   - TIME_OF_DAY - Risk due to local time of day
//   - TRANSACTION_REPORT_EMAIL - Risk due to transaction reports on email
//   - TRANSACTION_REPORT_IP - Risk due to transaction reports on IP
//   - TRANSACTION_REPORT_PHONE - Risk due to transaction reports on phone
//   - TRANSACTION_REPORT_SHIP - Risk due to transaction reports on shipping address
//   - EMAIL_ACTIVITY - Riskiness of email based on minFraud network activity
//   - PHONE_ACTIVITY - Riskiness of phone based on minFraud network activity
//   - SHIP_ACTIVITY - Riskiness of ship address based on minFraud network activity
type Reason struct {
	// Code is a machine-readable code identifying the reason.
	Code string `json:"code" example:"EMAIL_DOMAIN"`
	// Reason is a human-readable explanation of the reason.
	Reason string `json:"reason" example:"Riskiness of email domain"`
}

// Subscores contains scores for individual risk factors used to calculate the overall risk score.
// Deprecated: use RiskScoreReason instead.
type Subscores struct {
	// AVSResult is the risk associated with the AVS result (0.01 to 99).
	AVSResult *float32 `json:"avs_result,omitempty" example:"0.5"`
	// BillingAddress is the risk associated with the billing address (0.01 to 99).
	BillingAddress *float32 `json:"billing_address,omitempty" example:"0.5"`
	// BillingAddressDistanceToIPLocation is the risk associated with the distance
	// between billing address and IP location (0.01 to 99).
	BillingAddressDistanceToIPLocation *float32 `json:"billing_address_distance_to_ip_location,omitempty" example:"0.5"`
	// Browser is the risk associated with browser attributes (0.01 to 99).
	Browser *float32 `json:"browser,omitempty" example:"0.5"`
	// Chargeback is individualized risk of chargeback for the IP address (0.01 to 99).
	Chargeback *float32 `json:"chargeback,omitempty" example:"0.5"`
	// Country is the risk associated with the country (0.01 to 99).
	Country *float32 `json:"country,omitempty" example:"0.5"`
	// CountryMismatch is the risk associated with the combination of IP country,
	// card issuer country, billing country, and shipping country (0.01 to 99).
	CountryMismatch *float32 `json:"country_mismatch,omitempty" example:"0.5"`
	// CVVResult is the risk associated with the CVV result (0.01 to 99).
	CVVResult *float32 `json:"cvv_result,omitempty" example:"0.5"`
	// Device is the risk associated with the device (0.01 to 99).
	Device *float32 `json:"device,omitempty" example:"0.5"`
	// EmailAddress is the risk associated with the email address (0.01 to 99).
	EmailAddress *float32 `json:"email_address,omitempty" example:"0.5"`
	// EmailDomain is the general risk associated with the email domain (0.01 to 99).
	EmailDomain *float32 `json:"email_domain,omitempty" example:"0.5"`
	// EmailLocalPart is the risk associated with the email local part (0.01 to 99).
	EmailLocalPart *float32 `json:"email_local_part,omitempty" example:"0.5"`
	// IssuerIDNumber is the risk associated with the IIN given billing location
	// and usage history (0.01 to 99).
	IssuerIDNumber *float32 `json:"issuer_id_number,omitempty" example:"0.5"`
	// OrderAmount is the risk associated with the order amount (0.01 to 99).
	OrderAmount *float32 `json:"order_amount,omitempty" example:"0.5"`
	// PhoneNumber is the risk associated with the phone number (0.01 to 99).
	PhoneNumber *float32 `json:"phone_number,omitempty" example:"0.5"`
	// ShippingAddress is the risk associated with the shipping address (0.01 to 99).
	ShippingAddress *float32 `json:"shipping_address,omitempty" example:"0.5"`
	// ShippingAddressDistanceToIPLocation is the risk associated with the distance
	// between shipping address and IP location (0.01 to 99).
	ShippingAddressDistanceToIPLocation *float32 `json:"shipping_address_distance_to_ip_location,omitempty" example:"0.5"`
	// TimeOfDay is the risk associated with the local time of day (0.01 to 99).
	TimeOfDay *float32 `json:"time_of_day,omitempty" example:"0.5"`
}
