package main

import (
	"github.com/DaKasa-Co/identities/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/login", controllers.Login())
	r.POST("/api/register", controllers.Register())
	r.POST("/api/recovery", controllers.OpenAccountRecovery())
	r.POST("/api/chall-recovery", controllers.UpdateByRecovery())

	if err := r.Run(":9080"); err != nil {
		panic(err)
	}
}
