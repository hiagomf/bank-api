package account

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/actions/account"
	"github.com/hiagomf/bank-api/server/infraestructure/actions/account/postgres"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type repository struct {
	pg *postgres.PGAccount
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGAccount{DB: db},
	}
}

// Insert - insere uma conta
func (r *repository) Insert(data *account.Account) (err error) {
	return r.pg.Insert(data)
}

// Update - atualiza uma conta
func (r *repository) Update(data *account.Account) (err error) {
	return r.pg.Update(data)
}

// SelectOne - busca uma conta
func (r *repository) SelectOne(id *int64) (res *account.Account, err error) {
	return r.pg.SelectOne(id)
}

// SelectPaginated - realiza uma busca paginada
func (r *repository) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account.AccountPag, err error) {
	return r.pg.SelectPaginated(parameters)
}

// Disable - desabilita uma conta
func (r *repository) Disable(id *int64) (err error) {
	return r.pg.Disable(id)
}

// ConvertToInfra - realiza a conversão de alguma model informada para model de infraestructure
func (r *repository) ConvertToInfra(data interface{}) (res *account.Account, err error) {
	res = &account.Account{}

	if err = utils.ConvertStruct(data, res); err != nil {
		return res, oops.Err(err)
	}
	return res, nil
}

// GetDataInfra - retorna um data de infra, para não ferir a arquitetura
func (r *repository) GetDataInfra() (res *account.Account) {
	return new(account.Account)
}
