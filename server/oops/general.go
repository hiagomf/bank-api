package oops

import "errors"

var (
	// ErrDadoNaoEncontrado define que um dado esperado
	// não pôde ser encontrado
	ErrDadoNaoEncontrado = Error{
		Msg:        "Erro interno: não foi possível obter dado requerido para continuar operações",
		Code:       internalCode,
		StatusCode: 400,
		Err:        errors.New("erro interno: não foi possível obter dado requerido para continuar operações"),
	}

	// ErrFiltroInvalido indica que um filtro definido em uma
	// rota tem o tipo inválido
	ErrFiltroInvalido = Error{
		Msg:        "Valor especificado em filtro de rota tem tipo inválido",
		Code:       defaultCode,
		StatusCode: 400,
		Err:        errors.New("valor especificado em filtro de rota tem tipo inválido"),
	}

	// ErrMemcachedConn indica que a conexão com o memcache foi
	// mal sucedida
	ErrMemcachedConn = Error{
		Msg:        "Nenhuma conexão com o memcached aberta",
		Code:       internalCode,
		StatusCode: 500,
		Err:        errors.New("nenhuma conexão com o memcached aberta"),
	}

	// ErrInvalidPlatform indica que a plataforma usada
	// nao é valida
	ErrInvalidPlatform = Error{
		Msg:        "Acesso a partir de plataforma inválida",
		Code:       defaultCode,
		StatusCode: 409,
		Err:        errors.New("acesso a partir de plataforma inválida"),
	}
)
