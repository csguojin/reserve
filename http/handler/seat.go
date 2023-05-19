package handler

import (
	"net/http"
	"strconv"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/service"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
	"github.com/gin-gonic/gin"
)

func GetAllSeatsHandler(c *gin.Context) {
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

	seats, err := service.GetAllSeats(roomID)
	if err != nil {
		logger.L.Errorln(err)
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

func CreateSeatHandler(c *gin.Context) {
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

	var seat *model.Seat
	err = c.ShouldBindJSON(&seat)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seat.ID = 0
	seat.RoomID = roomID
	seat, err = service.CreateSeat(seat)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seat)
}

func GetSeatHandler(c *gin.Context) {
	seatIDStr := c.Param("seat_id")
	if seatIDStr == "" {
		logger.L.Errorln(util.ErrSeatIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrSeatIDNil.Error()})
		return
	}
	seatID, err := strconv.Atoi(seatIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seat, err := service.GetSeat(seatID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seat)
}

func UpdateSeatHandler(c *gin.Context) {
	seatIDStr := c.Param("seat_id")
	if seatIDStr == "" {
		logger.L.Errorln(util.ErrSeatIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrSeatIDNil.Error()})
		return
	}
	seatID, err := strconv.Atoi(seatIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var seat *model.Seat
	err = c.ShouldBindJSON(&seat)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seat.ID = seatID
	seat, err = service.UpdateSeat(seat)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seat)
}

func DeleteSeatHandler(c *gin.Context) {
	seatIDStr := c.Param("seat_id")
	if seatIDStr == "" {
		logger.L.Errorln(util.ErrSeatIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrSeatIDNil.Error()})
		return
	}
	seatID, err := strconv.Atoi(seatIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteSeat(seatID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
