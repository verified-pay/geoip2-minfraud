package model

import (
	"encoding/json"
	"time"
)

type MinfraudReq struct {
	Billing      *BillingReq        `json:"billing"`
	Shipping     *ShippingReq       `json:"shipping"`
	CreditCard   *CreditCardReq     `json:"credit_card"`
	CustomInputs *CustomInputsReq   `json:"custom_inputs"`
	Account      *AccountReq        `json:"account"`
	Device       *DeviceReq         `json:"device"`
	Email        *EmailReq          `json:"email"`
	Event        *EventReq          `json:"event"`
	Order        *OrderReq          `json:"order"`
	Payment      *PaymentReq        `json:"payment"`
	ShoppingCart []*ShoppingCartReq `json:"shopping_cart"`
}

// GetID returns a unique deterministic ID for this request.
func (req *MinfraudReq) GetID() string {
	data, _ := json.Marshal(req)
	return getSha256(data)
}

type BillingReq struct {
	Address          string `json:"address" example:"101 Address Rd."`
	Address2         string `json:"address_2,omitempty" example:"Unit 5"`
	City             string `json:"city" example:"Mountain View"`
	Company          string `json:"company,omitempty" example:"MaxMind"`
	Country          string `json:"country" example:"US"`
	FirstName        string `json:"first_name" example:"Jane"`
	LastName         string `json:"last_name" example:"Doe"`
	PhoneCountryCode string `json:"phone_country_code,omitempty" example:"1"`      // 1
	PhoneNumber      string `json:"phone_number,omitempty" example:"123-456-7890"` // 999-999-9999
	Postal           string `json:"postal" example:"94043"`
	Region           string `json:"region,omitempty" example:"CA"`
}

type ShippingReq struct {
	Address          string `json:"address" example:"101 Address Rd."`
	Address2         string `json:"address_2,omitempty" example:"Unit 5"`
	City             string `json:"city" example:"Mountain View"`
	Company          string `json:"company,omitempty" example:"MaxMind"`
	Country          string `json:"country" example:"US"`
	DeliverySpeed    string `json:"delivery_speed,omitempty" example:"same_day_or_overnight"`
	FirstName        string `json:"first_name" example:"Jane"`
	LastName         string `json:"last_name" example:"Doe"`
	PhoneCountryCode string `json:"phone_country_code,omitempty" example:"1"`      // 1
	PhoneNumber      string `json:"phone_number,omitempty" example:"123-456-7890"` // 999-999-9999
	Postal           string `json:"postal" example:"94043"`
	Region           string `json:"region,omitempty" example:"CA"`
}

type CreditCardReq struct {
	AvsResult             string `json:"avs_result,omitempty" example:"Y"`
	BankName              string `json:"bank_name,omitempty" example:"Bank of America"`
	BankPhoneCountryCode  string `json:"bank_phone_country_code,omitempty" example:"1"`      // 1
	BankPhoneNumber       string `json:"bank_phone_number,omitempty" example:"123-456-7890"` // 999-999-9999
	Country               string `json:"country,omitempty" example:"US"`
	CvvResult             string `json:"cvv_result,omitempty" example:"N"`
	IssuerIdNumber        string `json:"issuer_id_number" example:"411111"`
	LastDigits            string `json:"last_digits" example:"7643"` // last 4
	Token                 string `json:"token,omitempty" example:"123456abcd1234"`
	Was3dSecureSuccessful bool   `json:"was_3d_secure_successful,omitempty" example:"true"`
}

// CustomInputsReq holds additional fields to be sent as JSON key-value pairs.
type CustomInputsReq map[string]interface{}

type AccountReq struct {
	UserID      string `json:"user_id,omitempty" example:"3132"`
	UsernameMD5 string `json:"username_md5,omitempty" example:"4f9726678c438914fa04bdb8c1a24088"`
}

type DeviceReq struct {
	AcceptLanguage string  `json:"accept_language,omitempty" example:"en-US"`
	IPAddress      string  `json:"ip_address" example:"104.21.44.66"`
	SessionAge     float32 `json:"session_age,omitempty" example:"3600"`
	SessionID      string  `json:"session_id,omitempty" example:"a333a4e127f880d8820e56a66f40717c"`
	UserAgent      string  `json:"user_agent,omitempty" example:"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)"`
}

type EmailReq struct {
	Address string `json:"address" example:"test@maxmind.com"` // plain or MD5
	//AddressMD5 string `json:"address_md5"`
	Domain string `json:"domain,omitempty" example:"maxmind.com"`
}

type EventReq struct {
	ShopID        string    `json:"shop_id,omitempty" example:"s2123"`
	Time          time.Time `json:"time" example:"2026-05-21T04:34:44Z"`
	TransactionID string    `json:"transaction_id" example:"txn3134133"`
	Type          string    `json:"type" example:"purchase"`
}

type OrderReq struct {
	AffiliateID    string  `json:"affiliate_id,omitempty" example:"af12"`
	Amount         float32 `json:"amount" example:"323.21"`
	Currency       string  `json:"currency" example:"USD"`
	DiscountCode   string  `json:"discount_code,omitempty" example:"FIRST10"`
	HasGiftMessage bool    `json:"has_gift_message" example:"false"`
	IsGift         bool    `json:"is_gift" example:"false"`
	ReferrerUri    string  `json:"referrer_uri,omitempty" example:"https://example.com/landing"`
	SubaffiliateID string  `json:"subaffiliate_id,omitempty" example:"saf42"`
}

type PaymentReq struct {
	DeclineCode   string `json:"decline_code,omitempty" example:"invalid_number"`
	Processor     string `json:"processor,omitempty" example:"stripe"`
	WasAuthorized bool   `json:"was_authorized" example:"true"`
}

type ShoppingCartReq struct {
	Category string  `json:"category" example:"pets"`
	ItemID   string  `json:"item_id" example:"ad23232"`
	Price    float32 `json:"price" example:"20.43"`
	Quantity uint32  `json:"quantity" example:"2"`
}
