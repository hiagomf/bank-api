package postgres

import "github.com/hiagomf/bank-api/server/config/database"

type PGAccount struct {
	DB *database.DBTransaction
}
