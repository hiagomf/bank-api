package payment_slip

import "github.com/hiagomf/bank-api/server/infraestructure/actions/payment_slip"

type IPaymentSlip interface {
	Insert(data *payment_slip.PaymentSlip) (err error)
	UpdateDigitableLine(id *int64, value *string) (err error)
	ConvertToInfra(data interface{}) (res *payment_slip.PaymentSlip, err error)
}
