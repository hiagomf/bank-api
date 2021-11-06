package database

import (
	"context"
	"database/sql"

	"github.com/hiagomf/bank-api/server/config"
)

// Databases é a interface para multiplos banco de dados
type Databases interface {
	OpenConnection(*config.DataBase) error
	CloseConnection()
	NewTx(context.Context, *sql.TxOptions) (interface{}, error)
}

// Transaction é a interface para transacao dos bancos de dados
type Transaction interface {
	Commit() error
	Rollback() error
}
