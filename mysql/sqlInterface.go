package mysql

import (
	"database/sql"
)

type SqlSvrInterface interface {
	Query(sqlStr string) (*sql.Rows, error)
	Exec(sqlStr string) (*sql.Result, error)
	Prepare(sqlStr string) (*sql.Stmt, error)
	DoCmd() (*sql.Tx, error)
	Close()
}
