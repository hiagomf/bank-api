package account

import (
	"github.com/hiagomf/bank-api/server/infraestructure/records/account"
	"github.com/hiagomf/bank-api/server/utils"
)

type IAccount interface {
	Insert(data *account.Account) (err error)
	Update(data *account.Account) (err error)
	SelectOne(id *string) (res *account.Account, err error)
	SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account.AccountPag, err error)
	Disable(id *string) (err error)
}
