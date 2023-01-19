package database

import (
	"database/sql"
	"log"

	"github.com/abc3354/CODEV-back/services/env"
)

var db *sql.DB

func Get() *sql.DB {
	if db == nil {
		openingDB, err := sql.Open("postgres", env.Get().Datasource)
		if err != nil {
			log.Fatal(err)
		}
		db = openingDB
	}
	return db
}

func Close() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
