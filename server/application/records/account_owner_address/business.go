package account_owner_address

import (
	"context"
	"regexp"
	"strconv"

	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/domain/records/account_owner"
	"github.com/hiagomf/bank-api/server/domain/records/account_owner_address"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

// Insert - insere um endereço relacionado a um titula, desabilitando os antigos já cadastrados
func Insert(ctx context.Context, req *Request) (id *int64, err error) {
	var msgErrorDefault = "Erro ao inserir endereço de titular de conta"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	// Validando se o titular existe
	if _, err := account_owner.GetRepository(tx).SelectOne(req.OwnerID); err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}

	// Convertendo para camada de infra para não ferir a arquitetura
	ownerAddressRepo := account_owner_address.GetRepository(tx)
	data, err := ownerAddressRepo.ConvertToInfra(req)
	if err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}

	// retirando sinais do campo de CEP
	re := regexp.MustCompile(`(\.|\-)+`)
	cep := re.ReplaceAllString(*req.ZipCode, "")
	data.ZipCode = &cep

	var params utils.ParametrosRequisicao
	params.Filtros = make(map[string][]string)
	params.Filtros["zip_code"] = []string{*data.ZipCode}
	params.Filtros["public_place"] = []string{*data.PublicPlace}
	params.Filtros["number"] = []string{*data.Number}
	params.Filtros["district"] = []string{*data.District}
	params.Filtros["city"] = []string{*data.City}
	params.Filtros["state"] = []string{*data.State}
	params.Filtros["country"] = []string{*data.Country}
	params.Filtros["account_owner_id"] = []string{strconv.FormatInt(*data.OwnerID, 10)}
	params.Filtros["deleted"] = []string{"false"}
	params.Limite = 15

	// Buscando se existe registro com os dados a cima
	list, err := ownerAddressRepo.SelectPaginated(&params)
	if err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}

	// Caso o registro já exista exibe um erro
	if len(list.Data) > 0 {
		return data.ID, oops.Wrap(oops.NovoErr("endereço já registrado"), msgErrorDefault)
	}

	// Desabilitando registros ativos anteriores
	if err = ownerAddressRepo.DisableAllActives(data.OwnerID); err != nil {
		return id, oops.Wrap(err, msgErrorDefault)
	}

	// Inserindo na base
	if err = ownerAddressRepo.Insert(data); err != nil {
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
	var msgErrorDefault = "Erro ao alterar endereço de titular de conta"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	// Validando se o titular existe
	if _, err := account_owner.GetRepository(tx).SelectOne(req.OwnerID); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	// Convertendo para camada de infra para não ferir a arquitetura
	ownerAddressRepo := account_owner_address.GetRepository(tx)
	data, err := ownerAddressRepo.ConvertToInfra(req)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	// retirando sinais do campo do CEP
	re := regexp.MustCompile(`(\.|\-)+`)
	cep := re.ReplaceAllString(*req.ZipCode, "")
	data.ZipCode = &cep
	data.ID = id

	var params utils.ParametrosRequisicao
	params.Filtros = make(map[string][]string)
	params.Filtros["zip_code"] = []string{*data.ZipCode}
	params.Filtros["public_place"] = []string{*data.PublicPlace}
	params.Filtros["number"] = []string{*data.Number}
	params.Filtros["district"] = []string{*data.District}
	params.Filtros["city"] = []string{*data.City}
	params.Filtros["state"] = []string{*data.State}
	params.Filtros["country"] = []string{*data.Country}
	params.Filtros["account_owner_id"] = []string{strconv.FormatInt(*data.OwnerID, 10)}
	params.Filtros["deleted"] = []string{"false"}
	params.Total = true

	// Buscando se existe registro com os dados a cima
	list, err := ownerAddressRepo.SelectPaginated(&params)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	// impedindo caso o endereço já esteja devidamente registrado
	if list.Total != nil && *list.Total != 0 {
		return oops.Wrap(oops.NovoErr("endereço já registrado"), msgErrorDefault)
	}

	// Realizando a alteração
	if err = ownerAddressRepo.Update(data); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	return
}

// SelectOne - busca um tipo de titular de conta com base no ID informado
func SelectOne(ctx context.Context, id *int64) (res *Response, err error) {
	var msgErrorDefault = "Erro ao buscar endereço de titular de conta"
	res = new(Response)

	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	repository := account_owner_address.GetRepository(tx)
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
	var msgErrorDefault = "Erro ao buscar endereço de titular de conta"

	res = new(ResponsePag)
	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	repository := account_owner_address.GetRepository(tx)
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
	var msgErrorDefault = "Erro ao desativar  endereço de titular de conta"

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}
	defer tx.Rollback()

	// desabilitando conta
	if err = account_owner_address.GetRepository(tx).Disable(id); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrorDefault)
	}

	return
}
