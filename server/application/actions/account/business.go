package account

import (
	"context"
	"strconv"

	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/domain/actions/account"
	"github.com/hiagomf/bank-api/server/domain/actions/account_detail"
	"github.com/hiagomf/bank-api/server/domain/records/account_owner"
	"github.com/hiagomf/bank-api/server/domain/records/agency"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

// OpenAccount - realiza a abertura da conta
func OpenAccount(ctx context.Context, req *Request) (res *Response, err error) {
	var msgErrorDefault = "Erro ao abrir conta"
	res = new(Response)

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	accountRepo := account.GetRepository(tx)

	// Validando ID do titular
	if _, err := account_owner.GetRepository(tx).SelectOne(req.OwnerID); err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	// Validando ID da agência
	agency, err := agency.GetRepository(tx).SelectOne(req.AgencyID)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	// Validando se a conta já existe
	var params utils.ParametrosRequisicao
	params.Filtros = make(map[string][]string)
	params.Filtros["account_owner_id"] = []string{strconv.FormatInt(*req.OwnerID, 10)}
	params.Filtros["agency_id"] = []string{strconv.FormatInt(*req.AgencyID, 10)}
	params.Total = true

	list, err := accountRepo.SelectPaginated(&params)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	if list.Total != nil && *list.Total != 0 {
		return res, oops.Wrap(oops.NovoErr("já existe uma conta para esse usuário, consulte situação"), msgErrorDefault)
	}

	// encriptando senha pra salvar
	hash, err := utils.EncryptPassword(*req.Password)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	data := accountRepo.GetDataInfra()
	data.Password = &hash
	data.AgencyID = req.AgencyID
	data.OwnerID = req.OwnerID

	if err = accountRepo.Insert(data); err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	detailRepo := account_detail.GetRepository(tx)

	// Preenchendo detalhes
	dataDetail := detailRepo.GetDataInfra()
	dataDetail.AccountID = data.ID

	// Inserindo detalhes de poupança, para gerenciar o saldo
	if err = detailRepo.Insert(dataDetail); err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	res.AgencyCode = agency.Code
	res.Number = data.Number
	res.VerifyingDigit = data.VerifyingDigit

	return
}

// CloseAccount - realiza o fechamento da conta caso não exista saldo nela
func CloseAccount(ctx context.Context, id *int64) (err error) {
	var msgErrorDefault = "Erro ao fechar conta"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	if err = account.GetRepository(tx).Disable(id); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	return
}
