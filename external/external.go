package external

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"gopkg.in/mail.v2"
)

// MediaStorage represents the media storage
type MediaStorage struct {
	Name   string
	Key    string
	Secret string
	Folder string
}

// Email represents the email service
type Email struct {
	Host      string
	Port      string
	Account   EmailAccount
	Receivers []string
	Content   []byte
}

// EmailAccount represents the service account responsible for sending emails
type EmailAccount struct {
	Email    string
	Password string
}

var (
	LoadedStorage = MediaStorage{
		Name:   os.Getenv("CLOUDINARY_CLOUD_NAME"),
		Key:    os.Getenv("CLOUDINARY_API_KEY"),
		Secret: os.Getenv("CLOUDINARY_API_SECRET"),
		Folder: os.Getenv("CLOUDINARY_UPLOAD_FOLDER"),
	}

	LoadedEmail = Email{
		Host: os.Getenv("EMAIL_HOST"),
		Port: os.Getenv("EMAIL_PORT"),
		Account: EmailAccount{
			Email:    os.Getenv("EMAIL_USER_ACCOUNT"),
			Password: os.Getenv("EMAIL_PASSWORD_ACCOUNT"),
		},
	}
)

type EmailSender interface {
	SendEmailToRecoverAccount(receivers []string, validation string) error
}

type MediaStorageHandler interface {
	UploadMedia(file interface{}) (string, error)
}

// SendEmailToRecoveryAccount sends recovery message to owner's account
func (e Email) SendEmailToRecoverAccount(receivers []string, validation string) error {
	port, err := strconv.Atoi(e.Port)
	if err != nil {
		return err
	}

	m := mail.NewMessage()
	m.SetHeader("From", e.Account.Email)
	m.SetHeader("To", receivers...)
	m.SetHeader("Subject", "Recuperação da conta DaKasa")
	m.SetBody("text/html", "A senha de recuperação da conta é:<br><hr><h1 style='text-align:center;'>"+validation+"</h1>")

	d := mail.NewDialer(e.Host, port, e.Account.Email, e.Account.Password)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	err = d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}

// UploadMedia upload to cloud respective media
func (s *MediaStorage) UploadMedia(file interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(s.Name, s.Key, s.Secret)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: s.Folder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
