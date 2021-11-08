package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/agency"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type PGAgency struct {
	DB *database.DBTransaction
}

// SelectPaginated - retorna as agências paginadas de acordo com os parâmetros informados
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

// SelectOne - realiza a busca de uma agência no banco
func (pg *PGAgency) SelectOne(id *int64) (res *agency.Agency, err error) {
	res = new(agency.Agency)

	if err = pg.DB.Builder.
		Select(`
			TA.id,
			TA.created_at,
			TA.updated_at,
			TA.deleted_at,
			TA.bank_id,
			TA.code,
			TA.main_agency,
			TA.zip_code,
			TA.public_place,
			TA.number,
			TA.complement,
			TA.district,
			TA.city,
			TA.state,
			TA.country
		`).
		From(`public.t_agency TA`).
		Where(squirrel.Eq{
			"TA.id": id,
		}).
		Scan(
			&res.ID,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
			&res.BankID,
			&res.Code,
			&res.MainAgency,
			&res.ZipCode,
			&res.PublicPlace,
			&res.Number,
			&res.Complement,
			&res.District,
			&res.City,
			&res.State,
			&res.Country,
		); err != nil {
		if err == sql.ErrNoRows {
			return res, oops.NovoErr("agência não encontrada")
		}
		return res, oops.Err(err)
	}
	return
}
