package dal

import (
	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func GetAllRooms(db *gorm.DB) ([]*model.Room, error) {
	var rooms []*model.Room
	err := db.Find(&rooms).Error
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return rooms, nil
}
