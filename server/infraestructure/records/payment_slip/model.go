package payment_slip

import "time"

type PaymentSlip struct {
	ID               *int64     `sql:"id" conversorTag:"id"`
	Header           *string    `sql:"header" conversorTag:"header"`                       // cabeçalho (banco emissor)
	Assignor         *string    `sql:"assignor" conversorTag:"assignor"`                   // cedente
	IssuanceDate     *time.Time `sql:"issuance_date" conversorTag:"issuance_date"`         // data de emissão
	DueDate          *time.Time `sql:"due_date" conversorTag:"due_date"`                   // data de vencimento
	AgencyCode       *int64     `sql:"agency_code" conversorTag:"agency_code"`             // número da agência
	AccountCode      *int64     `sql:"account_code" conversorTag:"account_code"`           // número da conta
	VerifyingDigit   *int64     `sql:"verifying_digit" conversorTag:"verifying_digit"`     // dígito verificador
	GrossValue       *float64   `sql:"gross_value" conversorTag:"gross_value"`             // valor bruto, sem descontos, juros, etc
	Deduction        *float64   `sql:"deduction" conversorTag:"deduction"`                 // dedução
	Discount         *float64   `sql:"discount" conversorTag:"discount"`                   // desconto
	Penalty          *float64   `sql:"penalty" conversorTag:"penalty"`                     // multa
	Fees             *float64   `sql:"fees" conversorTag:"fees"`                           // juros
	AmoutCharged     *float64   `sql:"amount_charged" conversorTag:"amount_charged"`       // valor cobrado
	PayingDocument   *string    `sql:"paying_document" conversorTag:"paying_document"`     // Documento CPF/CNPJ do pagante
	PayingName       *string    `sql:"paying_name" conversorTag:"paying_name"`             // Nome do pagante
	ReceiverDocument *string    `sql:"receiver_document" conversorTag:"receiver_document"` // Documento CPF/CNPJ do beneficiário
	ReceiverName     *string    `sql:"receiver_name" conversorTag:"receiver_name"`         // Nome do beneficiário
	PaymentLocal     *string    `sql:"payment_local" conversorTag:"payment_local"`         // local de pagamento
	DigitableLine    *string    `sql:"digitable_line" conversorTag:"digitable_line"`       // linha digitável
	Instructions     *string    `sql:"instructions" conversorTag:"instructions"`           // instruções para pagamento
	Msg1             *string    `sql:"msg_1" conversorTag:"msg_1"`                         // mensagem 1
	Msg2             *string    `sql:"msg_2" conversorTag:"msg_2"`                         // mensagem 2
	Msg3             *string    `sql:"msg_3" conversorTag:"msg_3"`                         // mensagem 3
}

type PaymentSlipPag struct {
	Data  []PaymentSlip
	Next  *bool
	Total *int64
}
