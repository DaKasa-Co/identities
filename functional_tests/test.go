package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bitfield/script"
)

var unexpectedErrorInCreateRequest = "Occurs an unexpected error when create new request"

func main() {
	var body []byte

	fmt.Println("Send HTTP Request without API-KEY. Should return error")
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/v1/register", nil)
	if err != nil {
		log.Fatal(unexpectedErrorInCreateRequest)
	}

	_, err = script.Do(req).Stdout()
	fmt.Println()
	if !strings.Contains(err.Error(), "401") {
		log.Fatal("request without API-Key should return 401 status")
	}

	fmt.Println("Send HTTP Request without some request field. Should return error")
	body = []byte(`{"username": "teste", "email": "someemail@gmail.com"}`)
	req, err = http.NewRequest(http.MethodPost, "http://localhost:9080/v1/register", bytes.NewBuffer(body))
	req.Header.Add("X-Api-Key", "SomeApiKey")
	if err != nil {
		log.Fatal(unexpectedErrorInCreateRequest)
	}

	_, err = script.Do(req).Stdout()
	fmt.Println()
	if !strings.Contains(err.Error(), "400") {
		log.Fatal("request without JSON should return 400 status")
	}

	// TO DO: THE CHECK OF THIS TOPIC IS BROKEN. ALGORITHM NEED BE REALLY FUNCTIONAL
	fmt.Println("Send HTTP Request with all datas correct. Should return success")
	body = []byte(`{"username": "gio silva", "username": "gio._.", "email": "someemail@gmail.com", "password": "someGoodPassword", "birthday": "2003-05-11", "phoneNumber": 1291820931901823}`)
	req, err = http.NewRequest(http.MethodPost, "http://localhost:9080/v1/register", bytes.NewBuffer(body))
	req.Header.Add("X-Api-Key", "SomeApiKey")
	if err != nil {
		log.Fatal(unexpectedErrorInCreateRequest)
	}

	_, err = script.Do(req).Stdout()
	fmt.Println()
	if !strings.Contains(err.Error(), "400") {
		log.Fatal("request without JSON should return 400 status")
	}
}
