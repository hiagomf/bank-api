package account_owner

import (
	"github.com/hiagomf/bank-api/server/infraestructure/records/account_owner"
	"github.com/hiagomf/bank-api/server/utils"
)

type IAccountOwner interface {
	Insert(data *account_owner.AccountOwner) (err error)
	Update(data *account_owner.AccountOwner) (err error)
	SelectOne(id *int64) (res *account_owner.AccountOwner, err error)
	SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account_owner.AccountOwnerPag, err error)
	Disable(id *int64) (err error)
	ConvertToInfra(data interface{}) (res *account_owner.AccountOwner, err error)
}
