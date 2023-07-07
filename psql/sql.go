package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenSQL() (*sql.DB, error) {
	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Name,
	)

	db, err := sql.Open(Driver, dataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
