package account_owner

import "time"

// AccountOwner define a estrutura de dados do titular da conta
type AccountOwner struct {
	ID         *int64     `sql:"id" conversorTag:"id"`
	CreatedAt  *time.Time `sql:"created_at" conversorTag:"created_at"`
	UpdatedAt  *time.Time `sql:"updated_at" conversorTag:"updated_at"`
	DeletedAt  *time.Time `sql:"deleted_at" conversorTag:"deleted_at"`
	Name       *string    `sql:"name" conversorTag:"name"`
	Document   *string    `sql:"document" conversorTag:"document"`
	BirthDate  *time.Time `sql:"birth_date" conversorTag:"birth_date"`
	FatherName *string    `sql:"father_name" conversorTag:"father_name"`
	MotherName *string    `sql:"mother_name" conversorTag:"mother_name"`
}

// AccountOwnerPag define a estrutura de paginação de titular de conta
type AccountOwnerPag struct {
	Data  []AccountOwner
	Next  *bool
	Total *int64
}
