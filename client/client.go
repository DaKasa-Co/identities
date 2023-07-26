package client

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/DaKasa-Co/identities/model"
	database "github.com/DaKasa-Co/identities/psql"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

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
		infos.Avatar, err = UploadMedia(infos.Avatar)
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

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func UploadMedia(file interface{}) (string, error) {
	name := os.Getenv("CLOUDINARY_CLOUD_NAME")
	key := os.Getenv("CLOUDINARY_API_KEY")
	secret := os.Getenv("CLOUDINARY_API_SECRET")
	folder := os.Getenv("CLOUDINARY_UPLOAD_FOLDER")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(name, key, secret)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: folder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
