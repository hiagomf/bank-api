package account

import (
	"github.com/hiagomf/bank-api/server/infraestructure/actions/account"
	"github.com/hiagomf/bank-api/server/utils"
)

type IAccount interface {
	Insert(data *account.Account) (err error)
	Update(data *account.Account) (err error)
	SelectOne(id *int64) (res *account.Account, err error)
	SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account.AccountPag, err error)
	Disable(id *int64) (err error)
	ConvertToInfra(data interface{}) (res *account.Account, err error)
	GetDataInfra() (res *account.Account)
}
