package handler

import (
	"net/http"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/service"
	"github.com/gin-gonic/gin"
)

type UserRsp struct {
	ID       int64  `json:"id" `
	Username string `json:"username" `
	Email    string `json:"email"`
}

func RegisterHandler(c *gin.Context) {
	var user *model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = service.CreateUser(user)
	if err != nil {
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
