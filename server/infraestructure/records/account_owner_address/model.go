package account_owner_address

import "time"

// AccountOwnerAddress define a estrutura de dados do endereço titular da conta
type AccountOwnerAddress struct {
	ID          *int64     `sql:"id" conversorTag:"id"`
	CreatedAt   *time.Time `sql:"created_at" conversorTag:"created_at"`
	UpdatedAt   *time.Time `sql:"updated_at" conversorTag:"updated_at"`
	DeletedAt   *time.Time `sql:"deleted_at" conversorTag:"deleted_at"`
	ZipCode     *string    `sql:"zip_code" conversorTag:"zip_code"`
	PublicPlace *string    `sql:"public_place" conversorTag:"public_place"`
	Number      *string    `sql:"number" conversorTag:"number"`
	Complement  *string    `sql:"complement" conversorTag:"complement"`
	District    *string    `sql:"district" conversorTag:"district"`
	City        *string    `sql:"city" conversorTag:"city"`
	State       *string    `sql:"state" conversorTag:"state"`
	Country     *string    `sql:"country" conversorTag:"country"`
	OwnerID     *int64     `sql:"account_owner_id" conversorTag:"account_owner_id"`
}

// AccountOwnerAddressPag define a estrutura de paginação do endereço do titular de conta
type AccountOwnerAddressPag struct {
	Data  []AccountOwnerAddress
	Next  *bool
	Total *int64
}
