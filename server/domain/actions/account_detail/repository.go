package account_detail

import (
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/infraestructure/actions/account_detail"
	"github.com/hiagomf/bank-api/server/infraestructure/actions/account_detail/postgres"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

type repository struct {
	pg *postgres.PGAccountDetail
}

func novoRepo(db *database.DBTransaction) *repository {
	return &repository{
		pg: &postgres.PGAccountDetail{DB: db},
	}
}

// Insert - insere um detalhe de conta
func (r *repository) Insert(data *account_detail.AccountDetail) (err error) {
	return r.pg.Insert(data)
}

// GetAccountDetail - realiza a busca de detalhes de conta, trazendo também a senha para verificação
func (r *repository) GetAccountDetail(data *account_detail.Access) (res *account_detail.AccountDetail, err error) {
	return r.pg.GetAccountDetail(data)
}

// GetDataInfra retorna a model de infra para não ferir o DDD
func (r *repository) GetDataInfra() (res *account_detail.AccountDetail) {
	return new(account_detail.AccountDetail)
}

// GetDataInfra retorna a model de infra para não ferir o DDD
func (r *repository) GetAccessInfra() (res *account_detail.Access) {
	return new(account_detail.Access)
}

// ConvertToInfra - realiza a conversão de alguma model informada para model de infraestructure
func (r *repository) ConvertAcessToInfra(data interface{}) (res *account_detail.Access, err error) {
	res = &account_detail.Access{}

	if err = utils.ConvertStruct(data, res); err != nil {
		return res, oops.Err(err)
	}
	return res, nil
}
