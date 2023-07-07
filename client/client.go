package client

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckRequiredFieldExists(key string, value interface{}) error {
	if value != nil {
		return fmt.Errorf("field %s is required", key)
	}

	return nil
}

func CheckIsValidEmail(email string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegex.MatchString(email) {
		return fmt.Errorf("respective email address '%s' is not valid", email)
	}

	return nil
}

func CheckIfNotHasSpecialCharacters(key string, value string) error {
	r := regexp.MustCompile(`^[A-Za-zÀ-ÿ\s]+$`)

	if !r.MatchString(value) {
		return fmt.Errorf("special characters not allowed in field %s", key)
	}

	return nil
}

func CheckIsValidUsername(username string) error {
	r := regexp.MustCompile(`^[\.\-\w]{3,13}$`)

	if !r.MatchString(username) {
		return fmt.Errorf("respective username '%s' is not valid", username)
	}

	return nil
}

func CheckIsValidPassword(password string) error {
	r := regexp.MustCompile(`^.{8,225}$`)

	if !r.MatchString(password) {
		return fmt.Errorf("respective password is not valid")
	}

	return nil
}

func CheckBirthday(birth time.Time) error {
	start := time.Now().AddDate(-213, 0, 0)
	end := time.Now().AddDate(-13, 0, 0)

	if birth.Before(start) || birth.After(end) {
		return fmt.Errorf("respective birthday %v is not valid", birth.Format(time.DateOnly))
	}

	return nil
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
