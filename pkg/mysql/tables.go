package mysql

import (
	"fmt"
	"log"
)

func createTable(query string, table string) {
	var _, err = Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s Table creation success\n", table)
	}
}
