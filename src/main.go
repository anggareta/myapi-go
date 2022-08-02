package main

import (
	"fmt"
	"log"
	"net/http"

	"myapi/pkg/api"
	"myapi/pkg/dbconn"
	"myapi/pkg/env"
)

var (
	envConfig map[string]string
	Addr      string
)

func main() {

	envConfig = env.NewEnvConfig(".env")

	Addr = envConfig["server"]
	fmt.Println(Addr)

	dbInfo := dbconn.DBConnect{
		Server:   envConfig["server"],
		User:     envConfig["user"],
		Password: envConfig["password"],
		Port:     api.StringToInt(envConfig["port"]),
		Database: envConfig["database"],
	}
	db, err := dbInfo.SetConn()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected!\n")
	defer db.Close()
	// api.Server = envConfig["server"]
	// api.User = envConfig["user"]
	// api.Password = envConfig["password"]
	// api.Port = api.StringToInt(envConfig["port"])
	// api.Database = envConfig["database"]
	api.ConnDB = api.QueryMssql{
		Db: db,
	}

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
