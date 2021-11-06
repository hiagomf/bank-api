package bank

import (
	"github.com/hiagomf/bank-api/server/infraestructure/records/bank"
	"github.com/hiagomf/bank-api/server/utils"
)

type IBank interface {
	SelectPaginated(parameters *utils.ParametrosRequisicao) (res *bank.BankPag, err error)
}
