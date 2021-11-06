package postgres

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/agency"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type PGAgency struct {
	DB *database.DBTransaction
}

// SelectPaginated - retorna os bancos paginados de acordo com os parâmetros informados
func (pg *PGAgency) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *agency.AgencyPag, err error) {
	var agencymodel agency.Agency
	res = new(agency.AgencyPag)

	fields, _, err := parameters.ValidarCampos(&agencymodel)
	if err != nil {
		return res, oops.Err(err)
	}

	preQuery := pg.DB.Builder.
		Select(fields...).
		From(`public.t_agency TA`)

	// Definindo filtros que poderão ser utilizados na consulta
	whereClause := parameters.CriarFiltros(preQuery, map[string]utils.Filtro{
		"not_in_id":   utils.CriarFiltros("TA.id", utils.FlagFiltroNotIn),
		"deleted":     utils.CriarFiltros("(TA.deleted_at IS NOT NULL) = ?::BOOL", utils.FlagFiltroEq),
		"bank_id":     utils.CriarFiltros("TA.bank_id = ?::INTEGER", utils.FlagFiltroEq),
		"main_agency": utils.CriarFiltros("TA.main_agency = ?::BOOL", utils.FlagFiltroEq),
	})

	data, next, total, err := utils.ConfigurarPaginacao(parameters, &agencymodel, &whereClause)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Data, res.Next, res.Total = data.([]agency.Agency), next, total
	return
}
