package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/csguojin/reserve/config"
	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/http/handler"
	"github.com/csguojin/reserve/http/middleware"
	"github.com/csguojin/reserve/service"
	"github.com/csguojin/reserve/util/logger"
)

func main() {
	defer logger.L.Sync()

	db := dal.GetDB()
	dalClient := dal.GetDal(db)
	svc := service.NewService(dalClient)
	h := handler.NewHandler(svc)

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", h.RegisterHandler)
		v1.POST("/login", h.LoginHandler)

		v1.GET("/rooms", h.GetAllRoomsHandler)
		v1.GET("/rooms/:room_id/seats", h.GetAllSeatsHandler)

		v1.GET("/users/:user_id/reservations", middleware.AuthUserMiddleware(), h.GetAllResvsByUserHandler)
		v1.POST("/users/:user_id/reservations", middleware.AuthUserMiddleware(), h.CreateResvHandler)
		v1.GET("/users/:user_id/reservations/:resv_id", middleware.AuthUserMiddleware(), h.GetResvHandler)
		v1.POST("/users/:user_id/reservations/:resv_id/signin", middleware.AuthUserMiddleware(), h.SigninHandler)
		v1.POST("/users/:user_id/reservations/:resv_id/signout", middleware.AuthUserMiddleware(), h.SignoutHandler)
		v1.POST("/users/:user_id/reservations/:resv_id/cancel", middleware.AuthUserMiddleware(), h.CancelResvHandler)

		v1.PUT("/users/:user_id", middleware.AuthUserMiddleware(), h.UpdateUserHandler)
	}

	admin := v1.Group("admin")
	{
		admin.POST("/login", h.AdminLoginHandler)
		admin.GET("/admins", middleware.AuthAdminMiddleware(), h.GetAllAdminsHandler)
		admin.POST("/admins", middleware.AuthAdminMiddleware(), h.CreateAdminHandler)
		admin.GET("/admins/:admin_id", middleware.AuthAdminMiddleware(), h.GetAdminHandler)
		admin.DELETE("/admins/:admin_id", middleware.AuthAdminMiddleware(), h.DeleteAdminHandler)

		admin.GET("/rooms", middleware.AuthAdminMiddleware(), h.GetAllRoomsHandler)
		admin.POST("/rooms", middleware.AuthAdminMiddleware(), h.CreateRoomHandler)
		admin.GET("/rooms/:room_id", middleware.AuthAdminMiddleware(), h.GetRoomHandler)
		admin.PUT("/rooms/:room_id", middleware.AuthAdminMiddleware(), h.UpdateRoomHandler)
		admin.DELETE("/rooms/:room_id", middleware.AuthAdminMiddleware(), h.DeleteRoomHandler)

		admin.GET("/rooms/:room_id/seats", middleware.AuthAdminMiddleware(), h.GetAllSeatsHandler)
		admin.POST("/rooms/:room_id/seats", middleware.AuthAdminMiddleware(), h.CreateSeatHandler)
		admin.GET("/rooms/:room_id/seats/:seat_id", middleware.AuthAdminMiddleware(), h.GetSeatHandler)
		admin.PUT("/rooms/:room_id/seats/:seat_id", middleware.AuthAdminMiddleware(), h.UpdateSeatHandler)
		admin.DELETE("/rooms/:room_id/seats/:seat_id", middleware.AuthAdminMiddleware(), h.DeleteSeatHandler)

		admin.GET("/users", middleware.AuthAdminMiddleware(), h.GetAllUsersHandler)
		admin.GET("/users/:user_id", middleware.AuthAdminMiddleware(), h.GetUserHandler)
		admin.PUT("/users/:user_id", middleware.AuthAdminMiddleware(), h.UpdateUserHandler)
		admin.DELETE("/users/:user_id", middleware.AuthAdminMiddleware(), h.DeleteUserHandler)

		admin.GET("/users/:user_id/reservations", middleware.AuthAdminMiddleware(), h.GetAllResvsByUserHandler)
		admin.POST("/users/:user_id/reservations", middleware.AuthAdminMiddleware(), h.CreateResvHandler)
		admin.GET("/users/:user_id/reservations/:resv_id", middleware.AuthAdminMiddleware(), h.GetResvHandler)
		admin.POST("/users/:user_id/reservations/:resv_id/signin", middleware.AuthAdminMiddleware(), h.SigninHandler)
		admin.POST("/users/:user_id/reservations/:resv_id/signout", middleware.AuthAdminMiddleware(), h.SignoutHandler)
		admin.POST("/users/:user_id/reservations/:resv_id/cancel", middleware.AuthAdminMiddleware(), h.CancelResvHandler)
	}

	router.Run(":" + viper.GetString("server.port"))
}
