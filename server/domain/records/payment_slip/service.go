package payment_slip

import "github.com/hiagomf/bank-api/server/config/database"

type Service struct {
	repo IPaymentSlip
}

func GetRepository(tx *database.DBTransaction) IPaymentSlip {
	return novoRepo(tx)
}

func GetService(rReceived IPaymentSlip) *Service {
	return &Service{
		repo: rReceived,
	}
}
