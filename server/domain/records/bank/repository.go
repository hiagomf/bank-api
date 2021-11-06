package bank

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/bank/postgres"
)

type repository struct {
	pg *postgres.PGBank
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGBank{DB: db},
	}
}
