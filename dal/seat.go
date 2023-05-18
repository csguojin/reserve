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

func CreateSeat(db *gorm.DB, seat *model.Seat) (*model.Seat, error) {
	err := db.Create(seat).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return GetSeat(db, seat.ID)
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

func UpdateSeat(db *gorm.DB, seat *model.Seat) (*model.Seat, error) {
	err := db.Model(&model.Seat{}).Where("id = ?", seat.ID).Updates(&seat).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return GetSeat(db, seat.ID)
}

func DeleteSeat(db *gorm.DB, seatID int) error {
	err := db.Delete(&model.Seat{}, seatID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
