package main

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	database "github.com/DaKasa-Co/identities/psql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var ID uuid.UUID
var Validation int

func getFirstGeneratedTicket() (uuid.UUID, int) {
	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), 5432, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}

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

func registerMissingRequiredFields() error {
	fmt.Println("Send HTTP Request without some request field. Should return error")
	body := []byte(`{"username": "teste", "email": "someemail@gmail.com"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/register", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 400 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func registerSuccess() error {
	fmt.Println("Send HTTP Request with all datas correct. Should return success")
	body := []byte(`{"name": "gio silva", "username": "gio._.", "email": "someemail@gmail.com", "password": "someGoodPassword123%", "birthday": "2003-05-11", "phoneNumber": 1291820931901823}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/register", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func registerConflictWithUsernameFieldThatAlreadyExists() error {
	fmt.Println("Send HTTP Request with username and email that already exists in database. Should return error")
	body := []byte(`{"name": "gio silva", "username": "gio._.", "email": "someemasil@gmail.com", "password": "someGoodPassword123%", "birthday": "2003-05-11", "phoneNumber": 12918209321901823}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/register", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 409 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func registerConflictWithEmailFieldThatAlreadyExists() error {
	fmt.Println("Send HTTP Request with username and email that already exists in database. Should return error")
	body := []byte(`{"name": "gio silva", "username": "gio._..", "email": "someemail@gmail.com", "password": "someGoodPassword123%", "birthday": "2003-05-11", "phoneNumber": 12918230931901823}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/register", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 409 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func registerConflictWithPhoneNumberFieldThatAlreadyExists() error {
	fmt.Println("Send HTTP Request with username and email that already exists in database. Should return error")
	body := []byte(`{"name": "gio silva", "username": "gio._..", "email": "someemasil@gmail.com", "password": "someGoodPassword123%", "birthday": "2003-05-11", "phoneNumber": 1291820931901823}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/register", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 409 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func recoveryErrorUserNotFound() error {
	fmt.Println("Send HTTP Request with phonenumber that doens't exist in user database. Should return error")
	body := []byte(`{"phoneNumber": 12345}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 404 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func recoverySuccessWithPhoneNumber() error {
	fmt.Println("Create ticket recovery with phone number. Should return success.")
	body := []byte(`{"phoneNumber": 1291820931901823}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func recoverySuccessWithEmail() error {
	fmt.Println("Create ticket recovery with email. Should return success.")
	body := []byte(`{"email": "someemail@gmail.com"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func recoverySuccessWithUsername() error {
	fmt.Println("Create ticket recovery with username. Should return success.")
	body := []byte(`{"username": "gio._."}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func challRecoveryErrorBadPassword() error {
	fmt.Println("New bad password. Should return error")
	body := []byte(`{"password": "BadPass123"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/chall-recovery", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 400 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func challRecoveryWrongValidation() error {
	fmt.Println("Validation number is incorrect. Should return error")
	body := []byte(`{"password": "BadPass123#", "status": {"ticket": "` + ID.String() + `", "validation": {"tmp": 222}}}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/chall-recovery", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 403 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func challRecoverySuccess() error {
	fmt.Println("Recovery challenge with positive result. Should return success")
	body := []byte(`{"password": "GoodPass123#", "status": {"ticket": "` + ID.String() + `", "validation": {"tmp": ` + strconv.Itoa(Validation) + `}}}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/chall-recovery", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func loginWrongCredentials() error {
	fmt.Println("Send HTTP Request with bad credentials. Should return error")
	body := []byte(`{"email": "someemail@gmail.com", "password": "someBadPassword"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/login", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 403 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func loginSuccess() error {
	fmt.Println("Send HTTP Request with correct credentials. Should return success and JWT token in the header")
	body := []byte(`{"email": "someemail@gmail.com", "password": "GoodPass123#"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/login", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New(string(body))
	}

	if res.Header.Get("X-JWT") == "" {
		fmt.Printf("Headers: %v\n", res.Header)
		return errors.New("JWT header is required")
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func main() {
	err := registerMissingRequiredFields()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = registerSuccess()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = registerConflictWithUsernameFieldThatAlreadyExists()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = registerConflictWithEmailFieldThatAlreadyExists()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = registerConflictWithPhoneNumberFieldThatAlreadyExists()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = recoveryErrorUserNotFound()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = recoverySuccessWithPhoneNumber()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = recoverySuccessWithEmail()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = recoverySuccessWithUsername()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	ID, Validation = getFirstGeneratedTicket()
	err = challRecoveryErrorBadPassword()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = challRecoveryWrongValidation()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = challRecoverySuccess()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = loginWrongCredentials()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = loginSuccess()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}
}
