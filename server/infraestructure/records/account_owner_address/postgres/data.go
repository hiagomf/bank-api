package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/account_owner_address"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type PGAccountOwnerAddress struct {
	DB *database.DBTransaction
}

// Insert - realiza a inserção de um registro no banco
func (pg *PGAccountOwnerAddress) Insert(data *account_owner_address.AccountOwnerAddress) (err error) {
	cols, vals, err := utils.FormatarInsertUpdate(data)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Insert("public.t_account_owner_address").
		SetMap(valores).
		Suffix(`RETURNING id`).
		Scan(&data.ID); err != nil {
		return oops.Err(err)
	}
	return
}

// Update - realiza a alteração de um registro no banco
func (pg *PGAccountOwnerAddress) Update(data *account_owner_address.AccountOwnerAddress) (err error) {
	cols, vals, err := utils.FormatarInsertUpdate(data)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Update("t_account_owner_address").
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
func (pg *PGAccountOwnerAddress) SelectOne(id *int64) (res *account_owner_address.AccountOwnerAddress, err error) {
	res = new(account_owner_address.AccountOwnerAddress)

	if err = pg.DB.Builder.
		Select(`
			TAOA.id,
			TAOA.created_at,
			TAOA.updated_at,
			TAOA.deleted_at,
			TAOA.zip_code,
			TAOA.public_place,
			TAOA.number,
			TAOA.complement,
			TAOA.district,
			TAOA.city,
			TAOA.state,
			TAOA.country
		`).
		From(`public.t_account_owner_address TAOA`).
		Where(squirrel.Eq{
			"TAOA.id": id,
		}).
		Scan(
			&res.ID,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
			&res.ZipCode,
			&res.PublicPlace,
			&res.Number,
			&res.Complement,
			&res.District,
			&res.City,
			&res.State,
			&res.Country,
		); err != nil {
		return res, oops.Err(err)
	}
	return
}

// SelectPaginated - realiza a busca paginada de tipos de ocorrências no banco
func (pg *PGAccountOwnerAddress) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account_owner_address.AccountOwnerAddressPag, err error) {
	var occurrenceModel account_owner_address.AccountOwnerAddress
	res = new(account_owner_address.AccountOwnerAddressPag)

	fields, _, err := parameters.ValidarCampos(&occurrenceModel)
	if err != nil {
		return res, oops.Err(err)
	}

	preQuery := pg.DB.Builder.
		Select(fields...).
		From(`public.t_account_owner_address TAO`)

	// Definindo filtros que poderão ser utilizados na consulta
	whereClause := parameters.CriarFiltros(preQuery, map[string]utils.Filtro{
		"not_in_id":        utils.CriarFiltros("TAO.id", utils.FlagFiltroNotIn),
		"deleted":          utils.CriarFiltros("(TAO.deleted_at IS NOT NULL) = ?::BOOL", utils.FlagFiltroEq),
		"zip_code":         utils.CriarFiltros("TAO.zip_code = ?", utils.FlagFiltroEq),
		"public_place":     utils.CriarFiltros("TAO.public_place = ?", utils.FlagFiltroEq),
		"number":           utils.CriarFiltros("TAO.number = ?", utils.FlagFiltroEq),
		"complement":       utils.CriarFiltros("TAO.complement = ?", utils.FlagFiltroEq),
		"district":         utils.CriarFiltros("TAO.district = ?", utils.FlagFiltroEq),
		"city":             utils.CriarFiltros("TAO.city = ?", utils.FlagFiltroEq),
		"state":            utils.CriarFiltros("TAO.state = ?", utils.FlagFiltroEq),
		"country":          utils.CriarFiltros("TAO.country = ?", utils.FlagFiltroEq),
		"account_owner_id": utils.CriarFiltros("TAO.account_owner_id = ?", utils.FlagFiltroEq),
	})

	data, next, total, err := utils.ConfigurarPaginacao(parameters, &occurrenceModel, &whereClause)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Data, res.Next, res.Total = data.([]account_owner_address.AccountOwnerAddress), next, total
	return
}

// Disable realiza o soft delete no banco do registro informado
func (pg *PGAccountOwnerAddress) Disable(id *int64) (err error) {
	if err = pg.DB.Builder.
		Update("public.t_account_owner_address").
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
		return oops.Err(err)
	}
	return
}

// DisableAllActives realiza o soft delete no banco do registros que se encaixem na condição
func (pg *PGAccountOwnerAddress) DisableAllActives(ownerID *int64) (err error) {
	if err = pg.DB.Builder.
		Update("public.t_account_owner_address").
		Set("deleted_at", squirrel.Expr("NOW()")).
		Where(squirrel.Eq{
			"account_owner_id": ownerID,
			"deleted_at":       nil,
		}).
		Scan(); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return oops.Err(err)
	}
	return
}
