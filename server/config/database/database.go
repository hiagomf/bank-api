package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hiagomf/bank-api/server/config"
	"github.com/hiagomf/bank-api/server/config/database/postgres"
	"go.uber.org/zap"
)

const (
	// TimeoutWarningQuery define a mensagem padrão para um aviso
	// de que o tempo limite de uma consulta foi atingido
	TimeoutWarningQuery = "Tempo esperado para consulta foi excedido"
)

var (
	connections map[string]Databases
	roSuffix    = "-ro"
)

// DBTransaction é usado para agregar transações para todos os
// banco de dados disponíveis
type DBTransaction struct {
	postgres *sql.Tx
	ctx      context.Context
	Builder  sq.StatementBuilderType
}

// OpenConnection itera sobre os banco de dados indicados na configuração
// e tenta abrir uma conexão com eles
func OpenConnection() error {
	initConnectionsMap()

	conf := config.GetConfig()

	for d := 0; d < len(conf.Databases); d++ {
		connName := conf.Databases[d].Nick
		if conf.Databases[d].ReadOnly {
			connName += roSuffix
		}
		if _, set := connections[connName]; set {
			if err := connections[connName].OpenConnection(&conf.Databases[d]); err != nil {
				return err
			}
		}
	}

	return nil
}

// CloseConnections intera sobre o mapa de conexões abertas e as fecha
func CloseConnections() {
	for _, v := range connections {
		v.CloseConnection()
	}
}

func initConnectionsMap() {
	if connections != nil {
		return
	}
	connections = make(map[string]Databases)
	connections["bank-api"] = &postgres.Postgres{}
	connections["bank-api"+roSuffix] = &postgres.Postgres{}
}

// NewTransaction tenta abrir uma nova transacao para a conexão
// com o banco de dados do bank-api
func NewTransaction(ctx context.Context, readOnly bool) (*DBTransaction, error) {
	t := &DBTransaction{}
	db := "bank-api"
	if readOnly {
		db += "-ro"
	}

	pgsql, err := connections[db].NewTx(ctx, &sql.TxOptions{
		ReadOnly:  readOnly,
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, err
	}

	t.postgres = pgsql.(*sql.Tx)
	t.Builder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(t)
	t.ctx = ctx

	return t, nil
}

// Commit aplica as alterações pendentes em uma transacao
func (t *DBTransaction) Commit() (err error) {
	err = t.postgres.Commit()
	return
}

// Rollback desfaz todas as alterações pendentes em uma transacao
func (t *DBTransaction) Rollback() {
	_ = t.postgres.Rollback()
}

// Exec implementa a interface do método Exec
func (t *DBTransaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.postgres.Exec(getCaller()+query, args...)
}

// QueryRow implementa a interface do método QueryRow
func (t *DBTransaction) QueryRow(query string, args ...interface{}) sq.RowScanner {
	return t.postgres.QueryRow(getCaller()+query, args...)
}

// Query implementa a interface do método Query
func (t *DBTransaction) Query(query string, args ...interface{}) (*sql.Rows, error) {
	ch := make(chan bool, 1)
	conf := config.GetConfig()

	go func() {
		select {
		case <-time.After(time.Duration(conf.QueryTimeout * float32(time.Second))):
			logMsg := map[string]interface{}{
				"aviso":    TimeoutWarningQuery,
				"consulta": query,
				"valores":  args,
			}

			msg, err := json.Marshal(logMsg)
			if err != nil {
				log.Printf("%+v\n", logMsg)
			} else {
				log.Println(string(msg))
			}

			return
		case <-ch:
			return
		}
	}()

	r, err := t.postgres.Query(getCaller()+query, args...)
	ch <- true
	return r, err
}

// CloseQuery fecha a conexão da query
func (t *DBTransaction) CloseQuery(rows io.Closer) {
	err := rows.Close()
	if err != nil {
		zap.L().Error("DATABASE [database.CloseQuery(rows *sql.Rows)]", zap.Error(err))
	}
}

func getCaller() (saida string) {
	var (
		archive string
		line    int
	)

	saida = "-- \r\n"

	for i := 3; i < 8; i++ {
		_, archive, line, _ = runtime.Caller(i)
		if strings.Contains(archive, "/services/") {
			saida += "-- " + strconv.FormatInt(int64(i), 10) + " :: " + archive + ":" + strconv.FormatInt(int64(line), 10) + "\r\n"
		}
	}

	return
}
