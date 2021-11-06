package agency

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/agency"
	"github.com/hiagomf/bank-api/server/infraestructure/records/agency/postgres"
	"github.com/hiagomf/bank-api/server/utils"
)

type repository struct {
	pg *postgres.PGAgency
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGAgency{DB: db},
	}
}

// SelectPaginated - retorna as agências paginadas de acordo com os parâmetros informados
func (r *repository) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *agency.AgencyPag, err error) {
	return r.pg.SelectPaginated(parameters)
}
