package main

import (
	"github.com/dgb35/telemogus_backend/db"
	"github.com/dgb35/telemogus_backend/handlers"
	"github.com/dgb35/telemogus_backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()

	r.POST("/signup", handlers.SignUp)
	r.POST("/login", handlers.Login)

	authorized := r.Group("/", utils.AuthRequired)
	{
		authorized.POST("/:masterId/ligths_off", handlers.LigthsOff)
		authorized.POST("/:masterId/ligths_on", handlers.LigthsOn)
		authorized.POST("/:masterId/static_on", handlers.StaticOn)
		authorized.POST("/:masterId/static_off", handlers.StaticOff)
	}

	r.Run() // Listen and serve on 0.0.0.0:8080
}
