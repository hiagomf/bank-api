package account_owner_address

import (
	"github.com/hiagomf/bank-api/server/infraestructure/records/account_owner_address"
	"github.com/hiagomf/bank-api/server/utils"
)

type IAccountOwnerAddress interface {
	Insert(data *account_owner_address.AccountOwnerAddress) (err error)
	Update(data *account_owner_address.AccountOwnerAddress) (err error)
	SelectOne(id *int64) (res *account_owner_address.AccountOwnerAddress, err error)
	SelectPaginated(parameters *utils.ParametrosRequisicao) (res *account_owner_address.AccountOwnerAddressPag, err error)
	Disable(id *int64) (err error)
	DisableAllActives(id *int64) (err error)
	ConvertToInfra(data interface{}) (res *account_owner_address.AccountOwnerAddress, err error)
}
