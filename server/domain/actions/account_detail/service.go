package account_detail

import "github.com/hiagomf/bank-api/server/config/database"

type Service struct {
	repo IAccountDetail
}

func GetRepository(tx *database.DBTransaction) IAccountDetail {
	return novoRepo(tx)
}

func GetService(rReceived IAccountDetail) *Service {
	return &Service{
		repo: rReceived,
	}
}
