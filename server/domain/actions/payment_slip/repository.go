package payment_slip

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/actions/payment_slip"
	"github.com/hiagomf/bank-api/server/infraestructure/actions/payment_slip/postgres"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type repository struct {
	pg *postgres.PGPaymentSlip
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGPaymentSlip{DB: db},
	}
}

// Insert - insere um boleto
func (r *repository) Insert(data *payment_slip.PaymentSlip) (err error) {
	return r.pg.Insert(data)
}

// UpdateDigitableLine - altera a linha digitável
func (r *repository) UpdateDigitableLine(id *int64, value *string) (err error) {
	return r.pg.UpdateDigitableLine(id, value)
}

// ConvertToInfra - realiza a conversão de alguma model informada para model de infraestructure
func (r *repository) ConvertToInfra(data interface{}) (res *payment_slip.PaymentSlip, err error) {
	res = &payment_slip.PaymentSlip{}

	if err = utils.ConvertStruct(data, res); err != nil {
		return res, oops.Err(err)
	}
	return res, nil
}
