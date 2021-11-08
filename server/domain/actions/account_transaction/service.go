package account_transaction

import "github.com/hiagomf/bank-api/server/config/database"

type Service struct {
	repo IAccountTransaction
}

func GetRepository(tx *database.DBTransaction) IAccountTransaction {
	return novoRepo(tx)
}

func GetService(rReceived IAccountTransaction) *Service {
	return &Service{
		repo: rReceived,
	}
}
