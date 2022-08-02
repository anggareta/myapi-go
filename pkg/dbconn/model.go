package dbconn

import (
	"database/sql"
	"fmt"

	mssql "github.com/denisenkom/go-mssqldb"
)

type DBConnect struct {
	Server   string
	User     string
	Password string
	Port     int
	Database string
}

//type QueryMssql struct {
//	Db *sql.DB
//}

func (m DBConnect) SetConn() (*sql.DB, error) {
	//flag.Parse()

	connString := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%d;database=%s;", m.Server, m.User, m.Password, m.Port, m.Database,
	)
	//db, err := sql.Open("mssql", connString)

	connector, err := mssql.NewConnector(connString)
	// if err != nil {
	// 	return nil, err
	// }
	//connector.SessionInitSQL = "SET ANSI_NULLS ON"

	db := sql.OpenDB(connector)

	return db, err
}
