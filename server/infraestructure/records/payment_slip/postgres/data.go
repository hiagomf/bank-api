package postgres

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/records/payment_slip"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type PGPaymentSlip struct {
	DB *database.DBTransaction
}

// SelectPaginated - realiza a busca paginada
func (pg *PGPaymentSlip) SelectPaginated(parameters *utils.ParametrosRequisicao) (res *payment_slip.PaymentSlipPag, err error) {
	var bankModel payment_slip.PaymentSlip
	res = new(payment_slip.PaymentSlipPag)

	fields, _, err := parameters.ValidarCampos(&bankModel)
	if err != nil {
		return res, oops.Err(err)
	}

	preQuery := pg.DB.Builder.
		Select(fields...).
		From(`public.t_payment_slip TPS`)

	// Definindo filtros que poderão ser utilizados na consulta
	whereClause := parameters.CriarFiltros(preQuery, map[string]utils.Filtro{
		"not_in_id":       utils.CriarFiltros("TPS.id", utils.FlagFiltroNotIn),
		"header":          utils.CriarFiltros("header = ?", utils.FlagFiltroEq),
		"assignor":        utils.CriarFiltros("assignor = ?", utils.FlagFiltroEq),
		"due_date":        utils.CriarFiltros("due_date = ?", utils.FlagFiltroEq),
		"agency_code":     utils.CriarFiltros("agency_code = ?", utils.FlagFiltroEq),
		"account_code":    utils.CriarFiltros("account_code = ?", utils.FlagFiltroEq),
		"verifying_digit": utils.CriarFiltros("verifying_digit = ?", utils.FlagFiltroEq),
		"paying_document": utils.CriarFiltros("paying_document = ?", utils.FlagFiltroEq),
		"digitable_line":  utils.CriarFiltros("digitable_line = ?", utils.FlagFiltroEq),
	})

	data, next, total, err := utils.ConfigurarPaginacao(parameters, &bankModel, &whereClause)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Data, res.Next, res.Total = data.([]payment_slip.PaymentSlip), next, total
	return
}
