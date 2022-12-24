package db_helper

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPGDB(dbUrl string) *sql.DB {
	database, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}
	return database

}
