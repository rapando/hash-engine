package database

import (
	"database/sql"
	"fmt"
	"sync"
)

func Save(wg *sync.WaitGroup, protocol, rawPassword, hashedPassword string, db *sql.DB) {
	defer wg.Done()
	query := "INSERT INTO passwords (id, protocol, raw_password, hashed_password) " +
		"VALUES (UUID_SHORT(), ?, ?, ?)"
	_, err := db.Exec(query, protocol, rawPassword, hashedPassword)
	fmt.Printf("save : %v\n", err)
}
