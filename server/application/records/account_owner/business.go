package account_owner

import (
	"context"
	"regexp"
	"strconv"

	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/domain/records/account_owner"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

// Insert - insere um titular de conta
func Insert(ctx context.Context, req *Request) (id *int64, err error) {
	var msgErrorDefault = "Erro ao inserir titular de conta"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	// validando CPF
	if ok := utils.IsCPF(*req.Document); !ok {
		return id, oops.Wrap(oops.NovoErr("CPF inválido"), msgErrorDefault)
	}

	// Convertendo para camada de infra para não ferir a arquitetura
	ownerRepo := account_owner.GetRepository(tx)
	data, err := ownerRepo.ConvertToInfra(req)
	if err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}

	// retirando sinais do campo de documento
	re := regexp.MustCompile(`(\.|\-)+`)
	cpf := re.ReplaceAllString(*req.Document, "")
	data.Document = &cpf

	// Definindo filtros de busca de registro por documento, garantindo que caso o cliente já exista, deve apenas prosseguir
	var params utils.ParametrosRequisicao
	params.Filtros = make(map[string][]string)
	params.Filtros["document"] = []string{*data.Document}
	params.Limite = 15

	// Buscando lista de usuários com aquele documento: OBS -> campo document é UNIQUE no banco, caso encontre ele retorna o ID
	list, err := ownerRepo.SelectPaginated(&params)
	if err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}

	// Caso o registro já exista exibe um erro
	if len(list.Data) > 0 {
		return data.ID, oops.Wrap(oops.NovoErr("Já existe uma conta com esse CPF"), msgErrorDefault)
	}

	// Inserindo na base
	if err = ownerRepo.Insert(data); err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}

	id = data.ID
	return
}

// Update - altera um titular de conta
func Update(ctx context.Context, req *Request, id *int64) (err error) {
	var msgErrorDefault = "Erro ao alterar titular de conta"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	// validando CPF
	if ok := utils.IsCPF(*req.Document); !ok {
		return oops.Wrap(oops.NovoErr("CPF inválido"), msgErrorDefault)
	}

	// Convertendo para camada de infra para não ferir a arquitetura
	ownerRepo := account_owner.GetRepository(tx)
	data, err := ownerRepo.ConvertToInfra(req)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	// retirando sinais do campo de documento
	re := regexp.MustCompile(`(\.|\-)+`)
	cpf := re.ReplaceAllString(*req.Document, "")
	data.Document = &cpf
	data.ID = id

	// Filtros para validação se CPF é pertencente a outro usuário
	var params utils.ParametrosRequisicao
	params.Filtros = make(map[string][]string)
	params.Filtros["document"] = []string{*data.Document}
	params.Filtros["not_in_id"] = []string{strconv.FormatInt(*data.ID, 10)}
	params.Total = true

	list, err := ownerRepo.SelectPaginated(&params)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	if list.Total != nil && *list.Total != 0 {
		return oops.Wrap(oops.NovoErr("esse CPF pertence a outro titular"), msgErrorDefault)
	}

	// Realizando a alteração
	if err = ownerRepo.Update(data); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	return
}

// SelectOne - busca um tipo de titular de conta com base no ID informado
func SelectOne(ctx context.Context, id *int64) (res *Response, err error) {
	var msgErrorDefault = "Erro ao buscar titular de conta"
	res = new(Response)

	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	repository := account_owner.GetRepository(tx)
	dataInfra, err := repository.SelectOne(id)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	if err = utils.ConvertStruct(dataInfra, res); err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}

	return
}

// SelectPaginated - buscar titulares de conta com base nos query params informados paginados
func SelectPaginated(ctx context.Context, params *utils.ParametrosRequisicao) (res *ResponsePag, err error) {
	var msgErrorDefault = "Erro ao buscar ocorrências paginadas"

	res = new(ResponsePag)
	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	repository := account_owner.GetRepository(tx)
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

// Disable - desativa um titular de conta com base no ID informado
func Disable(ctx context.Context, id *int64) (err error) {
	var msgErrorDefault = "Erro ao desativar conta"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	// desabilitando conta
	if err = account_owner.GetRepository(tx).Disable(id); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	return
}
