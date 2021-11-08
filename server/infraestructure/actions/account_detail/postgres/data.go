package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/actions/account_detail"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type PGAccountDetail struct {
	DB *database.DBTransaction
}

// Insert - realiza a inserção de um registro no banco
func (pg *PGAccountDetail) Insert(data *account_detail.AccountDetail) (err error) {
	cols, vals, err := utils.FormatarInsertUpdate(data)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Insert("public.t_account_detail").
		SetMap(valores).
		Suffix(`RETURNING "id"`).
		Scan(&data.ID); err != nil {
		return oops.Err(err)
	}
	return
}

// GetAccountDetail - busca detalhes da conta com base ns dados informados
func (pg *PGAccountDetail) GetAccountDetail(data *account_detail.Access) (res *account_detail.AccountDetail, err error) {
	res = new(account_detail.AccountDetail)

	if err = pg.DB.Builder.
		Select(`
			TAD.id,
			TAD.created_at,
			TAD.updated_at,
			TAD.deleted_at,
			TAD.blocked,
			TAD.balance,
			TAD.account_id,
			TA.password
		`).
		From(`t_account_detail TAD`).
		Join(`t_account TA ON TAD.account_id = TA.id`).
		Join(`t_agency TAG ON TA.agency_id = TAG.id AND TAG.code = ?`, data.AgencyCode).
		Where(squirrel.Eq{
			"TA.number":          data.AccountNumber,
			"TA.verifying_digit": data.VerifyingDigit,
		}).
		Scan(
			&res.ID,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
			&res.Blocked,
			&res.Balance,
			&res.AccountID,
			&res.AccountPassword,
		); err != nil {
		if err == sql.ErrNoRows {
			return res, oops.NovoErr("conta não encontrada")
		}
		return res, oops.Err(err)
	}
	return
}
