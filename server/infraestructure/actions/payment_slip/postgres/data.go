package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/actions/payment_slip"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type PGPaymentSlip struct {
	DB *database.DBTransaction
}

// Insert - realiza a inserção de um registro no banco
func (pg *PGPaymentSlip) Insert(data *payment_slip.PaymentSlip) (err error) {
	cols, vals, err := utils.FormatarInsertUpdate(data)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Insert("public.t_payment_slip").
		SetMap(valores).
		Suffix(`RETURNING "id"`).
		Scan(&data.ID); err != nil {
		return oops.Err(err)
	}
	return
}

// UpdateDigitableLine - altera a linha digitável
func (pg *PGPaymentSlip) UpdateDigitableLine(id *int64, value *string) (err error) {

	if err = pg.DB.Builder.
		Update(`t_payment_slip`).
		Set("digitable_line", value).
		Where(squirrel.Eq{
			"id": id,
		}).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); err != nil {
		return oops.Err(err)
	}
	return
}
