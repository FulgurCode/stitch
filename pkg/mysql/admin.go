package mysql

import (
	"fmt"

	"github.com/VAISHAKH-GK/benster-website/models"
)

// Getting admin user with username
func GetAdminUser(username string) (models.Admin, error) {
	var admin models.Admin

	var query = fmt.Sprintf("SELECT username,password from admin;")
	var result, err = Db.Query(query)
	if err != nil {
		return admin, err
	}

	for result.Next() {
		err = result.Scan(&admin.Username, &admin.Password)
	}

	return admin, err
}
