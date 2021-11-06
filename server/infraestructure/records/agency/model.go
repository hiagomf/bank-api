package agency

import "time"

type Agency struct {
	ID          *int64     `sql:"id" conversorTag:"id"`
	CreatedAt   *time.Time `sql:"created_at" conversorTag:"created_at"`
	UpdatedAt   *time.Time `sql:"updated_at" conversorTag:"updated_at"`
	DeletedAt   *time.Time `sql:"deleted_at" conversorTag:"deleted_at"`
	Code        *int64     `sql:"code" conversorTag:"code"`
	MainAgency  *bool      `sql:"main_agency" conversorTag:"main_agency"`
	ZipCode     *string    `sql:"zip_code" conversorTag:"zip_code"`
	PublicPlace *string    `sql:"public_place" conversorTag:"public_place"`
	Number      *string    `sql:"number" conversorTag:"number"`
	Complement  *string    `sql:"complement" conversorTag:"complement"`
	District    *string    `sql:"district" conversorTag:"district"`
	City        *string    `sql:"city" conversorTag:"city"`
	State       *string    `sql:"state" conversorTag:"state"`
	Country     *string    `sql:"country" conversorTag:"country"`
	BankID      *int64     `sql:"bank_id" conversorTag:"bank_id"`
}

type AgencyPag struct {
	Data  []Agency
	Next  *bool
	Total *int64
}
