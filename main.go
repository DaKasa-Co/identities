package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DaKasa-Co/identities/client"
	"github.com/DaKasa-Co/identities/model"
	database "github.com/DaKasa-Co/identities/psql"
	"github.com/DaKasa-Co/identities/securities"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(securities.Authenticate())

	r.POST("/api/login", func(c *gin.Context) {
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
			"(email = $1 OR username = $2) AND password = crypt($3, password)"
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
	})

	r.POST("/api/register", func(c *gin.Context) {
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

		checkQuery := "SELECT id FROM users WHERE email=$1 OR username=$2"
		res, err := db.Query(checkQuery, req.Email, req.Username)
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
			err = errors.New("users with respective email ou username already exists")
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
	})

	if err := r.Run(":9080"); err != nil {
		panic(err)
	}
}
