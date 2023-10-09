package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// @@ REGISTER @@
func TestRegister() {
	fmt.Println("🧪 REGISTER 🧪")
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
	fmt.Print("✅ Success\n\n")
}

func registerMissingRequiredFields() error {
	fmt.Println("💡 Request without some request field. Should return error")
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 400)
	}

	desiredResponse := `{"msg":"field <field> is required"}`
	if !strings.Contains(string(body), "required") {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

func registerSuccess() error {
	fmt.Println("💡 Request with all datas correct. Should return success")
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 201)
	}

	desiredResponse := `{"msg":"register user with success"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

func registerConflictWithUsernameFieldThatAlreadyExists() error {
	fmt.Println("💡 Username that already exists in database. Should return error")
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 409)
	}

	desiredResponse := `{"msg":"users with respective email, username or phone number already exists"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

func registerConflictWithEmailFieldThatAlreadyExists() error {
	fmt.Println("💡 Email that already exists in database. Should return error")
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 409)
	}

	desiredResponse := `{"msg":"users with respective email, username or phone number already exists"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

func registerConflictWithPhoneNumberFieldThatAlreadyExists() error {
	fmt.Println("💡 PhoneNumber that already exists in database. Should return error. Should return error")
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 409)
	}

	desiredResponse := `{"msg":"users with respective email, username or phone number already exists"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

// @@ LOGIN @@
func TestLogin() {
	fmt.Println("🧪 LOGIN 🧪")
	err := loginIncorrectCredentials()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = loginWithSuccess()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}
	fmt.Print("✅ Success\n\n")
}

func loginIncorrectCredentials() error {
	fmt.Println("💡 Try login but incorrect credentials. Should return error")
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
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 403 {
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 403)
	}

	desiredResponse := `{"msg":"incorrect credentials"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

func loginWithSuccess() error {
	fmt.Println("💡 Try login with correct credentials. Should return success")
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
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 200)
	}

	if res.Header.Get("X-JWT") == "" {
		fmt.Printf("Headers: %v\n", res.Header)
		return errors.New("JWT header is required")
	}

	desiredResponse := `{"msg":"user has logged in with success"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

// @@ VALIDATE RECOVERY TICKET @@
func TestValidateRecoveryTicket() {
	fmt.Println("🧪 VALIDATE TICKET RECOVERY 🧪")
	err := validateRecoveryWrongValidation()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}

	err = validateRecoverySuccess()
	if err != nil {
		fmt.Println("❌ Fail")
		log.Fatal(err)
	}
	fmt.Print("✅ Success\n\n")
}

func validateRecoveryWrongValidation() error {
	fmt.Println("💡 Try validate ticket with incorrect validation number. Should return error")
	ID, _ := getFirstGeneratedTicket()

	body := []byte(`{"password": "BadPass123#", "status": {"ticket": "` + ID.String() + `", "validation": {"tmp": 222}}}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery/validate", bytes.NewBuffer(body))
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 403)
	}

	desiredResponse := `{"msg":"failed in validate ticket recovery"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

func validateRecoverySuccess() error {
	fmt.Println("💡 Validate recovery ticket with success")
	ID, Validation := getFirstGeneratedTicket()

	body := []byte(`{"password": "GoodPass123#", "status": {"ticket": "` + ID.String() + `", "validation": {"tmp": ` + strconv.Itoa(Validation) + `}}}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery/validate", bytes.NewBuffer(body))
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 200)
	}

	desiredResponse := `{"msg":"validate ticket with success"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}
	return nil
}

// @@ CHALL @@
func TestChallengeRecoveryTicket() {
	fmt.Println("🧪 CHALL RECOVERY 🧪")
	err := challRecoveryErrorBadPassword()
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
	fmt.Print("✅ Success\n\n")
}

func challRecoveryErrorBadPassword() error {
	fmt.Println("💡 New bad password. Should return error")
	body := []byte(`{"password": "BadPass123"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery/chall", bytes.NewBuffer(body))
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 400)
	}

	desiredResponse := `{"msg":"password must have at least one special character"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

func challRecoveryWrongValidation() error {
	fmt.Println("💡 Validation number is incorrect. Should return error")
	ID, _ := getFirstGeneratedTicket()

	body := []byte(`{"password": "BadPass123#", "status": {"ticket": "` + ID.String() + `", "validation": {"tmp": 222}}}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery/chall", bytes.NewBuffer(body))
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 403)
	}

	desiredResponse := `{"msg":"failed in recovery validation"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

func challRecoverySuccess() error {
	fmt.Println("💡 Recovery challenge with positive result. Should return success")
	ID, Validation := getFirstGeneratedTicket()

	body := []byte(`{"password": "GoodPass123#", "status": {"ticket": "` + ID.String() + `", "validation": {"tmp": ` + strconv.Itoa(Validation) + `}}}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery/chall", bytes.NewBuffer(body))
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 200)
	}

	desiredResponse := `{"msg":"recovery account with success"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}
	return nil
}

// @@ INIT RECOVERY @@
func TestCreateRecoveryTicket() {
	fmt.Println("🧪 INIT RECOVERY 🧪")
	err := recoveryErrorUserNotFound()
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
	fmt.Print("✅ Success\n\n")
}

func recoveryErrorUserNotFound() error {
	fmt.Println("💡 Try recovery account with phone number that doens't exist in user database. Should return error")
	body := []byte(`{"phoneNumber": 12345}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery/create", bytes.NewBuffer(body))
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 404)
	}

	desiredResponse := `{"msg":"user not found"}`
	if string(body) != desiredResponse {
		return fmt.Errorf("\n\n-got: %s\n+want: %s", string(body), desiredResponse)
	}

	return nil
}

func recoverySuccessWithPhoneNumber() error {
	fmt.Println("💡 Create ticket recovery with phone number. Should return success.")
	body := []byte(`{"phoneNumber": 1291820931901823}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery/create", bytes.NewBuffer(body))
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 201)
	}

	return nil
}

func recoverySuccessWithEmail() error {
	fmt.Println("💡 Create ticket recovery with email. Should return success.")
	body := []byte(`{"email": "someemail@gmail.com"}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery/create", bytes.NewBuffer(body))
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
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 201)
	}

	return nil
}

func recoverySuccessWithUsername() error {
	fmt.Println("💡 Create ticket recovery with username. Should return success.")
	body := []byte(`{"username": "gio._."}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9080/api/recovery/create", bytes.NewBuffer(body))
	req.Header.Add("X-API-Key", "SomeApiKey")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		fmt.Println(string(body))
		return fmt.Errorf("\n\n-got: %d\n+want: %d", res.StatusCode, 201)
	}
	return nil
}
