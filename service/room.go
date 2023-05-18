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

func CreateRoom(room *model.Room) (*model.Room, error) {
	room, err := dal.CreateRoom(dal.GetDB(), room)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return room, nil
}

func GetRoom(roomID int) (*model.Room, error) {
	room, err := dal.GetRoom(dal.GetDB(), roomID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return room, nil
}

func UpdateRoom(room *model.Room) (*model.Room, error) {
	room, err := dal.UpdateRoom(dal.GetDB(), room)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return room, nil
}

func DeleteRoom(roomID int) error {
	err := dal.DeleteRoom(dal.GetDB(), roomID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
