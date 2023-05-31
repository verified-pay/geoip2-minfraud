package model

type MaxmindWarning string

const (
	MaxmindWarningBillingCityNotFound     MaxmindWarning = "BILLING_CITY_NOT_FOUND"
	MaxmindWarningBillingCountryMissing   MaxmindWarning = "BILLING_COUNTRY_MISSING"
	MaxmindWarningBillingCountryNotFound  MaxmindWarning = "BILLING_COUNTRY_NOT_FOUND"
	MaxmindWarningBillingPostalNotFound   MaxmindWarning = "BILLING_POSTAL_NOT_FOUND"
	MaxmindWarningBillingRegionNotFound   MaxmindWarning = "BILLING_REGION_NOT_FOUND"
	MaxmindWarningEmailAddressUnusable    MaxmindWarning = "EMAIL_ADDRESS_UNUSABLE"
	MaxmindWarningInputInvalid            MaxmindWarning = "INPUT_INVALID"
	MaxmindWarningInputUnknown            MaxmindWarning = "INPUT_UNKNOWN"
	MaxmindWarningIpAddressInvalid        MaxmindWarning = "IP_ADDRESS_INVALID"
	MaxmindWarningIpAddressNotFound       MaxmindWarning = "IP_ADDRESS_NOT_FOUND"
	MaxmindWarningIpAddressReserved       MaxmindWarning = "IP_ADDRESS_RESERVED"
	MaxmindWarningShippingCityNotFound    MaxmindWarning = "SHIPPING_CITY_NOT_FOUND"
	MaxmindWarningShippingCountryMissing  MaxmindWarning = "SHIPPING_COUNTRY_MISSING"
	MaxmindWarningShippingCountryNotFound MaxmindWarning = "SHIPPING_COUNTRY_NOT_FOUND"
	MaxmindWarningShippingPostalNotFound  MaxmindWarning = "SHIPPING_POSTAL_NOT_FOUND"
	MaxmindWarningShippingRegionNotFound  MaxmindWarning = "SHIPPING_REGION_NOT_FOUND"
)

type Warning struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses/#schema--response--warning

	Code         MaxmindWarning `json:"code"`
	InputPointer string         `json:"input_pointer"`
	Warning      string         `json:"warning"`
}
