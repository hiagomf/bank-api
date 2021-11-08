package account_owner

import "time"

type Request struct {
	Name       *string    `json:"name" binding:"required,gte=10" conversorTag:"name"`
	Document   *string    `json:"document" binding:"required" conversorTag:"document"`
	BirthDate  *time.Time `json:"birth_date" binding:"required" conversorTag:"birth_date"`
	FatherName *string    `json:"father_name" binding:"required,gte=10" conversorTag:"father_name"`
	MotherName *string    `json:"mother_name" binding:"required,gte=10" conversorTag:"mother_name"`
}

type Response struct {
	ID         *int64     `json:"id,omitempty" conversorTag:"id"`
	CreatedAt  *time.Time `json:"created_at,omitempty" conversorTag:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" conversorTag:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" conversorTag:"deleted_at"`
	Name       *string    `json:"name,omitempty" conversorTag:"name"`
	Document   *string    `json:"document,omitempty" conversorTag:"document"`
	BirthDate  *time.Time `json:"birth_date,omitempty" conversorTag:"birth_date"`
	FatherName *string    `json:"father_name,omitempty" conversorTag:"father_name"`
	MotherName *string    `json:"mother_name,omitempty" conversorTag:"mother_name"`
}

type ResponsePag struct {
	Data  []Response `json:"data,omitempty"`
	Next  *bool      `json:"next,omitempty"`
	Total *int64     `json:"total,omitempty"`
}
