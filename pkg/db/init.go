package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
}
