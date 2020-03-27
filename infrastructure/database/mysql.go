package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysql(host string, port int, net string, dbName string, user string, password string) (*sql.DB, error) {
	dbConnection := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, password, net, host, port, dbName)

	dbConn, err := sql.Open("mysql", dbConnection)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
