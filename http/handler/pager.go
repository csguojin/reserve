package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func parsePager(c *gin.Context) *model.Pager {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		logger.L.Debugln(err)
		page = 1
	}
	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if err != nil {
		logger.L.Debugln(err)
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
