package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// OpenSQL helps instantiate postgresql service
func OpenSQL() (*sql.DB, error) {
	host := getHost()

	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, Port, User, Password, Name,
	)

	db, err := sql.Open(Driver, dataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getHost() string {
	host := os.Getenv("POSTGRESDB_SERVICE_HOST")
	if host == "" {
		host = os.Getenv("DB_HOST")
	}

	return host
}
