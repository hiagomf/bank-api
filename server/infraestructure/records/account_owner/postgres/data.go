package postgres

import "github.com/hiagomf/bank-api/server/config/database"

type PGAccountOwner struct {
	DB *database.DBTransaction
}
