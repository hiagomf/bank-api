package agency

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/agency/postgres"
)

type repository struct {
	pg *postgres.PGAgency
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGAgency{DB: db},
	}
}
