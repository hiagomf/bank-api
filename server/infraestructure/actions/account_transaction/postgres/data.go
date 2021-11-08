package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/oops"
)

type PGAccountTransaction struct {
	DB *database.DBTransaction
}

func (pg *PGAccountTransaction) Deposit(id *int64, value *float64) (err error) {

	if err = pg.DB.Builder.
		Update(`t_account_detail`).
		Set("balance", value).
		Where(squirrel.Eq{
			"id": id,
		}).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); err != nil {
		return oops.Err(err)
	}
	return
}
