package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DaKasa-Co/identities/client"
	"github.com/DaKasa-Co/identities/external"
	"github.com/DaKasa-Co/identities/model"
	database "github.com/DaKasa-Co/identities/psql"
	"github.com/DaKasa-Co/identities/securities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Login is responsible to user authntication in service
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

// Register is responsible to sign up user in service
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

// OpenAccountRecovery is responsible to open ticket about recovery process
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

		id := uuid.New()
		validation := rand.Intn(999999-100000+1) + 100000
		query = "INSERT INTO recovery (id, user_id, validation) VALUES ($1, $2, $3)"
		_, err = db.Exec(query, id, rows[0].ID, validation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		if os.Getenv("TEST_IGNORE_EMAIL") != "true" {
			err = external.LoadedEmail.SendEmailToRecoverAccount([]string{rows[0].Email}, strconv.Itoa(validation))
			if err != nil {
				c.JSON(http.StatusServiceUnavailable, err.Error())
			}
		}

		c.JSON(http.StatusCreated, "{\"id\": \""+id.String()+"\"}")
	}
}

// UpdateByRecovery checks if who request ticket recovery is the account's owner. In positive case, updates to a new inserted password
func UpdateByRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := model.Identity{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, client.ErrorResponse(err))
			return
		}

		if err = client.CheckIsValidPassword(req.Password); err != nil {
			c.JSON(http.StatusBadRequest, client.ErrorResponse(err))
			return
		}

		db, err := database.OpenSQL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		query := "SELECT user_id, expire_at FROM recovery WHERE " +
			"id = $1 AND validation = $2;"
		res, err := db.Query(query, req.Status.Ticket, req.Status.Validation.Tmp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		defer res.Close()
		rows := []database.Recovery{}
		for res.Next() {
			row := new(database.Recovery)

			err = res.Scan(&row.UserID, &row.ExpireAt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
				return
			}

			rows = append(rows, *row)
		}

		if len(rows) != 1 {
			err = fmt.Errorf("failed in recovery validation")
			c.JSON(http.StatusForbidden, client.ErrorResponse(err))
			return
		}

		query = "DELETE FROM recovery WHERE id = $1;"
		_, err = db.Exec(query, req.Status.Ticket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		if rows[0].ExpireAt.Before(time.Now()) {
			err = fmt.Errorf("recovery request has been expired")
			c.JSON(http.StatusGone, client.ErrorResponse(err))
			return
		}

		query = "UPDATE users SET password = $1 WHERE id = $2;"
		_, err = db.Exec(query, req.Password, rows[0].UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
