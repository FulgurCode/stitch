package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() {
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASS = os.Getenv("DB_PASS")
	var DB_NAME = os.Getenv("DB_NAME")

	// Creating connection uri and making connection to database
	var connectionUri = fmt.Sprintf("%s:%s@/%s",DB_USER, DB_PASS, DB_NAME)
	var db, err = sql.Open("mysql", connectionUri)
	err = db.Ping()

	if err != nil {
		fmt.Println("FAILED MYSQL CONNECTION:"  + err.Error())
		return
	}

	Db = db
	setUpTables()
}
