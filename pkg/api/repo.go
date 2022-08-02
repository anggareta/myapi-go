package api

import (
	"database/sql"
	"fmt"

	"myapi/pkg/dbconn"
)

var currentId int

//var todos Todos
//var emp Employee

// var (
// 	server   = "10.82.15.16"
// 	port     = 1433
// 	user     = "sa"
// 	password = "mansek03"
// 	database = "SMS"
// )

type QueryMssql struct {
	Db *sql.DB
}

// func NewQueryMssql(db *sql.DB) IClient {
// 	return &QueryMssql{Db: db}
// }

type MyDB dbconn.DBConnect

// func (m MyDB) SetConn() (*sql.DB, error) {
// 	connString := fmt.Sprintf(
// 		"server=%s;user id=%s;password=%s;port=%d;database=%s;", m.Server, m.User, m.Password, m.Port, m.Database,
// 	)
// 	db, err := sql.Open("mssql", connString)

// 	return db, err
// }

func (qm QueryMssql) getAllEmployee() (res Employees, c int, e error) {
	tsql := fmt.Sprintf("SELECT * FROM Employees;")
	rows, err := qm.Db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res, -1, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		r := Employee{}
		err := rows.Scan(&r.Id, &r.Name, &r.Location, &r.Waktu)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return res, -1, err
		}
		res = append(res, r)
		//fmt.Printf("ID: %s, Name: %s, Location: %s\n", r.Id, r.Name, r.Location)
		count++
	}
	return res, count, nil
}

func (qm QueryMssql) getEmployee(i int) (res Employee, e error) {
	tsql := fmt.Sprintf("SELECT * FROM Employees where ID=%d;", i)
	row := qm.Db.QueryRow(tsql)

	//var Id, Name, Location string
	r := Employee{}
	err := row.Scan(&r.Id, &r.Name, &r.Location, &r.Waktu)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res, err
	}
	// res = Employee{
	// 	Id:       Id,
	// 	Name:     Name,
	// 	Location: Location,
	// }
	fmt.Printf("ID: %s, Name: %s, Location: %s\n", r.Id, r.Name, r.Location)

	return r, nil
}

func (qm QueryMssql) menulis(karyawan Employee) (res string, err error) {
	tsql := fmt.Sprintf("exec [usp_InsertEmployeeTest] %s, %s;", karyawan.Name, karyawan.Location)
	rows, err := qm.Db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		//var name, location string
		var id string
		//r := Employee{}
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return res, err
		}
		res = id
		fmt.Printf("newID: %s\n", id)
		count++
	}
	return res, nil
}
