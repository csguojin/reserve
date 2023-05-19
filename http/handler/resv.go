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
		logger.L.Errorln(util.ErrUserIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrUserIDNil.Error()})
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resvs, err := service.GetResvsByUser(userID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resvs)
}

func CreateResvHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		logger.L.Errorln(util.ErrUserIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrUserIDNil.Error()})
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resv *model.Resv
	err = c.ShouldBindJSON(&resv)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if resv.SeatID <= 0 ||
		resv.StartTime == nil ||
		resv.EndTime == nil ||
		!resv.EndTime.After(*resv.StartTime) ||
		resv.EndTime.Sub(*resv.StartTime) >= time.Hour*24 {
		logger.L.Errorln(util.ErrResvTimeIllegal)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrResvTimeIllegal.Error()})
		return
	}

	resv.UserID = userID
	resv.Status = 0

	resv, err = service.CreateResv(resv)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}

func SigninHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		logger.L.Errorln(util.ErrUserIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrUserIDNil.Error()})
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resvIDStr := c.Param("resv_id")
	if resvIDStr == "" {
		logger.L.Errorln(util.ErrResvIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrResvIDNil.Error()})
		return
	}
	resvID, err := strconv.Atoi(resvIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resv, err := service.Signin(resvID, userID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}

func SignoutHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		logger.L.Errorln(util.ErrUserIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrUserIDNil.Error()})
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resvIDStr := c.Param("resv_id")
	if resvIDStr == "" {
		logger.L.Errorln(util.ErrResvIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrResvIDNil.Error()})
		return
	}
	resvID, err := strconv.Atoi(resvIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resv, err := service.Signout(resvID, userID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}

func CancelResvHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		logger.L.Errorln(util.ErrUserIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrUserIDNil.Error()})
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resvIDStr := c.Param("resv_id")
	if resvIDStr == "" {
		logger.L.Errorln(util.ErrResvIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrResvIDNil.Error()})
		return
	}
	resvID, err := strconv.Atoi(resvIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resv, err := service.CancelResv(resvID, userID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}
