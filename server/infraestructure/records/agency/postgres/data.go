package postgres

import "github.com/hiagomf/bank-api/server/config/database"

type PGAgency struct {
	DB *database.DBTransaction
}
