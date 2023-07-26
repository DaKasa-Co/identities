package main

import (
	"github.com/DaKasa-Co/identities/controllers"
	"github.com/DaKasa-Co/identities/securities"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(securities.Authenticate())

	r.POST("/api/login", controllers.Login())
	r.POST("/api/register", controllers.Register())

	if err := r.Run(":9080"); err != nil {
		panic(err)
	}
}
