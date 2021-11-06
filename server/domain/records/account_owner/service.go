package account_owner

import "github.com/hiagomf/bank-api/server/config/database"

type Service struct {
	repo IAccountOwner
}

func GetRepository(tx *database.DBTransaction) IAccountOwner {
	return novoRepo(tx)
}

func GetService(rReceived IAccountOwner) *Service {
	return &Service{
		repo: rReceived,
	}
}
