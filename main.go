package main

import (
	"github.com/DaKasa-Co/identities/controllers"
	"github.com/DaKasa-Co/identities/securities"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(securities.CORSMiddleware())

	r.POST("/api/login", controllers.Login())
	r.POST("/api/register", controllers.Register())
	r.POST("/api/recovery/create", controllers.OpenAccountRecovery())
	r.POST("/api/recovery/chall", controllers.UpdateByRecovery())
	r.POST("/api/recovery/validate", controllers.CheckTicketRecovery())

	if err := r.Run(":9080"); err != nil {
		panic(err)
	}
}
