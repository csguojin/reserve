package service

import (
	"context"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func (s *svc) GetAllRooms(ctx context.Context, pager *model.Pager) ([]*model.Room, error) {
	rooms, err := s.dal.GetAllRooms(ctx, pager)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return rooms, nil
}

func (s *svc) CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	room, err := s.dal.CreateRoom(ctx, room)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return room, nil
}

func (s *svc) GetRoom(ctx context.Context, roomID int) (*model.Room, error) {
	room, err := s.dal.GetRoom(ctx, roomID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return room, nil
}

func (s *svc) UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	room, err := s.dal.UpdateRoom(ctx, room)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return room, nil
}

func (s *svc) DeleteRoom(ctx context.Context, roomID int) error {
	err := s.dal.DeleteRoom(ctx, roomID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
