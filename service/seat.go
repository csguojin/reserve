package service

import (
	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func GetAllSeats(roomID int) ([]*model.Seat, error) {
	seats, err := dal.GetAllSeats(dal.GetDB(), roomID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seats, nil
}

func CreateSeat(seat *model.Seat) (*model.Seat, error) {
	seat, err := dal.CreateSeat(dal.GetDB(), seat)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func GetSeat(seatID int) (*model.Seat, error) {
	seat, err := dal.GetSeat(dal.GetDB(), seatID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func UpdateSeat(seat *model.Seat) (*model.Seat, error) {
	seat, err := dal.UpdateSeat(dal.GetDB(), seat)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func DeleteSeat(seatID int) error {
	err := dal.DeleteSeat(dal.GetDB(), seatID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
