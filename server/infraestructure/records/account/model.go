package account

import "time"

// Account define o modelo de conta
type Account struct {
	ID             *int64     `sql:"id" conversorTag:"id"`
	CreatedAt      *time.Time `sql:"created_at" conversorTag:"created_at"`
	UpdatedAt      *time.Time `sql:"updated_at" conversorTag:"updated_at"`
	DeletedAt      *time.Time `sql:"deleted_at" conversorTag:"deleted_at"`
	Number         *int64     `sql:"number" conversorTag:"number"`
	VerifyingDigit *int64     `sql:"verifying_digit" conversorTag:"verifying_digit"`
	Password       *string    `sql:"password" conversorTag:"password"`
	OwnerID        *int64     `sql:"account_owner_id" conversorTag:"account_owner_id"`
	AgencyID       *int64     `sql:"agency_id" conversorTag:"agency_id"`
}

// AccountPag define o modelo paginado de compra
type AccountPag struct {
	Data  []Account
	Next  *bool
	Total *int64
}
