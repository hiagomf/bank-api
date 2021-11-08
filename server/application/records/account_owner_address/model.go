package account_owner_address

import "time"

type Request struct {
	OwnerID     *int64  `json:"account_owner_id" binding:"required" conversorTag:"account_owner_id"`
	ZipCode     *string `json:"zip_code" binding:"required,gte=9,lte=10" conversorTag:"zip_code"`
	PublicPlace *string `json:"public_place" binding:"required" conversorTag:"public_place"`
	Number      *string `json:"number" binding:"required" conversorTag:"number"`
	Complement  *string `json:"complement" binding:"" conversorTag:"complement"`
	District    *string `json:"district" binding:"required" conversorTag:"district"`
	City        *string `json:"city" binding:"required" conversorTag:"city"`
	State       *string `json:"state" binding:"required" conversorTag:"state"`
	Country     *string `json:"country" binding:"required" conversorTag:"country"`
}

// Response define a estrutura de dados do endereço titular da conta
type Response struct {
	ID          *int64     `json:"id,omitempty" conversorTag:"id"`
	CreatedAt   *time.Time `json:"created_at,omitempty" conversorTag:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" conversorTag:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" conversorTag:"deleted_at"`
	ZipCode     *string    `json:"zip_code,omitempty" conversorTag:"zip_code"`
	PublicPlace *string    `json:"public_place,omitempty" conversorTag:"public_place"`
	Number      *string    `json:"number,omitempty" conversorTag:"number"`
	Complement  *string    `json:"complement,omitempty" conversorTag:"complement"`
	District    *string    `json:"district,omitempty" conversorTag:"district"`
	City        *string    `json:"city,omitempty" conversorTag:"city"`
	State       *string    `json:"state,omitempty" conversorTag:"state"`
	Country     *string    `json:"country,omitempty" conversorTag:"country"`
	OwnerID     *int64     `json:"account_owner_id,omitempty" binding:"required" conversorTag:"account_owner_id"`
}

// ResponsePag define a estrutura de paginação do endereço do titular de conta
type ResponsePag struct {
	Data  []Response `json:"data,omitempty"`
	Next  *bool      `json:"next,omitempty"`
	Total *int64     `json:"total,omitempty"`
}
