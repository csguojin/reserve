package handler

import (
	"net/http"
	"strconv"

	"github.com/csguojin/reserve/service"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
	"github.com/gin-gonic/gin"
)

func GetAllSeatsHandler(c *gin.Context) {
	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		logger.Errorln("room id is nil")
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrRoomNotFound})
		return
	}

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seats, err := service.GetAllSeats(roomID)
	if err != nil {
		logger.Errorln(err)
		switch err {
		case util.ErrRoomNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, nil)
		}
		return
	}

	c.JSON(http.StatusOK, seats)
}
