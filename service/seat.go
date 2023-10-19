package service

import (
	"context"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func (s *svc) GetAllSeats(ctx context.Context, roomID int, pager *model.Pager) ([]*model.Seat, error) {
	seats, err := s.dal.GetAllSeats(ctx, roomID, pager)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seats, nil
}

func (s *svc) CreateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	seat, err := s.dal.CreateSeat(ctx, seat)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func (s *svc) GetSeat(ctx context.Context, seatID int) (*model.Seat, error) {
	seat, err := s.dal.GetSeat(ctx, seatID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func (s *svc) UpdateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	seat, err := s.dal.UpdateSeat(ctx, seat)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func (s *svc) DeleteSeat(ctx context.Context, seatID int) error {
	err := s.dal.DeleteSeat(ctx, seatID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
