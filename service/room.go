package service

import (
	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func GetAllRooms() ([]*model.Room, error) {
	rooms, err := dal.GetAllRooms(dal.GetDB())
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return rooms, nil
}
