package mysql

import (
	"fmt"
	"log"
	"os"
)

// Create Basic admin user
func CreateAdminUser() {
	// Default admin user credentials
	var ADMIN_USER = os.Getenv("DB_ADMIN_USER")
	var ADMIN_PASS = os.Getenv("DB_ADMIN_PASS")

	// Create a default admin user
	var query = fmt.Sprintf("INSERT INTO admin(id,username,password) VALUES(UUID(), '%s', '%s');", ADMIN_USER, ADMIN_PASS)
	var _, err = Db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}

// Create tables if not exist already
func createTables() {
	var _, err = Db.Exec("CREATE TABLE IF NOT EXISTS admin(id uuid, username VARCHAR(20),password VARCHAR(51));")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("table creation success")
	}

	// Get current number of admin users
	result, err := Db.Query("SELECT COUNT(*) FROM admin;")
	if err != nil {
		log.Fatal(err)
	}

	result.Next()
	var length int
	result.Scan(&length)

	// Create new admin user if none exist already
	if length == 0 {
		CreateAdminUser()
	}
}
