package models

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

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
	Value      interface{}
}
type PPIOQuery struct {
	FromTable    string
	Fields       []string
	JoinTables   []JointTable
	WhereClauses []WhereClause
}

func (ppioQ *PPIOQuery) From(tableName string) {
	ppioQ.FromTable = tableName
}

func (ppioQ *PPIOQuery) Join(tableName string, jointColumnOrig string, jointColumnDest string) {
	if ppioQ.JoinTables == nil {
		ppioQ.JoinTables = make([]JointTable, 0)
	}
	newJointTable := JointTable{
		TableName:       tableName,
		JointColumnOrig: jointColumnOrig,
		JointColumnDest: jointColumnDest,
	}
	ppioQ.JoinTables = append(ppioQ.JoinTables, newJointTable)
}

func (ppioQ *PPIOQuery) Where(columnName string, value interface{}) {
	if ppioQ.WhereClauses == nil {
		ppioQ.WhereClauses = make([]WhereClause, 0)
	}
	newWhereClause := WhereClause{
		ColumnName: columnName,
		Value:      value,
	}
	ppioQ.WhereClauses = append(ppioQ.WhereClauses, newWhereClause)
}

func (ppioQ *PPIOQuery) ToSql() (string, error) {
	var queryBlder bytes.Buffer
	var queryFinal string
	if _, err := queryBlder.WriteString("SELECT "); err != nil {
		log.Printf("Could not create the query with buffer. Error: %v", err)
		return queryFinal, err
	}
	if ppioQ.Fields == nil {
		log.Printf("No field selected for the query.")
		return queryFinal, errors.New("No field selected")
	}
	var fieldsList string
	for _, field := range ppioQ.Fields {
		fieldsList = fmt.Sprintf("%s%s, ", fieldsList, field)
	}
	// Remove extra comma and space for the latest field.
	if _, err := queryBlder.WriteString(fieldsList[:len(fieldsList)-2]); err != nil {
		log.Printf("Could not create the query with buffer. Error: %v", err)
		return queryFinal, err
	}

	if len(ppioQ.FromTable) == 0 {
		log.Printf("From table not given for the query.")
		return queryFinal, errors.New("From table not given")
	}
	/* TODO make sure the fromTable[0] is not redundant with a joint table alias */
	if _, err := queryBlder.WriteString(fmt.Sprintf(" FROM %s %c", ppioQ.FromTable, ppioQ.FromTable[0])); err != nil {
		log.Printf("Could not create the query with buffer. Error: %v", err)
		return queryFinal, err
	}

	/* TODO JOIN TABLES */
	queryBlder.WriteString(" WHERE 1=1 ")
	for _, whereC := range ppioQ.WhereClauses {
		queryBlder.WriteString(fmt.Sprintf("AND %s=? ", whereC.ColumnName))
	}

	return queryBlder.String(), nil
}
