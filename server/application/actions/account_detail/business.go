package account_detail

import (
	"context"

	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/domain/actions/account_detail"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

// CheckDetails - verifica os detalhes da conta, inclusive o saldo
func CheckDetails(ctx context.Context, req *Request) (res *Response, err error) {
	var msgErrorDefault = "Erro ao buscar detalhes de conta"
	res = new(Response)

	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	detailRepo := account_detail.GetRepository(tx)

	accessData, err := detailRepo.ConvertAcessToInfra(req)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	// buscando conta e seus detalhes
	detailData, err := detailRepo.GetAccountDetail(accessData)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	// comparando hashes de senhas
	ok, err := utils.CheckEncryptedPassword(*accessData.Password, *detailData.AccountPassword)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	// verificando se a validação de senha passou
	if !ok {
		return res, oops.Wrap(oops.NovoErr("senha inválida"), msgErrorDefault)
	}

	if err = utils.ConvertStruct(detailData, res); err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	return
}
