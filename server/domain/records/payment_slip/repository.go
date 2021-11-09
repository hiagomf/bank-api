package payment_slip

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/payment_slip"
	"github.com/hiagomf/bank-api/server/infraestructure/records/payment_slip/postgres"
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

// SelectPaginated - busca os boletos de maneira paginada de acordo com seus query params
func (r *repository) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *payment_slip.PaymentSlipPag, err error) {
	return r.pg.SelectPaginated(parameters)
}
