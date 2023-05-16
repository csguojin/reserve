package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/csguojin/reserve/service"
	"github.com/csguojin/reserve/util/logger"
)

func GetAllRoomsHandler(c *gin.Context) {
	rooms, err := service.GetAllRooms()
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, rooms)
}
