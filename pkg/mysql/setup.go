package mysql

import (
	"fmt"
	"log"
	"os"
)

// Create Basic admin user
func CreateAdminUser() {
	// Default admin user credentials
	var ADMIN_USER = os.Getenv("ADMIN_USER")
	var ADMIN_PASS = os.Getenv("ADMIN_PASS")

	// Create a default admin user
	var query = fmt.Sprintf("INSERT INTO admin(id,username,password) VALUES(UUID(), '%s', '%s');", ADMIN_USER, ADMIN_PASS)
	var _, err = Db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}

// Create tables if not exist already
func setUpTables() {
	createTable("CREATE TABLE IF NOT EXISTS admin(id uuid PRIMARY KEY, username VARCHAR(20),password CHAR(60));", "admin")

	createTable("CREATE TABLE IF NOT EXISTS product(id uuid PRIMARY KEY, name VARCHAR(100), category CHAR(100), price INT,description text );", "product")

	createTable("CREATE TABLE IF NOT EXISTS stock(product_id uuid, s int, m int, l int, xl int, xxl int, xxxl int, total int, FOREIGN KEY(product_id) REFERENCES product(id));", "stock")

	createTable("CREATE TABLE IF NOT EXISTS orders(id uuid, product_id uuid, name text, address text, house_name text, pin int, city text, phone int, payment text, total int, FOREIGN KEY(product_id) REFERENCES product(id));","order")

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
