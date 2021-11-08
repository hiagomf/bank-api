package account_transaction

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/actions/account_transaction/postgres"
)

type repository struct {
	pg *postgres.PGAccountTransaction
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGAccountTransaction{DB: db},
	}
}

// Deposit - deposita determinado valor na conta
func (r *repository) Deposit(id *int64, value *float64) (err error) {
	return r.pg.Deposit(id, value)
}
