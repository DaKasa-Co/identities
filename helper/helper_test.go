package helper

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/DaKasa-Co/identities/model"
	database "github.com/DaKasa-Co/identities/psql"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCheckRequiredFieldExists(t *testing.T) {
	type RequiredFields struct {
		Name        string
		Key         string
		Value       interface{}
		IsSuccess   bool
		Description string
	}

	validate := []RequiredFields{
		{
			Name:        "RequiredFieldHasContent",
			Key:         "SomeField",
			Value:       "miau",
			IsSuccess:   true,
			Description: "The required field has content. The result must be true.",
		},
		{
			Name:        "ExpectedContentInRequiredField",
			Key:         "FakeField",
			Value:       nil,
			IsSuccess:   false,
			Description: "This field is required but has no content. The result must be false",
		},
	}

	for _, v := range validate {
		t.Run(v.Name, func(t *testing.T) {
			result := false
			if err := CheckRequiredFieldExists(v.Key, v.Value); err == nil {
				result = true
			}

			assert.Equal(t, v.IsSuccess, result, v.Description)
		})
	}
}

func TestCheckIsValidEmail(t *testing.T) {
	type Emails struct {
		Name      string
		Email     string
		IsSuccess bool
	}

	validate := []Emails{
		{
			Name:      "jorginho@gmail.com",
			Email:     "jorginho@gmail.com",
			IsSuccess: true,
		},
		{
			Name:      "emailinvalid.com",
			Email:     "emailinvalid.com",
			IsSuccess: false,
		},
		{
			Name:      "#@%^%#$@#$@#.com",
			Email:     "#@%^%#$@#$@#.com",
			IsSuccess: false,
		},
		{
			Name:      "first.last@sub.do,com",
			Email:     "first.last@sub.do,com",
			IsSuccess: false,
		},
		{
			Name:      "_______@example.com",
			Email:     "_______@example.com",
			IsSuccess: true,
		},
	}

	for _, v := range validate {
		t.Run(v.Name, func(t *testing.T) {
			result := false
			if err := CheckIsValidEmail(v.Email); err == nil {
				result = true
			}

			assert.Equal(t, v.IsSuccess, result)
		})
	}
}

func TestCheckIfNotHasSpecialCharacters(t *testing.T) {
	type Emails struct {
		Name      string
		Value     string
		IsSuccess bool
	}

	validate := []Emails{
		{
			Name:      "CamelCase",
			Value:     "noSpecialCharacters",
			IsSuccess: true,
		},
		{
			Name:      "NameWithSpace",
			Value:     "Mario de Andrade",
			IsSuccess: true,
		},
		{
			Name:      "ComplexNameWithSpace",
			Value:     "Gerthüdes Custódio Peña",
			IsSuccess: true,
		},
		{
			Name:      "InvalidEmail",
			Value:     "first.last@sub.do,com",
			IsSuccess: false,
		},
		{
			Name:      "SomeCostWithCurrencyBRL",
			Value:     "R$1000,00",
			IsSuccess: false,
		},
	}

	for _, v := range validate {
		t.Run(v.Name, func(t *testing.T) {
			result := false
			if err := CheckIfNotHasSpecialCharacters("SomeKey", v.Value); err == nil {
				result = true
			}

			assert.Equal(t, v.IsSuccess, result)
		})
	}
}

func TestCheckIsValidUsername(t *testing.T) {
	type Usernames struct {
		Name      string
		Usernames string
		IsSuccess bool
	}

	validate := []Usernames{
		{
			Name:      "lo.",
			Usernames: "lo.",
			IsSuccess: true,
		},
		{
			Name:      "real_oficial",
			Usernames: "real_oficial",
			IsSuccess: true,
		},
		{
			Name:      "d4rk-sh4d0w",
			Usernames: "d4rk-sh4d0w",
			IsSuccess: true,
		},
		{
			Name:      "lo",
			Usernames: "lo",
			IsSuccess: false,
		},
		{
			Name:      "neusa silva",
			Usernames: "neusa silva",
			IsSuccess: false,
		},
		{
			Name:      "paralelepipedo_mathematics_and_philosophy_1234567890",
			Usernames: "paralelepipedo_mathematics_and_philosophy_1234567890",
			IsSuccess: false,
		},
	}

	for _, v := range validate {
		t.Run(v.Name, func(t *testing.T) {
			result := false
			if err := CheckIsValidUsername(v.Usernames); err == nil {
				result = true
			}

			assert.Equal(t, v.IsSuccess, result)
		})
	}
}

func TestCheckIsValidPassword(t *testing.T) {
	type Passwords struct {
		Name      string
		Password  string
		IsSuccess bool
	}

	validate := []Passwords{
		{
			Name:      "SmallPassword",
			Password:  "potato",
			IsSuccess: false,
		},
		{
			Name:      "WithoutUpperCase",
			Password:  "potatoisgoodhahaha",
			IsSuccess: false,
		},
		{
			Name:      "WithoutLowerCase",
			Password:  "POTATEISGOOD",
			IsSuccess: false,
		},
		{
			Name:      "WithoutNumber",
			Password:  "PotateIsGood",
			IsSuccess: false,
		},
		{
			Name:      "WithoutSpecialChars",
			Password:  "PotateIsGood157",
			IsSuccess: false,
		},
		{
			Name:      "DecentPassword",
			Password:  "PotateIsGood157@*",
			IsSuccess: true,
		},
	}

	for _, v := range validate {
		t.Run(v.Name, func(t *testing.T) {
			result := false
			if err := CheckIsValidPassword(v.Password); err == nil {
				result = true
			}

			assert.Equal(t, v.IsSuccess, result)
		})
	}
}

func TestCheckBirthday(t *testing.T) {
	baseDate := time.Now()

	type Dates struct {
		Name        string
		Date        time.Time
		IsSuccess   bool
		Description string
	}

	validate := []Dates{
		{
			Name:        "TooOld",
			Date:        baseDate.AddDate(-215, 0, 0),
			IsSuccess:   false,
			Description: "User is too old to register. Is the user a mummy?",
		},
		{
			Name:        "TooKid",
			Date:        baseDate.AddDate(-10, 0, 0),
			IsSuccess:   false,
			Description: "User is too young to register. Kids and the internet are things that don't usually go together",
		},
		{
			Name:        "GoodUserAge",
			Date:        baseDate.AddDate(-18, 0, 0),
			IsSuccess:   true,
			Description: "User has the maturity to use our service",
		},
	}

	for _, v := range validate {
		t.Run(v.Name, func(t *testing.T) {
			result := false
			if err := CheckBirthday(v.Date); err == nil {
				result = true
			}

			assert.Equal(t, v.IsSuccess, result)
		})
	}
}

func TestErrorResponse(t *testing.T) {
	RespectiveError := errors.New("boom")
	RespectiveResult := gin.H{"msg": RespectiveError.Error()}

	result := ErrorResponse(RespectiveError)

	if diff := cmp.Diff(RespectiveResult, result); diff != "" {
		t.Errorf("\nFailed test in func \"ErrorResponse\":\n%s", diff)
	}
}

func TestPrepareUserRegisterDatas(t *testing.T) {
	birth, err := time.Parse(time.DateOnly, "2002-01-02")
	if err != nil {
		fmt.Println(err)
	}

	type RequestData struct {
		Name        string
		Req         model.Identity
		Err         error
		PrepUser    database.Users
		Description string
	}

	validate := []RequestData{
		{
			Name: "FailInRequiredField",
			Req: model.Identity{

				Username:    "teste",
				Email:       "teste@email.com",
				Password:    "somePass123@",
				Birthday:    "BOOM!",
				PhoneNumber: 12736182763,
			},
			Err:         fmt.Errorf("field name is required"),
			PrepUser:    database.Users{},
			Description: "Required fields is missing and should return error.",
		},
		{
			Name: "FailInParseBirthday",
			Req: model.Identity{
				Name:        "teste",
				Username:    "teste",
				Email:       "teste@email.com",
				Password:    "somePass123@",
				Birthday:    "BOOM!",
				PhoneNumber: 12736182763,
			},
			Err:         &time.ParseError{Layout: "2006-01-02", Value: "BOOM!", LayoutElem: "2006", ValueElem: "BOOM!"},
			PrepUser:    database.Users{},
			Description: "Bad birthday request and should return error.",
		},
		{
			Name: "FailInValidateName",
			Req: model.Identity{
				Name:        "h4ck3rm4n",
				Username:    "teste",
				Email:       "teste@email.com",
				Password:    "somePass123@",
				Birthday:    "2006-01-02",
				PhoneNumber: 12736182763,
			},
			Err:         fmt.Errorf("special characters not allowed in field name"),
			PrepUser:    database.Users{},
			Description: "Bad name request and should return error.",
		},
		{
			Name: "FailInValidateUsername",
			Req: model.Identity{
				Name:        "teste da silva",
				Username:    "ab",
				Email:       "teste@email.com",
				Password:    "somePass123@",
				Birthday:    "2006-01-02",
				PhoneNumber: 12736182763,
			},
			Err:         fmt.Errorf("respective username 'ab' is not valid"),
			PrepUser:    database.Users{},
			Description: "Bad username request and should return error.",
		},
		{
			Name: "FailInValidateEmail",
			Req: model.Identity{
				Name:        "teste da silva",
				Username:    "abc.123",
				Email:       "testeemail.com",
				Password:    "somePass123@",
				Birthday:    "2006-01-02",
				PhoneNumber: 12736182763,
			},
			Err:         fmt.Errorf("respective email address 'testeemail.com' is not valid"),
			PrepUser:    database.Users{},
			Description: "Bad email request and should return error.",
		},
		{
			Name: "FailInValidatePassword",
			Req: model.Identity{
				Name:        "teste da silva",
				Username:    "abc.123",
				Email:       "teste@email.com",
				Password:    "weakpw",
				Birthday:    "2006-01-02",
				PhoneNumber: 12736182763,
			},
			Err:         fmt.Errorf("password must be at least 8 characters long"),
			PrepUser:    database.Users{},
			Description: "Bad password request and should return error.",
		},
		{
			Name: "FailInValidateBirthday",
			Req: model.Identity{
				Name:        "teste da silva",
				Username:    "abc.123",
				Email:       "teste@email.com",
				Password:    "somePass123@",
				Birthday:    "1500-01-02",
				PhoneNumber: 12736182763,
			},
			Err:         fmt.Errorf("respective birthday 1500-01-02 is not valid"),
			PrepUser:    database.Users{},
			Description: "Bad birthday request and should return error.",
		},
		{
			Name: "FailInUploadImage",
			Req: model.Identity{
				Name:        "teste da silva",
				Username:    "abc.123",
				Email:       "teste@email.com",
				Password:    "somePass123@",
				Birthday:    "2002-01-02",
				Avatar:      "potato",
				PhoneNumber: 12736182763,
			},
			Err:         fmt.Errorf("must provide API Secret"),
			PrepUser:    database.Users{},
			Description: "Fails when try submit avatar image.",
		},
		{
			Name: "PrepareUserDataWithSuccess",
			Req: model.Identity{
				Name:        "teste da silva",
				Username:    "abc.123",
				Email:       "teste@email.com",
				Password:    "somePass123@",
				Birthday:    "2002-01-02",
				PhoneNumber: 12736182763,
			},
			Err: nil,
			PrepUser: database.Users{
				Name:        "teste da silva",
				Username:    "abc.123",
				Email:       "teste@email.com",
				Password:    "somePass123@",
				Birthday:    birth,
				PhoneNumber: 12736182763,
			},
			Description: "All sent datas is correct and should return success with prepared datas.",
		},
	}

	for _, v := range validate {
		t.Run(v.Name, func(t *testing.T) {
			s, err := PrepareUserRegisterDatas(v.Req)

			if diff := cmp.Diff(v.Err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s:\n%s", v.Description, diff)
			}

			if diff := cmp.Diff(v.PrepUser, s); diff != "" {
				t.Errorf("\n%s:\n%s", v.Description, diff)
			}
		})
	}
}

func TestGenerateJWT(t *testing.T) {
	t.Setenv("JWT_KEY", "someKey")

	key, _ := GenerateJWT(uuid.New(), "username", "", "", time.Now().AddDate(-17, 0, 0))
	assert.True(t, len(strings.Split(key, ".")) == 3)
}
