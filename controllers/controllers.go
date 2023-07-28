package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/DaKasa-Co/identities/client"
	"github.com/DaKasa-Co/identities/external"
	"github.com/DaKasa-Co/identities/model"
	database "github.com/DaKasa-Co/identities/psql"
	"github.com/DaKasa-Co/identities/securities"
	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := model.Identity{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, client.ErrorResponse(err))
			return
		}

		if req.Email != "" {
			if err = client.CheckIsValidEmail(req.Email); err != nil {
				c.JSON(http.StatusForbidden, client.ErrorResponse(err))
				return
			}
		} else {
			if err = client.CheckIsValidUsername(req.Username); err != nil {
				c.JSON(http.StatusForbidden, client.ErrorResponse(err))
				return
			}
		}

		if err = client.CheckIsValidPassword(req.Password); err != nil {
			c.JSON(http.StatusForbidden, client.ErrorResponse(err))
			return
		}

		db, err := database.OpenSQL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		query := "SELECT id, username FROM users WHERE " +
			"(email = $1 OR username = $2) AND password = crypt($3, password);"
		res, err := db.Query(query, req.Email, req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		defer res.Close()
		rows := []database.Users{}
		for res.Next() {
			row := new(database.Users)

			err = res.Scan(&row.ID, &row.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
				return
			}

			rows = append(rows, *row)
		}

		if len(rows) != 1 {
			err = fmt.Errorf("incorrect credentials")
			c.JSON(http.StatusForbidden, client.ErrorResponse(err))
			return
		}

		jwt, err := securities.GenerateJWT(rows[0].ID, rows[0].Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		c.Header("X-JWT", jwt)
		c.JSON(http.StatusOK, nil)
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := model.Identity{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, client.ErrorResponse(err))
			return
		}

		s, err := client.PrepareUserRegisterDatas(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, client.ErrorResponse(err))
			return
		}

		db, err := database.OpenSQL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		checkQuery := "SELECT id FROM users WHERE email=$1 OR username=$2 OR phonenumber=$3"
		res, err := db.Query(checkQuery, s.Email, s.Username, s.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		defer res.Close()
		rows := []database.Users{}
		for res.Next() {
			row := new(database.Users)

			err = res.Scan(&row.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
				return
			}

			rows = append(rows, *row)
		}

		if len(rows) > 0 {
			err = errors.New("users with respective email, username or phone number already exists")
			c.JSON(http.StatusConflict, client.ErrorResponse(err))
			return
		}

		query := "INSERT INTO users" +
			"(name, username, email, password, birthday, phonenumber, address, picture) " +
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
		_, err = db.Exec(query, s.Name, s.Username, s.Email, s.Password, s.Birthday, s.PhoneNumber, s.Address, s.Avatar)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusCreated, nil)
	}
}

func OpenAccountRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := model.Identity{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, client.ErrorResponse(err))
			return
		}

		db, err := database.OpenSQL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		query := "SELECT id, email FROM users WHERE " +
			"email = $1 OR username = $2 OR phonenumber = $3;"
		res, err := db.Query(query, req.Email, req.Username, req.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		defer res.Close()
		rows := []database.Users{}
		for res.Next() {
			row := new(database.Users)

			err = res.Scan(&row.ID, &row.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
				return
			}

			rows = append(rows, *row)
		}

		if len(rows) != 1 {
			err = fmt.Errorf("user not found")
			c.JSON(http.StatusNotFound, client.ErrorResponse(err))
			return
		}

		email := rows[0].Email
		fmt.Println(rows[0])
		fmt.Println(email)
		validation := rand.Intn(999999-100000+1) + 100000
		query = "INSERT INTO recovery (id, validation) VALUES ($1, $2)"
		_, err = db.Exec(query, rows[0].ID, validation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		fmt.Println(email)
		err = external.LoadedEmail.SendEmailToRecoverAccount([]string{email}, strconv.Itoa(validation))
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, err.Error())
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
