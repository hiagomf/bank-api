package agency

import "github.com/hiagomf/bank-api/server/config/database"

type Service struct {
	repo IAgency
}

func GetRepository(tx *database.DBTransaction) IAgency {
	return novoRepo(tx)
}

func GetService(rReceived IAgency) *Service {
	return &Service{
		repo: rReceived,
	}
}
