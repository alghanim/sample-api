package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/ibmdb/go_ibm_db"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err)
	}

	return db
}

func ConnectToOracle() *sql.DB {
	db, err := sql.Open("godror", "<your username>/<your password>@service_name")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()

	return db
}
