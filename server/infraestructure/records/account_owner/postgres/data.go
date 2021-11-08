package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/account_owner"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type PGAccountOwner struct {
	DB *database.DBTransaction
}

// Insert - realiza a inserção de um registro no banco
func (pg *PGAccountOwner) Insert(data *account_owner.AccountOwner) (err error) {
	cols, vals, err := utils.FormatarInsertUpdate(data)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Insert("public.t_account_owner").
		SetMap(valores).
		Suffix(`RETURNING id`).
		Scan(&data.ID); err != nil {
		return oops.Err(err)
	}
	return
}

// Update - realiza a alteração de um registro no banco
func (pg *PGAccountOwner) Update(data *account_owner.AccountOwner) (err error) {
	cols, vals, err := utils.FormatarInsertUpdate(data)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Update("t_account_owner").
		SetMap(valores).
		Where(squirrel.Eq{
			"id":         data.ID,
			"deleted_at": nil,
		}).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); err != nil {
		return oops.Err(err)
	}

	return
}

// SelectOne - realiza a busca de uma ocorrência no banco
func (pg *PGAccountOwner) SelectOne(id *int64) (res *account_owner.AccountOwner, err error) {
	res = new(account_owner.AccountOwner)

	if err = pg.DB.Builder.
		Select(`
			TAO.id,
			TAO.created_at,
			TAO.updated_at,
			TAO.deleted_at,
			TAO.name,
			TAO.document,
			TAO.birth_date,
			TAO.father_name,
			TAO.mother_name
		`).
		From(`public.t_account_owner TAO`).
		Where(squirrel.Eq{
			"TAO.id": id,
		}).
		Scan(
			&res.ID,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
			&res.Name,
			&res.Document,
			&res.BirthDate,
			&res.FatherName,
			&res.MotherName,
		); err != nil {
		return res, oops.Err(err)
	}
	return
}

// SelectPaginated - realiza a busca paginada de tipos de ocorrências no banco
func (pg *PGAccountOwner) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account_owner.AccountOwnerPag, err error) {
	var occurrenceModel account_owner.AccountOwner
	res = new(account_owner.AccountOwnerPag)

	fields, _, err := parameters.ValidarCampos(&occurrenceModel)
	if err != nil {
		return res, oops.Err(err)
	}

	preQuery := pg.DB.Builder.
		Select(fields...).
		From(`public.t_account_owner TAO`)

	// Definindo filtros que poderão ser utilizados na consulta
	whereClause := parameters.CriarFiltros(preQuery, map[string]utils.Filtro{
		"not_in_id": utils.CriarFiltros("TAO.id", utils.FlagFiltroNotIn),
		"deleted":   utils.CriarFiltros("(TAO.deleted_at IS NOT NULL) = ?::BOOL", utils.FlagFiltroEq),
		"document":  utils.CriarFiltros("TAO.document = ?", utils.FlagFiltroEq),
	})

	data, next, total, err := utils.ConfigurarPaginacao(parameters, &occurrenceModel, &whereClause)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Data, res.Next, res.Total = data.([]account_owner.AccountOwner), next, total
	return
}

// Disable realiza o soft delete no banco do registro informado
func (pg *PGAccountOwner) Disable(id *int64) (err error) {
	if err = pg.DB.Builder.
		Update("public.t_account_owner").
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
