package main

import (
	"net/http"

	"github.com/DaKasa-Co/identities/client"
	"github.com/DaKasa-Co/identities/model"

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
	})
}
