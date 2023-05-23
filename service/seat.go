package service

import (
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func (s *svc) GetAllSeats(roomID int, pager *model.Pager) ([]*model.Seat, error) {
	seats, err := s.dal.GetAllSeats(roomID, pager)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seats, nil
}

func (s *svc) CreateSeat(seat *model.Seat) (*model.Seat, error) {
	seat, err := s.dal.CreateSeat(seat)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func (s *svc) GetSeat(seatID int) (*model.Seat, error) {
	seat, err := s.dal.GetSeat(seatID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func (s *svc) UpdateSeat(seat *model.Seat) (*model.Seat, error) {
	seat, err := s.dal.UpdateSeat(seat)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func (s *svc) DeleteSeat(seatID int) error {
	err := s.dal.DeleteSeat(seatID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
