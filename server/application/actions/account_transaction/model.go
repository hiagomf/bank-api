package account_transaction

type DepositRequest struct {
	AccountNumber  *int64   `json:"account_number" binding:"required" conversorTag:"account_number"`
	VerifyingDigit *int64   `json:"verifying_digit" binding:"required" conversorTag:"verifying_digit"`
	AgencyCode     *int64   `json:"agency_code" binding:"required" conversorTag:"agency_code"`
	Value          *float64 `json:"value" binding:"required,gt=0" conversorTag:"value"`
}

type TransferRequest struct {
	AccountNumber    *int64   `json:"account_number" binding:"required" conversorTag:"account_number"`
	VerifyingDigit   *int64   `json:"verifying_digit" binding:"required" conversorTag:"verifying_digit"`
	AgencyCode       *int64   `json:"agency_code" binding:"required" conversorTag:"agency_code"`
	Password         *string  `json:"password" binding:"required" conversorTag:"password"`
	Value            *float64 `json:"value" binding:"required,gt=0" conversorTag:"value"`
	ToAccountNumber  *int64   `json:"to_account_number" binding:"required" conversorTag:"to_account_number"`
	ToVerifyingDigit *int64   `json:"to_verifying_digit" binding:"required" conversorTag:"to_verifying_digit"`
	ToAgencyCode     *int64   `json:"to_agency_code" binding:"required" conversorTag:"to_agency_code"`
}
