package account_detail

import "time"

type Request struct {
	AccountNumber  *int64  `json:"account_number" binding:"required" conversorTag:"number"`
	VerifyingDigit *int64  `json:"verifying_digit" binding:"required" conversorTag:"verifying_digit"`
	AgencyCode     *int64  `json:"agency_code" binding:"required" conversorTag:"agency_code"`
	Password       *string `json:"password" binding:"required" conversorTag:"password"`
}

type Response struct {
	ID        *int64     `json:"id,omitempty" conversorTag:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty" conversorTag:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" conversorTag:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" conversorTag:"deleted_at"`
	Blocked   *bool      `json:"blocked,omitempty" conversorTag:"blocked"`
	Balance   *float64   `json:"balance,omitempty" conversorTag:"balance"`
	AccountID *int64     `json:"account_id,omitempty" conversorTag:"account_id"`
}
