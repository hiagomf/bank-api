package bank

import "time"

// Bank - definindo estrutura de banco
type Bank struct {
	ID        *int64     `sql:"id" conversorTag:"id"`
	CreatedAt *time.Time `sql:"created_at" conversorTag:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at" conversorTag:"updated_at"`
	DeletedAt *time.Time `sql:"deleted_at" conversorTag:"deleted_at"`
	Code      *int64     `sql:"code" conversorTag:"code"`
	Name      *string    `sql:"name" conversorTag:"name"`
}

// BankPag - estrutura de banco paginado
type BankPag struct {
	Data  []Bank
	Next  *bool
	Total *int64
}
