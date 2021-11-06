package utils

import (
	"errors"
	"reflect"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

var (
	// ErrFiltroNulo ocorre quando todos os campos estão nulos
	ErrFiltroNulo = errors.New("todos os campos estão nulos")
	// ErrFieldsNull All fields are null

	// ErrTagNaoEncontrada ocorre quando a tag "sql" não é encontrada
	// nos campos de uma estrutura
	ErrTagNaoEncontrada = errors.New("nenhum tag foi escontrada para a estrutura")
)

// ConfigurarPaginacao configura os dados de paginação para um query do banco de dados
// p é a estrutura de requisição de parâmetros que contém o limite e a informação sobre
// os campos, model é a estrutura de dados para os resultados, query é o builder da
// query que será executada, result é o slice com de valores que representam a resposta
// da requisição, next define se há uma próxima página e count define a quantidade de
// linhas disponiveis no resultado
func ConfigurarPaginacao(p *ParametrosRequisicao, model interface{}, query *sq.SelectBuilder, possuiOrdenador ...bool) (result interface{}, next *bool, count *int64, err error) {
	modelType := reflect.Indirect(reflect.ValueOf(model)).Type()
	slice := reflect.MakeSlice(reflect.SliceOf(modelType), 0, 0)
	if p.Total {
		var total int64
		err = query.
			QueryRow().
			Scan(&total)
		if err != nil {
			return
		}

		count = &total
	} else {
		pre := query.
			Limit(p.Limite + 1).
			Offset(p.Offset)

		// é necessário verificar se o erdenador é pre-selecionado e
		// adiciona-lo caso não seja
		if len(possuiOrdenador) != 1 || !possuiOrdenador[0] {
			pre = pre.OrderBy(p.ValidarOrdenador(model))
		}

		rows, err := pre.Query()
		if err != nil {
			return result, next, count, err
		}

		_, values, err := p.ValidarCampos(model)
		if err != nil {
			return result, next, count, err
		}

		slice = reflect.MakeSlice(reflect.SliceOf(modelType), 0, int(p.Limite+1))
		for rows.Next() {
			if err = rows.Scan(values...); err != nil {
				return result, next, count, err
			}
			slice = reflect.Append(slice, reflect.Indirect(reflect.ValueOf(model)))
		}

		hasNext := slice.Len() > int(p.Limite)
		if hasNext {
			slice = slice.Slice(0, int(p.Limite))
		}

		next = &hasNext
	}

	result = slice.Interface()

	return
}

// FormatarInsertUpdate formata colunas e valores para inserir
// ou realizar update
func FormatarInsertUpdate(in interface{}) ([]string, []interface{}, error) {
	var (
		tagFound = false
		cols     []string
		values   []interface{}
	)
	const tagName = "sql"

	elem := reflect.ValueOf(in).Elem()

	if elem.Kind() != reflect.Struct {
		return cols, values, errors.New("não é uma estrutura")
	}

	for s := 0; s < elem.NumField(); s++ {
		if !elem.Field(s).IsNil() {
			elemField := elem.Type().Field(s)

			tag := elemField.Tag.Get(tagName)
			if tag == "" {
				continue
			}

			// remove :: (typecast) from field description
			tag = strings.Split(tag, "::")[0]

			// remove ()
			tag = strings.ReplaceAll(tag, "(", "")
			tag = strings.ReplaceAll(tag, ")", "")

			cols = append(cols, tag)

			value := elem.Field(s).Interface()
			values = append(values, value)

			tagFound = true
		}
	}

	if !tagFound {
		return cols, values, ErrTagNaoEncontrada
	}

	if len(cols) == 0 || len(values) == 0 {
		return cols, values, ErrFiltroNulo
	}

	return cols, values, nil
}

// EncapsularQuery encapsula uma query
func EncapsularQuery(tx sq.BaseRunner, cols []string, sel sq.SelectBuilder) (query sq.SelectBuilder) {
	if len(cols) == 0 {
		cols = []string{"t.*"}
	}

	return sq.Select(cols...).FromSelect(sel, "t").RunWith(tx).PlaceholderFormat(sq.Dollar)
}
