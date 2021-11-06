package account

import "github.com/hiagomf/bank-api/server/config/database"

type Service struct {
	repo IAccount
}

func GetRepository(tx *database.DBTransaction) IAccount {
	return novoRepo(tx)
}

func GetService(rReceived IAccount) *Service {
	return &Service{
		repo: rReceived,
	}
}
