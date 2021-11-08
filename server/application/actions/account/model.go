package account

type Request struct {
	OwnerID  *int64  `json:"account_owner_id" binding:"required" conversorTag:"account_owner_id"`
	AgencyID *int64  `json:"agency_id" binding:"required" conversorTag:"agency_id"`
	Password *string `json:"password" binding:"required,gte=5" `
}

type Response struct {
	Number         *int64 `json:"account_number,omitempty" conversorTag:"number"`
	VerifyingDigit *int64 `json:"verifying_digit,omitempty" conversorTag:"verifying_digit"`
	AgencyCode     *int64 `json:"agency_code,omitempty"`
}

type ResponsePag struct {
	Data  []Response `json:"data,omitempty"`
	Next  *bool      `json:"next,omitempty"`
	Total *int64     `json:"total,omitempty"`
}
