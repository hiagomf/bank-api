package account

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/account"
	"github.com/hiagomf/bank-api/server/infraestructure/records/account/postgres"
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

func (r *repository) Insert(data *account.Account) (err error) {
	return r.pg.Insert(data)
}

func (r *repository) Update(data *account.Account) (err error) {
	return r.pg.Update(data)
}

func (r *repository) SelectOne(id *string) (res *account.Account, err error) {
	return r.pg.SelectOne(id)
}

func (r *repository) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account.AccountPag, err error) {
	return r.pg.SelectPaginated(parameters)
}

func (r *repository) Disable(id *string) (err error) {
	return r.pg.Disable(id)
}
