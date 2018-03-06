package models

import "database/sql"

type IDbObject interface {
	Insert(dbConn *sql.DB) int64
}
