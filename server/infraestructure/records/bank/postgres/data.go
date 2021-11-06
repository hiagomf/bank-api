package postgres

import "github.com/hiagomf/bank-api/server/config/database"

type PGBank struct {
	DB *database.DBTransaction
}
