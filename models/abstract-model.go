package models

import "database/sql"

type IDbObject interface {
	Insert(dbConn *sql.DB) int64
	GetByID(dbConn *sql.DB) error
	GetAll(dbConn *sql.DB) ([]IDbObject, error)
}
