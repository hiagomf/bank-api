package account_transaction

import (
	"context"

	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/domain/actions/account_detail"
	"github.com/hiagomf/bank-api/server/domain/actions/account_transaction"
	"github.com/hiagomf/bank-api/server/oops"
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
	if err = transactionRepo.Deposit(detailData.ID, &newValue); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	return
}
