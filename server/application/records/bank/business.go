package bank

import (
	"context"

	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/domain/records/bank"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

// SelectPaginated - busca os bancos de foma paginada com base nos query params informados
func SelectPaginated(ctx context.Context, params *utils.ParametrosRequisicao) (res *ResponsePag, err error) {
	var msgErrorDefault = "Erro ao buscar bancos paginados"

	res = new(ResponsePag)
	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	repository := bank.GetRepository(tx)

	list, err := repository.SelectPaginated(params)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	res.Data = make([]Response, len(list.Data))
	for i := 0; i < len(list.Data); i++ {
		if err = utils.ConvertStruct(&list.Data[i], &res.Data[i]); err != nil {
			return res, oops.Wrap(err, msgErrorDefault)
		}
	}

	res.Total, res.Next = list.Total, list.Next
	return
}
