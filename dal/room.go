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

func CreateRoom(db *gorm.DB, room *model.Room) (*model.Room, error) {
	err := db.Create(room).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return GetRoom(db, room.ID)
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

func UpdateRoom(db *gorm.DB, room *model.Room) (*model.Room, error) {
	err := db.Model(&model.Room{}).Where("id = ?", room.ID).Updates(&room).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return GetRoom(db, room.ID)
}

func DeleteRoom(db *gorm.DB, roomID int) error {
	err := db.Delete(&model.Room{}, roomID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
