package driver

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	dbURL := "app_authservice:Qwer1234@tcp(mysql-authservice:3306)/authservice"

	db, err := sql.Open("mysql", dbURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
