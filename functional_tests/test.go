package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func registerMissingApiKey() error {
	fmt.Println("Send HTTP Request without API-KEY. Should return error")
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/v1/register", nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 401 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New("request without API-Key should return 401 status")
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func registerMissingRequiredFields() error {
	fmt.Println("Send HTTP Request without some request field. Should return error")
	body := []byte(`{"username": "teste", "email": "someemail@gmail.com"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/v1/register", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 400 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New("request without JSON should return 400 status")
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func registerSuccess() error {
	fmt.Println("Send HTTP Request with all datas correct. Should return success")
	body := []byte(`{"name": "gio silva", "username": "gio._.", "email": "someemail@gmail.com", "password": "someGoodPassword", "birthday": "2003-05-11", "phoneNumber": 1291820931901823}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/v1/register", bytes.NewBuffer(body))
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
		return errors.New("request should return success with 201 status")
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func registerConflictWithUsernameFieldThatAlreadyExists() error {
	fmt.Println("Send HTTP Request with username and email that already exists in database. Should return error")
	body := []byte(`{"name": "gio silva", "username": "gio._.", "email": "someemasil@gmail.com", "password": "someGoodPassword", "birthday": "2003-05-11", "phoneNumber": 1291820931901823}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/v1/register", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 409 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New("request with username or email field that already exists should return 409 status")
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func loginMissingApiKey() error {
	fmt.Println("Send HTTP Request without API-KEY. Should return error")
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/v1/login", nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 401 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New("request without API-Key should return 401 status")
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func loginWrongCredentials() error {
	fmt.Println("Send HTTP Request with bad credentials. Should return error")
	body := []byte(`{"email": "someemail@gmail.com", "password": "someBadPassword"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/v1/login", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 403 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New("request with incorrect credentials should return 403 status")
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func loginSuccess() error {
	fmt.Println("Send HTTP Request with correct credentials. Should return success and JWT token in the header")
	body := []byte(`{"email": "someemail@gmail.com", "password": "someGoodPassword"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/v1/login", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		fmt.Printf("Status: %d\n", res.StatusCode)
		return errors.New("request with correct credentials should return 200 status")
	}

	if res.Header.Get("X-JWT") == "" {
		fmt.Printf("Headers: %v\n", res.Header)
		return errors.New("JWT header is required")
	}

	fmt.Print("✅ Success\n\n")
	return nil
}

func main() {
	err := registerMissingApiKey()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = registerMissingRequiredFields()
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

	err = loginMissingApiKey()
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
