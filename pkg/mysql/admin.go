package mysql

import (
	"fmt"

	"github.com/VAISHAKH-GK/benster-website/models"
)

// Getting admin user with username
func GetAdminUser(username string) (error, models.Admin) {
	var admin models.Admin

	var query = fmt.Sprintf("SELECT username,password from admin;")
	var result, err = Db.Query(query)
	if err != nil {
		return err, admin
	}

	for result.Next() {
		err = result.Scan(&admin.Username, &admin.Password)
	}

	return err, admin
}
