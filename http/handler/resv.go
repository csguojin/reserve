package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (h *HandlerStruct) GetResvsByUserHandler(c *gin.Context) {
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

	resvs, err := h.svc.GetResvsByUser(userID, parsePager(c))
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resvs)
}

func (h *HandlerStruct) CreateResvHandler(c *gin.Context) {
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

	resv, err = h.svc.CreateResv(resv)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}

func (h *HandlerStruct) SigninHandler(c *gin.Context) {
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

	resv, err := h.svc.Signin(resvID, userID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}

func (h *HandlerStruct) SignoutHandler(c *gin.Context) {
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

	resv, err := h.svc.Signout(resvID, userID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}

func (h *HandlerStruct) CancelResvHandler(c *gin.Context) {
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

	resv, err := h.svc.CancelResv(resvID, userID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resv)
}
