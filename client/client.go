package client

import (
	"fmt"
	"regexp"
	"time"
	"unicode"

	"github.com/DaKasa-Co/identities/external"
	"github.com/DaKasa-Co/identities/model"
	database "github.com/DaKasa-Co/identities/psql"
	"github.com/gin-gonic/gin"
)

// CheckRequiredFieldExists checks if required field has been writed
func CheckRequiredFieldExists(key string, value interface{}) error {
	var res bool

	switch value.(type) {
	case string:
		res = value == ""
	case int:
		res = value == 0
	default:
		res = value == nil
	}

	if res {
		return fmt.Errorf("field %s is required", key)
	}

	return nil
}

// CheckIsValidEmail checks if input string is a valid email
func CheckIsValidEmail(email string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegex.MatchString(email) {
		return fmt.Errorf("respective email address '%s' is not valid", email)
	}

	return nil
}

// CheckIfNotHasSpecialCharacters checks if respective field dont has special characters
func CheckIfNotHasSpecialCharacters(key string, value string) error {
	r := regexp.MustCompile(`^[A-Za-zÀ-ÿ\s]+$`)

	if !r.MatchString(value) {
		return fmt.Errorf("special characters not allowed in field %s", key)
	}

	return nil
}

// CheckIsValidUsername check if respective username is valid
func CheckIsValidUsername(username string) error {
	r := regexp.MustCompile(`^[\.\-\w]{3,13}$`)

	if !r.MatchString(username) {
		return fmt.Errorf("respective username '%s' is not valid", username)
	}

	return nil
}

// CheckIsValidPassword checks if respective password is valid
func CheckIsValidPassword(password string) error {
	if len(password) >= 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			return fmt.Errorf("password must have at least one capital letter")
		case unicode.IsLower(char):
			return fmt.Errorf("password must have at least one lowercase letter")
		case unicode.IsNumber(char):
			return fmt.Errorf("password must have at least one number")
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			return fmt.Errorf("password must have at least one special character")
		}
	}

	return nil
}

// CheckBirthday check if respective user has an age between 13 ~ 213 years
func CheckBirthday(birth time.Time) error {
	start := time.Now().AddDate(-213, 0, 0)
	end := time.Now().AddDate(-13, 0, 0)

	if birth.Before(start) || birth.After(end) {
		return fmt.Errorf("respective birthday %v is not valid", birth.Format(time.DateOnly))
	}

	return nil
}

// PrepareUserRegisterDatas validates sent user datas and prepares to register in service
func PrepareUserRegisterDatas(infos model.Identity) (database.Users, error) {
	requiredFields := map[string]interface{}{
		"name":     infos.Name,
		"username": infos.Username,
		"email":    infos.Email,
		"password": infos.Password,
		"birthday": infos.Birthday,
		"phone":    infos.PhoneNumber,
	}

	for k, v := range requiredFields {
		if err := CheckRequiredFieldExists(k, v); err != nil {
			return database.Users{}, err
		}
	}

	birthday, err := time.Parse(time.DateOnly, infos.Birthday)
	if err != nil {
		return database.Users{}, err
	}

	if err = CheckIfNotHasSpecialCharacters("name", infos.Name); err != nil {
		return database.Users{}, err
	}

	if err = CheckIsValidUsername(infos.Username); err != nil {
		return database.Users{}, err
	}

	if err = CheckIsValidEmail(infos.Email); err != nil {
		return database.Users{}, err
	}

	if err = CheckIsValidPassword(infos.Password); err != nil {
		return database.Users{}, err
	}

	if err = CheckBirthday(birthday); err != nil {
		return database.Users{}, err
	}

	if infos.Avatar != "" {
		infos.Avatar, err = external.LoadedStorage.UploadMedia(infos.Avatar)
		if err != nil {
			return database.Users{}, err
		}
	}

	return database.Users{
		Name:        infos.Name,
		Username:    infos.Username,
		Email:       infos.Email,
		Password:    infos.Password,
		Birthday:    birthday,
		PhoneNumber: infos.PhoneNumber,
		Address:     infos.Address,
		Avatar:      infos.Avatar,
	}, nil
}

// ErrorResponse is a auxiliary error response handler
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
