package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/csguojin/reserve/config"
	"github.com/csguojin/reserve/http/handler"
	"github.com/csguojin/reserve/util/logger"
)

func main() {
	defer logger.Logger.Sync()

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", handler.RegisterHandler)
		v1.POST("/login", handler.LoginHandler)

		v1.GET("/rooms", handler.GetAllRoomsHandler)
	}

	router.Run(":" + viper.GetString("server.port"))
}
