package account_owner_address

import "github.com/hiagomf/bank-api/server/config/database"

type Service struct {
	repo IAccountOwnerAddress
}

func GetRepository(tx *database.DBTransaction) IAccountOwnerAddress {
	return novoRepo(tx)
}

func GetService(rReceived IAccountOwnerAddress) *Service {
	return &Service{
		repo: rReceived,
	}
}
