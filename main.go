package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/csguojin/reserve/config"
	"github.com/csguojin/reserve/http/handler"
	"github.com/csguojin/reserve/util/logger"
)

func main() {
	defer logger.L.Sync()

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", handler.RegisterHandler)
		v1.POST("/login", handler.LoginHandler)

		v1.GET("/rooms", handler.GetAllRoomsHandler)
		v1.GET("/rooms/:room_id/seats", handler.GetAllSeatsHandler)

		v1.GET("/users/:user_id/reservations", handler.GetResvsByUserHandler)
		v1.POST("/users/:user_id/reservations", handler.CreateResvHandler)

		v1.POST("/users/:user_id/reservations/:resv_id/signin", handler.SigninHandler)
		v1.POST("/users/:user_id/reservations/:resv_id/signout", handler.SignoutHandler)
		v1.POST("/users/:user_id/reservations/:resv_id/cancel", handler.CancelResvHandler)
	}

	admin := v1.Group("admin")
	{
		admin.GET("/rooms", handler.GetAllRoomsHandler)
		admin.POST("/rooms", handler.CreateRoomHandler)
		admin.GET("/rooms/:room_id", handler.GetRoomHandler)
		admin.PUT("/rooms/:room_id", handler.UpdateRoomHandler)
		admin.DELETE("/rooms/:room_id", handler.DeleteRoomHandler)
	}

	router.Run(":" + viper.GetString("server.port"))
}
