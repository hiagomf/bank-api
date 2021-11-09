package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/bank"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type PGBank struct {
	DB *database.DBTransaction
}

// SelectOne - busca um banco pelo ID
func (pg *PGBank) SelectOne(id *int64) (res *bank.Bank, err error) {
	res = new(bank.Bank)

	if err = pg.DB.Builder.
		Select(`
			TB.id,
			TB.created_at,
			TB.updated_at,
			TB.deleted_at,
			TB.name,
			TB.code
		`).
		From(`public.t_bank TB`).
		Where(squirrel.Eq{
			"TB.id": id,
		}).
		Scan(
			&res.ID,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
			&res.Name,
			&res.Code,
		); err != nil {
		if err == sql.ErrNoRows {
			return res, oops.NovoErr("banco não encontrado, digite um ID válido")
		}
		return res, oops.Err(err)
	}
	return
}

// SelectPaginated - retorna os bancos paginados de acordo com os parâmetros informados
func (pg *PGBank) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *bank.BankPag, err error) {
	var bankModel bank.Bank
	res = new(bank.BankPag)

	fields, _, err := parameters.ValidarCampos(&bankModel)
	if err != nil {
		return res, oops.Err(err)
	}

	preQuery := pg.DB.Builder.
		Select(fields...).
		From(`public.t_bank TB`)

	// Definindo filtros que poderão ser utilizados na consulta
	whereClause := parameters.CriarFiltros(preQuery, map[string]utils.Filtro{
		"not_in_id": utils.CriarFiltros("TB.id", utils.FlagFiltroNotIn),
		"deleted":   utils.CriarFiltros("(TB.deleted_at IS NOT NULL) = ?::BOOL", utils.FlagFiltroEq),
	})

	data, next, total, err := utils.ConfigurarPaginacao(parameters, &bankModel, &whereClause)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Data, res.Next, res.Total = data.([]bank.Bank), next, total
	return
}
