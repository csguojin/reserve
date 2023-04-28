package dal

import (
	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func GetAllSeats(db *gorm.DB, roomID int) ([]*model.Seat, error) {
	var seats []*model.Seat
	err := db.Where(&model.Seat{RoomID: roomID}).Find(&seats).Error
	if err != nil {
		logger.Errorln(err)
		return nil, util.ErrRoomNotFound
	}
	return seats, nil
}
