package bank

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/bank"
	"github.com/hiagomf/bank-api/server/infraestructure/records/bank/postgres"
	"github.com/hiagomf/bank-api/server/utils"
)

type repository struct {
	pg *postgres.PGBank
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGBank{DB: db},
	}
}

// SelectPaginated - retorna os bancos paginados de acordo com os parâmetros informados
func (r *repository) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *bank.BankPag, err error) {
	return r.pg.SelectPaginated(parameters)
}
