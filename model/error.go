package model

type MaxmindErrorCode string

const (
	MaxmindErrUnknown            MaxmindErrorCode = ""
	MaxmindErrJsonInvalid        MaxmindErrorCode = "JSON_INVALID"
	MaxmindErrRequestInvalid     MaxmindErrorCode = "REQUEST_INVALID"
	MaxmindErrAuthInvalid        MaxmindErrorCode = "AUTHORIZATION_INVALID"
	MaxmindErrLicenseKeyRequired MaxmindErrorCode = "LICENSE_KEY_REQUIRED"
	MaxmindErrAccountIDRequired  MaxmindErrorCode = "ACCOUNT_ID_REQUIRED"
	MaxmindErrInsufficientFunds  MaxmindErrorCode = "INSUFFICIENT_FUNDS"
	MaxmindErrPermissionRequired MaxmindErrorCode = "PERMISSION_REQUIRED"
)

// Error JSON is returned on HTTP 4xx and 5xx responses.
type Error struct {
	// https://dev.maxmind.com/minfraud/api-documentation/responses/#error-body

	Code  MaxmindErrorCode `json:"code"`
	Error string           `json:"error"`
}
