package account_detail

import "github.com/hiagomf/bank-api/server/infraestructure/actions/account_detail"

type IAccountDetail interface {
	Insert(data *account_detail.AccountDetail) (err error)
	GetAccountDetail(data *account_detail.Access) (res *account_detail.AccountDetail, err error)
	GetDataInfra() (res *account_detail.AccountDetail)
	ConvertAcessToInfra(data interface{}) (res *account_detail.Access, err error)
	GetAccessInfra() (res *account_detail.Access)
}
