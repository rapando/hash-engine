package models

import (
	"database/sql"
	"fmt"
	"sync"
)

func Save(wg *sync.WaitGroup, protocol, rawPassword, hashedPassword string, db *sql.DB) {
	defer wg.Done()
	query := "INSERT INTO passwords (protocol, raw_password, hashed_password) " +
		"VALUES (?, ?, ?)"
	_, err := db.Exec(query, protocol, rawPassword, hashedPassword)
	fmt.Printf("save : %v\n", err)
}
