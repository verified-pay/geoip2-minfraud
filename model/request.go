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
	Address          string `json:"address"`
	Address2         string `json:"address_2,omitempty"`
	City             string `json:"city"`
	Company          string `json:"company,omitempty"`
	Country          string `json:"country"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	PhoneCountryCode string `json:"phone_country_code,omitempty"` // 1
	PhoneNumber      string `json:"phone_number,omitempty"`       // 999-999-9999
	Postal           string `json:"postal"`
	Region           string `json:"region,omitempty"`
}

type ShippingReq struct {
	Address          string `json:"address"`
	Address2         string `json:"address_2,omitempty"`
	City             string `json:"city"`
	Company          string `json:"company,omitempty"`
	Country          string `json:"country"`
	DeliverySpeed    string `json:"delivery_speed,omitempty"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	PhoneCountryCode string `json:"phone_country_code,omitempty"` // 1
	PhoneNumber      string `json:"phone_number,omitempty"`       // 999-999-9999
	Postal           string `json:"postal"`
	Region           string `json:"region,omitempty"`
}

type CreditCardReq struct {
	AvsResult             string `json:"avs_result,omitempty"`
	BankName              string `json:"bank_name,omitempty"`
	BankPhoneCountryCode  string `json:"bank_phone_country_code,omitempty"` // 1
	BankPhoneNumber       string `json:"bank_phone_number,omitempty"`       // 999-999-9999
	Country               string `json:"country,omitempty"`
	CvvResult             string `json:"cvv_result,omitempty"`
	IssuerIdNumber        string `json:"issuer_id_number"`
	LastDigits            string `json:"last_digits"` // last 4
	Token                 string `json:"token,omitempty"`
	Was3dSecureSuccessful bool   `json:"was_3d_secure_successful,omitempty"`
}

// CustomInputsReq holds additional fields to be sent as JSON key-value pairs.
type CustomInputsReq map[string]interface{}

type AccountReq struct {
	UserID      string `json:"user_id,omitempty"`
	UsernameMD5 string `json:"username_md5,omitempty"`
}

type DeviceReq struct {
	AcceptLanguage string  `json:"accept_language,omitempty"`
	IPAddress      string  `json:"ip_address"`
	SessionAge     float32 `json:"session_age,omitempty"`
	SessionID      string  `json:"session_id,omitempty"`
	UserAgent      string  `json:"user_agent,omitempty"`
}

type EmailReq struct {
	Address string `json:"address"` // plain or MD5
	//AddressMD5 string `json:"address_md5"`
	Domain string `json:"domain,omitempty"`
}

type EventReq struct {
	ShopID        string    `json:"shop_id,omitempty"`
	Time          time.Time `json:"time"`
	TransactionID string    `json:"transaction_id"`
	Type          string    `json:"type"`
}

type OrderReq struct {
	AffiliateID    string  `json:"affiliate_id,omitempty"`
	Amount         float32 `json:"amount"`
	Currency       string  `json:"currency"`
	DiscountCode   string  `json:"discount_code,omitempty"`
	HasGiftMessage bool    `json:"has_gift_message"`
	IsGift         bool    `json:"is_gift"`
	ReferrerUri    string  `json:"referrer_uri,omitempty"`
	SubaffiliateID string  `json:"subaffiliate_id,omitempty"`
}

type PaymentReq struct {
	DeclineCode   string `json:"decline_code,omitempty"`
	Processor     string `json:"processor,omitempty"`
	WasAuthorized bool   `json:"was_authorized"`
}

type ShoppingCartReq struct {
	Category string  `json:"category"`
	ItemID   string  `json:"item_id"`
	Price    float32 `json:"price"`
	Quantity uint32  `json:"quantity"`
}
