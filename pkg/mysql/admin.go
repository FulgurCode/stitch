package mysql

import (
	"github.com/FulgurCode/stitch/models"
)

// Getting admin user with username
func GetAdminUser(username string) (models.Admin, error) {
	var admin models.Admin

	var query = "SELECT username,password from admin;"
	var result, err = Db.Query(query)
	if err != nil {
		return admin, err
	}

	for result.Next() {
		err = result.Scan(&admin.Username, &admin.Password)
	}

	return admin, err
}

// Update admin password
func UpdateAdminPassword(admin models.Admin) error {
	var query = "UPDATE admin set password = ? WHERE username = ?;"

	var _, err = Db.Exec(query, admin.Password, admin.Username)

	return err
}
