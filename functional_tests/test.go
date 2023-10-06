package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	database "github.com/DaKasa-Co/identities/psql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	TestRegister()
	TestRecovery()
	TestChall()
	TestLogin()
}

func getFirstGeneratedTicket() (uuid.UUID, int) {
	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), 5432, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT id, validation FROM recovery;"
	res, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()
	rows := []database.Recovery{}
	for res.Next() {
		row := new(database.Recovery)

		err = res.Scan(&row.ID, &row.Validation)
		if err != nil {
			log.Fatal(err)
		}

		rows = append(rows, *row)
	}

	return rows[0].ID, rows[0].Validation
}
