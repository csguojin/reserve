package dal

import (
	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func GetAllRooms(db *gorm.DB) ([]*model.Room, error) {
	var rooms []*model.Room
	err := db.Find(&rooms).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return rooms, nil
}

func GetRoom(db *gorm.DB, id int) (*model.Room, error) {
	room := &model.Room{ID: id}
	err := db.First(&room, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrRoomNotFound
	}
	return room, nil
}
