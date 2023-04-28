package service

import (
	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func GetAllSeats(roomID int) ([]*model.Seat, error) {
	seats, err := dal.GetAllSeats(dal.GetDB(), roomID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return seats, nil
}
