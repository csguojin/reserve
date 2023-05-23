package service

import (
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func (s *svc) GetAllRooms(pager *model.Pager) ([]*model.Room, error) {
	rooms, err := s.dal.GetAllRooms(pager)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return rooms, nil
}

func (s *svc) CreateRoom(room *model.Room) (*model.Room, error) {
	room, err := s.dal.CreateRoom(room)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return room, nil
}

func (s *svc) GetRoom(roomID int) (*model.Room, error) {
	room, err := s.dal.GetRoom(roomID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return room, nil
}

func (s *svc) UpdateRoom(room *model.Room) (*model.Room, error) {
	room, err := s.dal.UpdateRoom(room)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return room, nil
}

func (s *svc) DeleteRoom(roomID int) error {
	err := s.dal.DeleteRoom(roomID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
