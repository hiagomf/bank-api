package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/account"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type PGAccount struct {
	DB *database.DBTransaction
}

// Insert - realiza a inserção de um registro no banco
func (pg *PGAccount) Insert(data *account.Account) (err error) {
	cols, vals, err := utils.FormatarInsertUpdate(data)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Insert("public.t_account").
		SetMap(valores).
		Suffix(`RETURNING id`).
		Scan(&data.ID); err != nil {
		return oops.Err(err)
	}
	return
}

// Update - realiza a alteração de um registro no banco
func (pg *PGAccount) Update(data *account.Account) (err error) {
	cols, vals, err := utils.FormatarInsertUpdate(data)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Update("t_account").
		SetMap(valores).
		Where(squirrel.Eq{
			"id":         data.ID,
			"deleted_at": nil,
		}).
		Suffix(`RETURNING "id"`).
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}

	return
}

// SelectOne - realiza a busca de uma ocorrência no banco
func (pg *PGAccount) SelectOne(id *string) (res *account.Account, err error) {
	res = new(account.Account)

	if err = pg.DB.Builder.
		Select(`
			TA.id,
			TA.created_at,
			TA.updated_at,
			TA.deleted_at,
			TA.number,
			TA.verifying_digit,
			TA.account_owner_id,
			TA.password,
			TA.agency_id
		`).
		From(`public.t_account TA`).
		Where(squirrel.Eq{
			"TA.id": id,
		}).
		Scan(
			&res.ID,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
			&res.Number,
			&res.VerifyingDigit,
			&res.OwnerID,
			&res.Password,
			&res.AgencyID,
		); err != nil {
		return res, oops.Err(err)
	}
	return
}

// SelectPaginated - realiza a busca paginada de tipos de ocorrências no banco
func (pg *PGAccount) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account.AccountPag, err error) {
	var occurrenceModel account.Account
	res = new(account.AccountPag)

	fields, _, err := parameters.ValidarCampos(&occurrenceModel)
	if err != nil {
		return res, oops.Err(err)
	}

	preQuery := pg.DB.Builder.
		Select(fields...).
		From(`public.t_account TA`)

	// Definindo filtros que poderão ser utilizados na consulta
	whereClause := parameters.CriarFiltros(preQuery, map[string]utils.Filtro{
		"not_in_id": utils.CriarFiltros("TA.id", utils.FlagFiltroNotIn),
		"deleted":   utils.CriarFiltros("(TA.deleted_at IS NOT NULL) = ?::BOOL", utils.FlagFiltroEq),
	})

	data, next, total, err := utils.ConfigurarPaginacao(parameters, &occurrenceModel, &whereClause)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Data, res.Next, res.Total = data.([]account.Account), next, total
	return
}

// Disable realiza o soft delete no banco do registro informado
func (pg *PGAccount) Disable(id *string) (err error) {
	if err = pg.DB.Builder.
		Update("public.t_account").
		Set("deleted_at", squirrel.Expr("NOW()")).
		Where(squirrel.Eq{
			"id":         id,
			"deleted_at": nil,
		}).
		Suffix(`RETURNING "id"`).
		Scan(new(string)); err != nil {
		if err == sql.ErrNoRows {
			return oops.NovoErr("esse registro já foi removido ou não existe")
		}
		return
	}
	return
}
