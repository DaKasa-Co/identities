package main

import (
	"net/http"
	"time"

	"github.com/DaKasa-Co/identities/client"
	"github.com/DaKasa-Co/identities/model"
	database "github.com/DaKasa-Co/identities/psql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var req model.Identity

	r.POST("/api/"+model.Version+"/login", func(c *gin.Context) {
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, client.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, req)
	})

	r.POST("/api/"+model.Version+"/register", func(c *gin.Context) {
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, client.ErrorResponse(err))
			return
		}

		birthday, err := time.Parse("2006-01-02", req.Birthday)
		if err != nil {
			c.JSON(http.StatusBadRequest, client.ErrorResponse(err))
			return
		}

		db, err := database.OpenSQL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, client.ErrorResponse(err))
			return
		}

		s := database.Users{
			Name:        req.Name,
			Username:    req.Username,
			Email:       req.Email,
			Password:    req.Password,
			Birthday:    birthday,
			PhoneNumber: req.PhoneNumber,
			Address:     req.Address,
			Avatar:      req.Avatar,
		}

		query := "INSERT INTO users" +
			"(name, username, email, password, birthday, phonenumber, address, picture) " +
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
		_, err = db.Exec(query, s.Name, s.Username, s.Email, s.Password, s.Birthday, s.PhoneNumber, s.Address, s.Avatar)
		if err != nil {
			c.JSON(http.StatusBadGateway, client.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusCreated, nil)
	})

	r.Run(":9080")
}
