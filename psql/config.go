package database

import (
	"os"
	"time"

	"github.com/google/uuid"
)

const (
	Driver = "postgres"
	Port   = 5432
)

var (
	Name     = os.Getenv("DB_NAME")
	User     = os.Getenv("DB_USER")
	Host     = os.Getenv("DB_HOST")
	Password = os.Getenv("DB_PASSWORD")
)

type Users struct {
	ID       uuid.UUID
	Name     string
	Username string
	Password string
	Birthday time.Time // maybe needs change
	Avatar   string    // needs change
	UpdateAt time.Time
}
