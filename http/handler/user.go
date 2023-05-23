package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

type UserRsp struct {
	ID       int    `json:"id" `
	Username string `json:"username" `
	Email    string `json:"email"`
}

func (h *HandlerStruct) RegisterHandler(c *gin.Context) {
	var user *model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = h.svc.CreateUser(user)
	if err != nil {
		logger.L.Errorln(err)
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
	ID       int    `json:"id"`
	Username string `json:"username" `
	Token    string `json:"token"`
}

func (h *HandlerStruct) LoginHandler(c *gin.Context) {
	var user *model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = h.svc.CheckUser(user.Username, user.Password)
	if err != nil {
		logger.L.Errorln(err)
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

	token, err := h.svc.GenerateToken(user)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userRsp := &UserLoginRsp{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}

	c.JSON(http.StatusOK, userRsp)
}

func (h *HandlerStruct) GetAllUsersHandler(c *gin.Context) {
	users, err := h.svc.GetAllUsers(parsePager(c))
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var usersRsp []*UserRsp
	for _, v := range users {
		usersRsp = append(usersRsp, &UserRsp{
			ID:       v.ID,
			Username: v.Username,
			Email:    v.Email,
		})
	}

	c.JSON(http.StatusOK, usersRsp)
}

func (h *HandlerStruct) GetUserHandler(c *gin.Context) {
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

	user, err := h.svc.GetUser(userID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	userRsp := &UserRsp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, userRsp)
}

func (h *HandlerStruct) UpdateUserHandler(c *gin.Context) {
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

	var user *model.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = userID

	user, err = h.svc.UpdateUser(user)
	if err != nil {
		logger.L.Errorln(err)
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

func (h *HandlerStruct) DeleteUserHandler(c *gin.Context) {
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

	err = h.svc.DeleteUser(userID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
