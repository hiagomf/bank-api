package account_owner

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/account_owner/postgres"
)

type repository struct {
	pg *postgres.PGAccountOwner
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGAccountOwner{DB: db},
	}
}
