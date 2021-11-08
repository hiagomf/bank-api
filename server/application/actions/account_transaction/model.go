package account_transaction

type DepositRequest struct {
	AccountNumber  *int64   `json:"account_number" binding:"required" conversorTag:"number"`
	VerifyingDigit *int64   `json:"verifying_digit" binding:"required" conversorTag:"verifying_digit"`
	AgencyCode     *int64   `json:"agency_code" binding:"required" conversorTag:"agency_code"`
	Value          *float64 `json:"value" binding:"required,gt=0" conversorTag:"value"`
}
