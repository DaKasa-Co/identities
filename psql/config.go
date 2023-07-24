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
)

type Users struct {
	ID          uuid.UUID `field:"id"`
	Name        string    `field:"name"`
	Username    string    `field:"username"`
	Email       string    `field:"email"`
	Password    string    `field:"password"`
	Birthday    time.Time `field:"birthday"`
	PhoneNumber int       `field:"phonenumber"`
	Address     string    `field:"address"`
	Avatar      string    `field:"picture"`
	UpdateAt    time.Time `field:"timestamp_update"`
	CreatedAt   time.Time `field:"timestamp_created"`
}
