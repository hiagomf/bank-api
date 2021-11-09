package payment_slip

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/domain/actions/account"
	"github.com/hiagomf/bank-api/server/domain/actions/account_detail"
	"github.com/hiagomf/bank-api/server/domain/actions/payment_slip"
	"github.com/hiagomf/bank-api/server/domain/records/account_owner"
	"github.com/hiagomf/bank-api/server/domain/records/agency"
	"github.com/hiagomf/bank-api/server/domain/records/bank"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

// GeneratePaymentSlip - gerar boleto de pagamento
func GeneratePaymentSlip(ctx context.Context, req *Request) (digitableLine *string, err error) {
	var msgErrorDefault = "Erro ao gerar boleto"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	paymentSlipRepo := payment_slip.GetRepository(tx)
	detailRepo := account_detail.GetRepository(tx)
	accessData := detailRepo.GetAccessInfra()

	accessData.AccountNumber = req.AccountCode
	accessData.VerifyingDigit = req.VerifyingDigit
	accessData.AgencyCode = req.AgencyCode

	// buscando conta e seus detalhes pelo CÓDIGO, não pelo ID
	accountDetail, err := detailRepo.GetAccountDetail(accessData)
	if err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}

	// Buscando conta
	account, err := account.GetRepository(tx).SelectOne(accountDetail.AccountID)
	if err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}

	owner, err := account_owner.GetRepository(tx).SelectOne(account.OwnerID)
	if err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}

	// Buscando agência
	agency, err := agency.GetRepository(tx).SelectOne(account.AgencyID)
	if err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}

	// Buscando banco
	bank, err := bank.GetRepository(tx).SelectOne(agency.BankID)
	if err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}

	dataInfra, err := paymentSlipRepo.ConvertToInfra(req)
	if err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}

	// Preenchendo dados básicos do boleto
	dataInfra.IssuanceDate = utils.PonteiroTime(time.Now())
	dataInfra.AmoutCharged = dataInfra.GrossValue
	dataInfra.ReceiverDocument = owner.Document
	dataInfra.ReceiverName = owner.Name

	// Aplicando dedução
	if dataInfra.Deduction != nil {
		dataInfra.AmoutCharged = utils.PonteiroFloat64(*dataInfra.AmoutCharged - *dataInfra.Deduction)
	}

	// Aplicando desconto
	if dataInfra.Discount != nil {
		dataInfra.AmoutCharged = utils.PonteiroFloat64(*dataInfra.AmoutCharged - *dataInfra.Discount)
	}

	// Aplicando multa
	if dataInfra.Penalty != nil {
		dataInfra.AmoutCharged = utils.PonteiroFloat64(*dataInfra.AmoutCharged + *dataInfra.Penalty)
	}

	// Aplicanto juros
	if dataInfra.Fees != nil {
		dataInfra.AmoutCharged = utils.PonteiroFloat64(*dataInfra.AmoutCharged + *dataInfra.Fees)
	}

	if err = paymentSlipRepo.Insert(dataInfra); err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}

	// Os três números iniciais indicam o código do banco emissor, de acordo com tabela da Febraban.
	bankCodeStr := strconv.FormatInt(*bank.Code, 10)    // convertendo código pra string
	qtBankDigits := utf8.RuneCountInString(bankCodeStr) // contando dígitos
	var pt1 string = bankCodeStr

	// preenchendo de acordo com a regra a pt1
	for i := qtBankDigits; i < 4; i++ {
		pt1 = `0` + pt1
	}

	// O quarto número representa o tipo da moeda: 9 para o Real e 0 para outras moedas.
	pt2 := `9`

	/**
	* Os próximos 25 números são definidos pelo banco emissor.
	* Cada instituição pode usá-los como preferir.
	* Geralmente, eles trazem informações sobre a pessoa ou empresa cobradora,
	* número da agência, número identificador do boleto, etc.
	* Implementação: Identificador do boleto incluso no número, só é possível gerar boletos de valores diferentes
	* Poderia ser implementado aqui o ID da venda por exemplo, para garantir a geração única de boleto
	**/
	idStr := strconv.FormatInt(*dataInfra.ID, 10)
	qtIdStr := utf8.RuneCountInString(idStr) // contando dígitos

	pt3 := idStr
	for i := qtIdStr; i < 25; i++ {
		pt3 = `0` + pt3
	}

	/**
	* O 30º número, que fica isolado em um campo, é o dígito verificador.
	* Ele é gerado a partir do cálculo dos números anteriores e tem a função de
	* garantir que os códigos estejam todos corretos.
	**/
	pt4 := `1`

	/**
	* Os quatro números que aparecem depois do dígito verificador representam a data de vencimento.
	* Este número é referente a quantidade de dias passados desde a data-base
	* estipulada pelo Banco Central: 7 de outubro de 1997.
	* Ou seja, o número de dias entre a data-base e a data de vencimento.
	**/
	dateBC := time.Date(1997, time.October, 7, 0, 0, 0, 0, time.Local)
	days := time.Since(dateBC).Hours() / 24
	pt5 := strconv.Itoa(int(days))

	/**
	* Os dez últimos algarismos indicam o valor do documento sem desconto.
	* Ex.: se o boleto tem o valor de R$1000,00, o final será: 0000100000.
	**/
	pt6 := fmt.Sprintf("%.2f", *req.GrossValue)
	pt6 = strings.Replace(pt6, ".", "", 1)
	qtGVDigits := utf8.RuneCountInString(pt6) // contando dígitos

	// preenchendo de acordo com a regra a pt6
	for i := qtGVDigits; i < 10; i++ {
		pt6 = `0` + pt6
	}

	// concatenando todas as partes
	dl := pt1 + pt2 + pt3 + pt4 + pt5 + pt6

	if err = paymentSlipRepo.UpdateDigitableLine(dataInfra.ID, &dl); err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return digitableLine, oops.Wrap(err, msgErrorDefault)
	}

	return &dl, err
}
