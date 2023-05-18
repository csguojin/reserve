package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/csguojin/reserve/config"
	"github.com/csguojin/reserve/http/handler"
	"github.com/csguojin/reserve/http/middleware"
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

		v1.GET("/users/:user_id/reservations", middleware.AuthUserMiddleware(), handler.GetResvsByUserHandler)
		v1.POST("/users/:user_id/reservations", middleware.AuthUserMiddleware(), handler.CreateResvHandler)

		v1.POST("/users/:user_id/reservations/:resv_id/signin", middleware.AuthUserMiddleware(), handler.SigninHandler)
		v1.POST("/users/:user_id/reservations/:resv_id/signout", middleware.AuthUserMiddleware(), handler.SignoutHandler)
		v1.POST("/users/:user_id/reservations/:resv_id/cancel", middleware.AuthUserMiddleware(), handler.CancelResvHandler)
	}

	admin := v1.Group("admin")
	{
		admin.POST("/login", handler.AdminLoginHandler)
		admin.GET("/admins", middleware.AuthAdminMiddleware(), handler.GetAllAdminsHandler)
		admin.POST("/admins", middleware.AuthAdminMiddleware(), handler.CreateAdminHandler)
		admin.GET("/admins/:admin_id", middleware.AuthAdminMiddleware(), handler.GetAdminHandler)
		admin.DELETE("/admins/:admin_id", middleware.AuthAdminMiddleware(), handler.DeleteAdminHandler)

		admin.GET("/rooms", middleware.AuthAdminMiddleware(), handler.GetAllRoomsHandler)
		admin.POST("/rooms", middleware.AuthAdminMiddleware(), handler.CreateRoomHandler)
		admin.GET("/rooms/:room_id", middleware.AuthAdminMiddleware(), handler.GetRoomHandler)
		admin.PUT("/rooms/:room_id", middleware.AuthAdminMiddleware(), handler.UpdateRoomHandler)
		admin.DELETE("/rooms/:room_id", middleware.AuthAdminMiddleware(), handler.DeleteRoomHandler)

		admin.GET("/rooms/:room_id/seats", middleware.AuthAdminMiddleware(), handler.GetAllSeatsHandler)
		admin.POST("/rooms/:room_id/seats", middleware.AuthAdminMiddleware(), handler.CreateSeatHandler)
		admin.GET("/rooms/:room_id/seats/:seat_id", middleware.AuthAdminMiddleware(), handler.GetSeatHandler)
		admin.PUT("/rooms/:room_id/seats/:seat_id", middleware.AuthAdminMiddleware(), handler.UpdateSeatHandler)
		admin.DELETE("/rooms/:room_id/seats/:seat_id", middleware.AuthAdminMiddleware(), handler.DeleteSeatHandler)

		admin.GET("/users", middleware.AuthAdminMiddleware(), handler.GetAllUsersHandler)
		admin.GET("/users/:user_id", middleware.AuthAdminMiddleware(), handler.GetUserHandler)
		admin.DELETE("/users/:user_id", middleware.AuthAdminMiddleware(), handler.DeleteUserHandler)
	}

	router.Run(":" + viper.GetString("server.port"))
}
