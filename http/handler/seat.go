package handler

import (
	"net/http"
	"strconv"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
	"github.com/gin-gonic/gin"
)

func (h *HandlerStruct) GetAllSeatsHandler(c *gin.Context) {
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

	seats, err := h.svc.GetAllSeats(roomID, parsePager(c))
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

func (h *HandlerStruct) CreateSeatHandler(c *gin.Context) {
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
	seat, err = h.svc.CreateSeat(seat)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seat)
}

func (h *HandlerStruct) GetSeatHandler(c *gin.Context) {
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

	seat, err := h.svc.GetSeat(seatID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seat)
}

func (h *HandlerStruct) UpdateSeatHandler(c *gin.Context) {
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
	seat, err = h.svc.UpdateSeat(seat)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seat)
}

func (h *HandlerStruct) DeleteSeatHandler(c *gin.Context) {
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

	err = h.svc.DeleteSeat(seatID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
