package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@/test_golang")
	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Database connected")
	return
}
