package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/service"
	"github.com/csguojin/reserve/util"
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

func CreateRoomHandler(c *gin.Context) {
	var room *model.Room
	err := c.ShouldBindJSON(&room)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	room, err = service.CreateRoom(room)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, room)
}

func GetRoomHandler(c *gin.Context) {
	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		logger.L.Errorln("room id is nil")
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrRoomNotFound})
		return
	}

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := service.GetRoom(roomID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, room)
}

func UpdateRoomHandler(c *gin.Context) {
	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		logger.L.Errorln("room id is nil")
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrRoomNotFound})
		return
	}

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var room *model.Room
	err = c.ShouldBindJSON(&room)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	room.ID = roomID
	room, err = service.UpdateRoom(room)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, room)

}

func DeleteRoomHandler(c *gin.Context) {
	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		logger.L.Errorln("room id is nil")
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrRoomNotFound})
		return
	}

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteRoom(roomID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
