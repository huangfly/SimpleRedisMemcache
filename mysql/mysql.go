package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "sync"
)

type DataBase struct {
	dbsql *sql.DB
}

//创建数据库
func NewDataBase(dbtype string, dbinfo string) SqlSvrInterface {
	dbbase, err := sql.Open(dbtype, dbinfo)
	if err != nil {
		log.Println("NewDataBase open error : ", err)
		return nil
	}
	err = dbbase.Ping()
	if err != nil {
		log.Println("NewDataBase ping error : ", err)
		return nil
	}

	return &DataBase{dbsql: dbbase}
}

//查询数据库，返回查询到的rows
func (this *DataBase) Query(sqlStr string) (*sql.Rows, error) {
	rows, err := this.dbsql.Query(sqlStr)
	return rows, err
}

//执行sql语句通常用于insert，update等操作，返回sql.Result
func (this *DataBase) Exec(sqlStr string) (*sql.Result, error) {
	result, err := this.dbsql.Exec(sqlStr)
	return &result, err
}

//预备处理sql语句，返回sql.Stmt
func (this *DataBase) Prepare(sqlStr string) (*sql.Stmt, error) {
	stmt, err := this.dbsql.Prepare(sqlStr)
	return stmt, err
}

//事物语句，返回sql.Tx
func (this *DataBase) DoCmd() (*sql.Tx, error) {
	tx, err := this.dbsql.Begin()
	return tx, err
}

//关闭数据库
func (this *DataBase)Close(){
	this.dbsql.Close()
}