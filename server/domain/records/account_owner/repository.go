package account_owner

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/account_owner"
	"github.com/hiagomf/bank-api/server/infraestructure/records/account_owner/postgres"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type repository struct {
	pg *postgres.PGAccountOwner
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGAccountOwner{DB: db},
	}
}

// Insert - insere o titular da conta
func (r *repository) Insert(data *account_owner.AccountOwner) (err error) {
	return r.pg.Insert(data)
}

// Update - altera os dados do titular da conta
func (r *repository) Update(data *account_owner.AccountOwner) (err error) {
	return r.pg.Update(data)
}

// SelectOne - busca o titular da conta
func (r *repository) SelectOne(id *int64) (res *account_owner.AccountOwner, err error) {
	return r.pg.SelectOne(id)
}

// SelectPaginated - lista titulares das contas
func (r *repository) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account_owner.AccountOwnerPag, err error) {
	return r.pg.SelectPaginated(parameters)
}

// Disable - desativa o titular da conta
func (r *repository) Disable(id *int64) (err error) {
	return r.pg.Disable(id)
}

// ConvertToInfra - realiza a conversão de alguma model informada para model de infraestructure
func (r *repository) ConvertToInfra(data interface{}) (res *account_owner.AccountOwner, err error) {
	res = &account_owner.AccountOwner{}

	if err = utils.ConvertStruct(data, res); err != nil {
		return res, oops.Err(err)
	}
	return res, nil
}
