package agency

import (
	"context"
	"strconv"

	bankApp "github.com/hiagomf/bank-api/server/application/records/bank"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/domain/records/agency"
	"github.com/hiagomf/bank-api/server/domain/records/bank"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

// SelectPaginatedByBank - busca as agências de banco de foma paginada com base nos query params informados
func SelectPaginatedByBank(ctx context.Context, bankID *int64, params *utils.ParametrosRequisicao) (res *ResponsePag, err error) {
	var msgErrorDefault = "Erro ao buscar agências paginadas"

	res = new(ResponsePag)
	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	// chamando repositórios
	repositoryB := bank.GetRepository(tx)
	repositoryA := agency.GetRepository(tx)

	// buscando banco na base, caso não encontre, retorna erro
	bank, err := repositoryB.SelectOne(bankID)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	// definindo banco no parâmetro de busca no banco de dado
	params.Filtros["bank_id"] = []string{strconv.FormatInt(*bank.ID, 10)}
	list, err := repositoryA.SelectPaginated(params)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	// preenchendo dados do banco
	res.Bank = new(bankApp.Response)
	if err = utils.ConvertStruct(bank, res.Bank); err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	// populando as agências
	res.Agencies = make([]Response, len(list.Data))
	for i := 0; i < len(list.Data); i++ {
		if err = utils.ConvertStruct(&list.Data[i], &res.Agencies[i]); err != nil {
			return res, oops.Wrap(err, msgErrorDefault)
		}
	}

	res.Total, res.Next = list.Total, list.Next
	return
}
