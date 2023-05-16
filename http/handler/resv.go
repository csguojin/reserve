package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/service"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func GetResvsByUserHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		logger.Errorln(util.ErrUserIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrUserIDNil.Error()})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resvs, err := service.GetResvsByUser(userID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resvs)
}

func CreateResvHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		logger.Errorln(util.ErrUserIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrUserIDNil.Error()})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resv *model.Resv
	err = c.ShouldBindJSON(&resv)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if resv.SeatID <= 0 ||
		resv.StartTime == nil ||
		resv.EndTime == nil ||
		!resv.EndTime.After(*resv.StartTime) ||
		resv.EndTime.Sub(*resv.StartTime) >= time.Hour*24 {
		logger.Errorln(util.ErrRequestBodyFormat)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrRequestBodyFormat.Error()})
		return
	}

	resv.UserID = userID
	resv.Status = 0

	resv, err = service.CreateResv(resv)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}

func UpdateResvHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		logger.Errorln(util.ErrUserIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrUserIDNil.Error()})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resvIDStr := c.Param("resv_id")
	if resvIDStr == "" {
		logger.Errorln(util.ErrUserIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrUserIDNil.Error()})
		return
	}

	resvID, err := strconv.Atoi(resvIDStr)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resv *model.Resv
	err = c.ShouldBindJSON(&resv)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resv.ID = resvID
	resv.UserID = userID

	resv, err = service.UpdateResv(resv)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}
