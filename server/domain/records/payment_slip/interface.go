package payment_slip

import (
	"github.com/hiagomf/bank-api/server/infraestructure/records/payment_slip"
	"github.com/hiagomf/bank-api/server/utils"
)

type IPaymentSlip interface {
	SelectPaginated(parameters *utils.ParametrosRequisicao) (res *payment_slip.PaymentSlipPag, err error)
}
