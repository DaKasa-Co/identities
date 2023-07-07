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
	ID          uuid.UUID
	Name        string
	Username    string
	Email       string
	Password    string
	Birthday    time.Time
	PhoneNumber int
	Address     string
	Avatar      string
	UpdateAt    time.Time
	CreatedAt   time.Time
}
