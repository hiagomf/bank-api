package agency

import (
	"time"

	"github.com/hiagomf/bank-api/server/application/records/bank"
)

type Response struct {
	ID          *int64     `json:"id,omitempty" conversorTag:"id"`
	CreatedAt   *time.Time `json:"created_at,omitempty" conversorTag:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" conversorTag:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" conversorTag:"deleted_at"`
	Code        *int64     `json:"code,omitempty" conversorTag:"code"`
	MainAgency  *bool      `json:"main_agency,omitempty" conversorTag:"main_agency"`
	ZipCode     *string    `json:"zip_code,omitempty" conversorTag:"zip_code"`
	PublicPlace *string    `json:"public_place,omitempty" conversorTag:"public_place"`
	Number      *string    `json:"number,omitempty" conversorTag:"number"`
	Complement  *string    `json:"complement,omitempty" conversorTag:"complement"`
	District    *string    `json:"district,omitempty" conversorTag:"district"`
	City        *string    `json:"city,omitempty" conversorTag:"city"`
	State       *string    `json:"state,omitempty" conversorTag:"state"`
	Country     *string    `json:"country,omitempty" conversorTag:"country"`
	BankID      *int64     `json:"bank_id,omitempty" conversorTag:"bank_id"`
}

type ResponsePag struct {
	Bank     *bank.Response `json:"bank,omitempty"`
	Agencies []Response     `json:"agencies,omitempty"`
	Next     *bool          `json:"next,omitempty"`
	Total    *int64         `json:"total,omitempty"`
}
