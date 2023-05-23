package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

type AdminRsp struct {
	ID    int    `json:"id" `
	Name  string `json:"name" `
	Email string `json:"email"`
}

func (h *HandlerStruct) CreateAdminHandler(c *gin.Context) {
	var admin *model.Admin
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	admin, err = h.svc.CreateAdmin(admin)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	adminRsp := &AdminRsp{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
	}

	c.JSON(http.StatusOK, adminRsp)
}

type AdminLoginRsp struct {
	ID    int    `json:"id"`
	Name  string `json:"name" `
	Token string `json:"token"`
}

func (h *HandlerStruct) AdminLoginHandler(c *gin.Context) {
	var admin *model.Admin
	err := c.ShouldBindJSON(&admin)
	if err != nil || admin.Name == "" || admin.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err = h.svc.CheckAdmin(admin.Name, admin.Password)
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

	token, err := h.svc.GenerateAdminToken(admin)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	adminRsp := &AdminLoginRsp{
		ID:    admin.ID,
		Name:  admin.Name,
		Token: token,
	}

	c.JSON(http.StatusOK, adminRsp)
}

func (h *HandlerStruct) GetAllAdminsHandler(c *gin.Context) {
	admins, err := h.svc.GetAllAdmins(parsePager(c))
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var adminsRsp []*AdminRsp
	for _, v := range admins {
		adminsRsp = append(adminsRsp, &AdminRsp{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		})
	}

	c.JSON(http.StatusOK, adminsRsp)
}

func (h *HandlerStruct) GetAdminHandler(c *gin.Context) {
	adminIDStr := c.Param("admin_id")
	if adminIDStr == "" {
		logger.L.Errorln(util.ErrAdminIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrAdminIDNil.Error()})
		return
	}
	adminID, err := strconv.Atoi(adminIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.svc.GetAdminNoPassword(adminID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	adminRsp := &AdminRsp{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
	}

	c.JSON(http.StatusOK, adminRsp)
}

func (h *HandlerStruct) DeleteAdminHandler(c *gin.Context) {
	adminIDStr := c.Param("admin_id")
	if adminIDStr == "" {
		logger.L.Errorln(util.ErrAdminIDNil)
		c.JSON(http.StatusBadRequest, gin.H{"error": util.ErrAdminIDNil.Error()})
		return
	}
	adminID, err := strconv.Atoi(adminIDStr)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.svc.DeleteAdmin(adminID)
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
