package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (h *HandlerStruct) GetAllRoomsHandler(c *gin.Context) {
	rooms, err := h.svc.GetAllRooms(parsePager(c))
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, rooms)
}

func (h *HandlerStruct) CreateRoomHandler(c *gin.Context) {
	var room *model.Room
	err := c.ShouldBindJSON(&room)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	room.ID = 0
	room, err = h.svc.CreateRoom(room)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, room)
}

func (h *HandlerStruct) GetRoomHandler(c *gin.Context) {
	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		logger.L.Errorln(util.ErrRoomIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrRoomIDNil.Error()})
		return
	}
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := h.svc.GetRoom(roomID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, room)
}

func (h *HandlerStruct) UpdateRoomHandler(c *gin.Context) {
	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		logger.L.Errorln(util.ErrRoomIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrRoomIDNil.Error()})
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
	room, err = h.svc.UpdateRoom(room)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, room)
}

func (h *HandlerStruct) DeleteRoomHandler(c *gin.Context) {
	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		logger.L.Errorln(util.ErrRoomIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrRoomIDNil.Error()})
		return
	}
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.svc.DeleteRoom(roomID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
