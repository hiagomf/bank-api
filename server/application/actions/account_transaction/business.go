package account_transaction

import (
	"context"

	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/domain/actions/account_detail"
	"github.com/hiagomf/bank-api/server/domain/actions/account_transaction"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

// Deposit - deposita no saldo da conta
func Deposit(ctx context.Context, req *DepositRequest) (err error) {
	var msgErrorDefault = "Erro ao buscar detalhes de conta"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	detailRepo := account_detail.GetRepository(tx)
	transactionRepo := account_transaction.GetRepository(tx)

	accessData, err := detailRepo.ConvertAcessToInfra(req)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	// buscando conta e seus detalhes
	detailData, err := detailRepo.GetAccountDetail(accessData)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	newValue := *detailData.Balance + *req.Value
	if err = transactionRepo.UpdateValue(detailData.ID, &newValue); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	return
}

func Transfer(ctx context.Context, req *TransferRequest) (err error) {
	var msgErrorDefault = "Erro ao buscar detalhes de conta"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	detailRepo := account_detail.GetRepository(tx)
	transactionRepo := account_transaction.GetRepository(tx)

	// separando a conta de origem para validação e operação com valores
	accessDataOrigin := detailRepo.GetAccessInfra()
	accessDataOrigin.AgencyCode = req.AgencyCode
	accessDataOrigin.AccountNumber = req.AccountNumber
	accessDataOrigin.VerifyingDigit = req.VerifyingDigit
	accessDataOrigin.Password = req.Password

	// buscando conta originária e seus detalhes
	detailDataOrigin, err := detailRepo.GetAccountDetail(accessDataOrigin)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	// comparando hashes de senhas
	ok, err := utils.CheckEncryptedPassword(*accessDataOrigin.Password, *detailDataOrigin.AccountPassword)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	// verificando se a validação de senha passou
	if !ok {
		return oops.Wrap(oops.NovoErr("senha inválida"), msgErrorDefault)
	}

	// Validando se o saldo na conta de origem é suficiente
	newBalance := *detailDataOrigin.Balance - *req.Value
	if newBalance < 0 {
		return oops.Wrap(oops.NovoErr("saldo insuficiente na conta de origem"), msgErrorDefault)
	}

	// buscando conta que irá receber o saldo e seus detalhes
	accessDataReceiver := detailRepo.GetAccessInfra()
	accessDataReceiver.AgencyCode = req.ToAgencyCode
	accessDataReceiver.AccountNumber = req.ToAccountNumber
	accessDataReceiver.VerifyingDigit = req.ToVerifyingDigit
	detailDataReceiver, err := detailRepo.GetAccountDetail(accessDataReceiver)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	// Atualizando o saldo da conta que envia o dinheiro
	if err = transactionRepo.UpdateValue(detailDataOrigin.ID, &newBalance); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	// Atualizando saldo na conta que receberá o valor
	newBalance = *detailDataReceiver.Balance + *req.Value
	if err = transactionRepo.UpdateValue(detailDataReceiver.ID, &newBalance); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	return
}
