package bank

import "github.com/hiagomf/bank-api/server/config/database"

type Service struct {
	repo IBank
}

func GetRepository(tx *database.DBTransaction) IBank {
	return novoRepo(tx)
}

func GetService(rReceived IBank) *Service {
	return &Service{
		repo: rReceived,
	}
}
