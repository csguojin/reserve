package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/service"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

type UserRsp struct {
	ID       int    `json:"id" `
	Username string `json:"username" `
	Email    string `json:"email"`
}

func RegisterHandler(c *gin.Context) {
	var user *model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = service.CreateUser(user)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userRsp := &UserRsp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, userRsp)
}

type UserLoginRsp struct {
	Username string `json:"username" `
	Token    string `json:"token"`
}

func LoginHandler(c *gin.Context) {
	var user *model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = service.CheckUser(user.Username, user.Password)
	if err != nil {
		logger.Errorln(err)
		switch err {
		case util.ErrUserNotFound:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case util.ErrUserAuthFail:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	token, err := service.GenerateToken(user)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userRsp := &UserLoginRsp{
		Username: user.Username,
		Token:    token,
	}

	c.JSON(http.StatusOK, userRsp)
}
