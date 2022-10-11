package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB = Connect()

func Connect() *sql.DB {
	DB, err := sql.Open("mysql", myEnv["DB_USER"]+":"+myEnv["DB_PASS"]+"@tcp("+myEnv["DB_URL"]+":"+myEnv["DB_PORT"]+")/"+myEnv["DB_NAME"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to Mysql")
	return DB
}

func Close() {
	DB.Close()
}
