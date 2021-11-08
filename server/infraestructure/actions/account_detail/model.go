package account_detail

import "time"

// AccountDetail define o modelo de conta
type AccountDetail struct {
	ID              *int64     `sql:"id" conversorTag:"id"`
	CreatedAt       *time.Time `sql:"created_at" conversorTag:"created_at"`
	UpdatedAt       *time.Time `sql:"updated_at" conversorTag:"updated_at"`
	DeletedAt       *time.Time `sql:"deleted_at" conversorTag:"deleted_at"`
	Blocked         *bool      `sql:"blocked" conversorTag:"blocked"`
	Balance         *float64   `sql:"balance" conversorTag:"balance"`
	AccountID       *int64     `sql:"account_id" conversorTag:"account_id"`
	AccountPassword *string
}

type Access struct {
	AccountNumber  *int64  `sql:"account_number" conversorTag:"number"`
	VerifyingDigit *int64  `sql:"verifying_digit" conversorTag:"verifying_digit"`
	AgencyCode     *int64  `sql:"agency_code" conversorTag:"agency_code"`
	Password       *string `sql:"password" conversorTag:"password"`
}
