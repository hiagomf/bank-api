package payment_slip

import "time"

type Request struct {
	Header         *string    `json:"header" binding:"required" conversorTag:"header"`                   // cabeçalho (banco emissor)
	Assignor       *string    `json:"assignor" binding:"required" conversorTag:"assignor"`               // cedente
	DueDate        *time.Time `json:"due_date" binding:"required" conversorTag:"due_date"`               // data de vencimento
	AgencyCode     *int64     `json:"agency_code" binding:"required" conversorTag:"agency_code"`         // número da agência
	AccountCode    *int64     `json:"account_code" binding:"required" conversorTag:"account_code"`       // número da conta
	VerifyingDigit *int64     `json:"verifying_digit" binding:"required" conversorTag:"verifying_digit"` // dígito verificador
	GrossValue     *float64   `json:"gross_value" binding:"required" conversorTag:"gross_value"`         // valor bruto, sem descontos, juros, etc
	Deduction      *float64   `json:"deduction" binding:"" conversorTag:"deduction"`                     // dedução
	Discount       *float64   `json:"discount" binding:"" conversorTag:"discount"`                       // desconto
	Penalty        *float64   `json:"penalty" binding:"" conversorTag:"penalty"`                         // multa
	Fees           *float64   `json:"fees" binding:"" conversorTag:"fees"`                               // juros
	PayingDocument *string    `json:"paying_document" binding:"required" conversorTag:"paying_document"` // documento CPF/CNPJ do pagante
	PayingName     *string    `json:"paying_name" binding:"required" conversorTag:"paying_name"`         // nome do pagante
	PaymentLocal   *string    `json:"payment_local" binding:"required" conversorTag:"payment_local"`     // local de pagamento
	Instructions   *string    `json:"instructions" conversorTag:"instructions"`                          // instruções para pagamento
	Msg1           *string    `json:"msg_1" conversorTag:"msg_1"`                                        // mensagem 1
	Msg2           *string    `json:"msg_2" conversorTag:"msg_2"`                                        // mensagem 2
	Msg3           *string    `json:"msg_3" conversorTag:"msg_3"`                                        // mensagem 3
}
