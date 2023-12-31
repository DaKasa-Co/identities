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
	Password = os.Getenv("DB_PASSWORD")
	Host     = os.Getenv("DB_HOST")
)

// Users represents Users table
type Users struct {
	ID          uuid.UUID `field:"id"`
	Name        string    `field:"name"`
	Username    string    `field:"username"`
	Email       string    `field:"email"`
	Password    string    `field:"password"`
	Birthday    time.Time `field:"birthday"`
	PhoneNumber int       `field:"phonenumber"`
	Address     string    `field:"address"`
	Avatar      string    `field:"avatar"`
	Stamp       string    `field:"stamp"`
	UpdateAt    time.Time `field:"update_at"`
	CreatedAt   time.Time `field:"created_at"`
}

// Recovery represents recovery table
type Recovery struct {
	ID         uuid.UUID `field:"id"`
	UserID     uuid.UUID `field:"user_id"`
	Validation int       `field:"validation"`
	ExpireAt   time.Time `field:"expire_at"`
}
