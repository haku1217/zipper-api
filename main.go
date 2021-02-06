package main

import (
	"fmt"
	"log"

	"github.com/haku1217/zipper/controller"
	"github.com/haku1217/zipper/model"
)

func main() {
	const MYSQL_HOST = "127.0.0.1"
	const MYSQL_TCP_PORT = "13322"
	const MYSQL_PWD = "root"
	const MYSQL_USER = "user"
	const DB_NAME = "sample_db"
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USER, MYSQL_PWD, MYSQL_HOST, MYSQL_TCP_PORT, DB_NAME)

	db, err := model.NewMySQL(DSN)
	if err != nil {
		log.Fatal(err)
	}

	controller.Router(db)
}
