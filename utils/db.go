package utils

import (
	"database/sql"
	"log"
	"time"
)

func DbConnect(dbURI string) (db *sql.DB, err error) {
	log.Println("Connecting to db... at", dbURI)

	db, err = sql.Open("mysql", dbURI)
	if err != nil {
		log.Println("Unable to connect to db because ", err)
		return
	}

	var _int int
	if err = db.QueryRow("SELECT 1").Scan(&_int); err != nil {
		log.Println("Unable to confirm db connection because ", err)
		return
	}

	db.SetMaxIdleConns(4999)
	db.SetMaxOpenConns(5000)
	db.SetConnMaxLifetime(10 * time.Second)
	log.Println("Db connection successful")
	return
}
