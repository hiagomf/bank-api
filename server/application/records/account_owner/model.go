package account_owner

import "time"

type Response struct {
	ID        *int64     `json:"id,omitempty" conversorTag:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty" conversorTag:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" conversorTag:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" conversorTag:"deleted_at"`
}

type ResponsePag struct {
	Data  []Response `json:"data,omitempty"`
	Next  *bool      `json:"next,omitempty"`
	Total *int64     `json:"total,omitempty"`
}
