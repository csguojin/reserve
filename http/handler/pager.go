package handler

import (
	"net/http"
	"strconv"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
	"github.com/gin-gonic/gin"
)

func parsePager(c *gin.Context) *model.Pager {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		page = 1
	}
	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if err != nil {
		logger.L.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		perPage = 10
	}
	if page <= 0 {
		page = 1
	}
	if perPage > 100 {
		perPage = 100
	}
	if perPage <= 0 {
		perPage = 1
	}

	return &model.Pager{
		Page:    page,
		PerPage: perPage,
	}
}
