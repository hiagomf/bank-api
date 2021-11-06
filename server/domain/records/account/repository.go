package account

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/account/postgres"
)

type repository struct {
	pg *postgres.PGAccount
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGAccount{DB: db},
	}
}
