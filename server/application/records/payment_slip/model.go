package payment_slip

import "time"

type Response struct {
	ID               *int64     `json:"id,omitempty" conversorTag:"id"`
	Header           *string    `json:"header,omitempty" conversorTag:"header"`                       // cabeçalho (banco emissor)
	Assignor         *string    `json:"assignor,omitempty" conversorTag:"assignor"`                   // cedente
	IssuanceDate     *time.Time `json:"issuance_date,omitempty" conversorTag:"issuance_date"`         // data de emissão
	DueDate          *time.Time `json:"due_date,omitempty" conversorTag:"due_date"`                   // data de vencimento
	AgencyCode       *int64     `json:"agency_code,omitempty" conversorTag:"agency_code"`             // número da agência
	AccountCode      *int64     `json:"account_code,omitempty" conversorTag:"account_code"`           // número da conta
	VerifyingDigit   *int64     `json:"verifying_digit,omitempty" conversorTag:"verifying_digit"`     // dígito verificador
	GrossValue       *float64   `json:"gross_value,omitempty" conversorTag:"gross_value"`             // valor bruto, sem descontos, juros, etc
	Deduction        *float64   `json:"deduction,omitempty" conversorTag:"deduction"`                 // dedução
	Discount         *float64   `json:"discount,omitempty" conversorTag:"discount"`                   // desconto
	Penalty          *float64   `json:"penalty,omitempty" conversorTag:"penalty"`                     // multa
	Fees             *float64   `json:"fees,omitempty" conversorTag:"fees"`                           // juros
	AmoutCharged     *float64   `json:"amount_charged,omitempty" conversorTag:"amount_charged"`       // valor cobrado
	PayingDocument   *string    `json:"paying_document,omitempty" conversorTag:"paying_document"`     // Documento CPF/CNPJ do pagante
	PayingName       *string    `json:"paying_name,omitempty" conversorTag:"paying_name"`             // Nome do pagante
	ReceiverDocument *string    `json:"receiver_document,omitempty" conversorTag:"receiver_document"` // Documento CPF/CNPJ do beneficiário
	ReceiverName     *string    `json:"receiver_name,omitempty" conversorTag:"receiver_name"`         // Nome do beneficiário
	PaymentLocal     *string    `json:"payment_local,omitempty" conversorTag:"payment_local"`         // local de pagamento
	DigitableLine    *string    `json:"digitable_line,omitempty" conversorTag:"digitable_line"`       // linha digitável
	Instructions     *string    `json:"instructions,omitempty" conversorTag:"instructions"`           // instruções para pagamento
	Msg1             *string    `json:"msg_1,omitempty" conversorTag:"msg_1"`                         // mensagem 1
	Msg2             *string    `json:"msg_2,omitempty" conversorTag:"msg_2"`                         // mensagem 2
	Msg3             *string    `json:"msg_3,omitempty" conversorTag:"msg_3"`                         // mensagem 3
}

type ResponsePag struct {
	Data  []Response `json:"data,omitempty"`
	Next  *bool      `json:"next,omitempty"`
	Total *int64     `json:"total,omitempty"`
}
