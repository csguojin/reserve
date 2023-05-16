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
		logger.L.Errorln(err)
		return nil, err
	}
	return seats, nil
}

func GetSeat(db *gorm.DB, id int) (*model.Seat, error) {
	seat := &model.Seat{ID: id}
	err := db.First(&seat, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrSeatNotFound
	}
	return seat, nil
}
