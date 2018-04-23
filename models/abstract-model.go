package models

import "database/sql"

type IDbObject interface {
	Insert(dbConn *sql.DB) int64
	GetByID(dbConn *sql.DB) error
	GetAll(dbConn *sql.DB) ([]IDbObject, error)
}

type Query interface {
	From(tableName string)
	Join(tableName string, jointColumnOrig string, jointColumnDest string)
	Where(columnName string, value interface{})
	ToSql() (string, error)
}

type JointTable struct {
	TableName       string
	JointColumnOrig string
	JointColumnDest string
}

type WhereClause struct {
	ColumnName string
	value      interface{}
}
type PPIOQuery struct {
	FromTable  string
	JoinTables []JointTable
	Where      []WhereClause
}
