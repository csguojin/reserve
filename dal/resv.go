package dal

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func CreateResv(db *gorm.DB, resv *model.Resv) (*model.Resv, error) {
	tx := db.Begin()
	if tx.Error != nil {
		logger.L.Errorln(tx.Error)
		return nil, tx.Error
	}

	var seat model.Seat
	if err := tx.
		Where("id = ?", resv.SeatID).
		Where("status = ?", 0).
		First(&seat).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if seat.Status != 0 {
		tx.Rollback()
		err := errors.New("seat is not available")
		logger.L.Errorln(err)
		return nil, err
	}

	var room *model.Room
	if err := tx.
		Where("id = ?", seat.RoomID).
		Where("status = ?", 0).
		First(&room).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if ok := checkRoomResv(room, resv); !ok {
		tx.Rollback()
		err := errors.New("room is not available")
		logger.L.Errorln(err)
		return nil, err
	}

	var existingResvCount int64
	if err := tx.
		Model(&model.Resv{}).
		Where("seat_id = ?", resv.SeatID).
		Where("status = ?", 0).
		Where("start_time < ? AND end_time > ?", resv.EndTime, resv.StartTime).
		Count(&existingResvCount).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if existingResvCount > 0 {
		tx.Rollback()
		logger.L.Errorln(util.ErrResvSeatTimeConflict)
		return nil, util.ErrResvSeatTimeConflict
	}

	existingResvCount = 0
	if err := tx.
		Model(&model.Resv{}).
		Where("user_id = ?", resv.UserID).
		Where("status = ?", 0).
		Where("start_time < ? AND end_time > ?", resv.EndTime, resv.StartTime).
		Count(&existingResvCount).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if existingResvCount > 0 {
		tx.Rollback()
		logger.L.Errorln(util.ErrResvUserTimeConflict)
		return nil, util.ErrResvUserTimeConflict
	}

	now := time.Now()
	resv.CreateTime = &now
	resv.Status = 0

	if err := tx.Create(resv).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	err := tx.Commit().Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	return resv, nil
}

func checkRoomResv(room *model.Room, resv *model.Resv) bool {
	layout := "15:04:05"
	openTime, err := time.Parse(layout, room.OpenTime)
	if err != nil {
		logger.L.Errorln("room opentime format error", room.OpenTime)
		return false
	}
	closeTime, err := time.Parse(layout, room.CloseTime)
	if err != nil {
		logger.L.Errorln("room closetime format error", room.CloseTime)
		return false
	}
	return resv.StartTime.Hour() >= openTime.Hour() &&
		resv.StartTime.Minute() >= openTime.Minute() &&
		resv.StartTime.Second() >= openTime.Second() &&
		resv.EndTime.Hour() <= closeTime.Hour() &&
		resv.EndTime.Minute() <= closeTime.Minute() &&
		resv.EndTime.Second() <= closeTime.Second()
}

func GetResv(db *gorm.DB, resvID int) (*model.Resv, error) {
	resv := &model.Resv{ID: resvID}
	err := db.First(&resv, resvID).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func UpdateResvStatus(db *gorm.DB, resv *model.Resv) (*model.Resv, error) {
	err := db.Model(&model.Resv{}).Where("id = ?", resv.ID).Updates(
		&model.Resv{
			SigninTime:  resv.SigninTime,
			SignoutTime: resv.SignoutTime,
			Status:      resv.Status,
		}).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return GetResv(db, resv.ID)
}

func updateResvTime(db *gorm.DB, newResv *model.Resv) (*model.Resv, error) {
	tx := db.Begin()
	if tx.Error != nil {
		logger.L.Errorln(tx.Error)
		return nil, tx.Error
	}

	oldResv := &model.Resv{ID: newResv.ID}
	err := db.First(&oldResv, newResv.ID).Error
	if err != nil {
		logger.L.Errorln(err)
		tx.Rollback()
		return nil, err
	}

	if newResv.SeatID != oldResv.SeatID {
		err = errors.New("resv seat cannot update")
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if newResv.UserID != oldResv.UserID {
		err = errors.New("resv user cannot update")
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	var existingResvCount int64
	if err := tx.
		Model(&model.Resv{}).
		Where("seat_id = ?", newResv.SeatID).
		Where("id != ?", newResv.ID).
		Where("status = ?", 0).
		Where("start_time < ? AND end_time > ?", newResv.EndTime, newResv.StartTime).
		Count(&existingResvCount).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if existingResvCount > 0 {
		tx.Rollback()
		logger.L.Errorln(util.ErrResvSeatTimeConflict)
		return nil, util.ErrResvSeatTimeConflict
	}

	existingResvCount = 0
	if err := tx.
		Model(&model.Resv{}).
		Where("user_id = ?", newResv.UserID).
		Where("id != ?", newResv.ID).
		Where("status = ?", 0).
		Where("start_time < ? AND end_time > ?", newResv.EndTime, newResv.StartTime).
		Count(&existingResvCount).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if existingResvCount > 0 {
		tx.Rollback()
		logger.L.Errorln(util.ErrResvUserTimeConflict)
		return nil, util.ErrResvUserTimeConflict
	}

	err = db.Model(&model.Resv{}).Where("id = ?", newResv.ID).Updates(
		&model.Resv{
			StartTime: newResv.StartTime,
			EndTime:   newResv.EndTime,
		}).Error
	if err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	return GetResv(db, newResv.ID)
}

func GetResvsByUser(db *gorm.DB, userID int) ([]*model.Resv, error) {
	var resvs []*model.Resv
	err := db.Where(&model.Resv{UserID: userID}).Find(&resvs).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resvs, nil
}

func GetResvsBySeat(db *gorm.DB, seatID int) ([]*model.Resv, error) {
	var resvs []*model.Resv
	err := db.Where(&model.Resv{SeatID: seatID}).Find(&resvs).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resvs, nil
}
