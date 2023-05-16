package dal

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func CreateResv(db *gorm.DB, resv *model.Resv) (*model.Resv, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var seat model.Seat
	if err := tx.Where("id = ?", resv.SeatID).First(&seat).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var room model.Room
	if err := tx.Where("id = ?", seat.RoomID).First(&room).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var existingResvCount int64
	if err := tx.
		Model(&model.Resv{}).
		Where("seat_id = ?", resv.SeatID).
		Where("start_time < ? AND end_time > ?", resv.EndTime, resv.StartTime).
		Count(&existingResvCount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if existingResvCount > 0 {
		tx.Rollback()
		return nil, errors.New("reservation time conflict")
	}

	now := time.Now()
	resv.CreateTime = &now

	if err := tx.Create(resv).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return resv, nil
}

func GetResv(db *gorm.DB, resvID int) (*model.Resv, error) {
	resv := &model.Resv{ID: resvID}
	err := db.First(&resv, resvID).Error
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func UpdateResv(db *gorm.DB, resv *model.Resv) (*model.Resv, error) {
	err := db.Model(&model.Resv{}).Where("id = ?", resv.ID).Updates(
		&model.Resv{
			StartTime:   resv.StartTime,
			EndTime:     resv.EndTime,
			SigninTime:  resv.SigninTime,
			SignoutTime: resv.SignoutTime,
			Status:      resv.Status,
		}).Error
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return GetResv(db, resv.ID)
}

func GetResvsByUser(db *gorm.DB, userID int) ([]*model.Resv, error) {
	var resvs []*model.Resv
	err := db.Where(&model.Resv{UserID: userID}).Find(&resvs).Error
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resvs, nil
}

func GetResvsBySeat(db *gorm.DB, seatID int) ([]*model.Resv, error) {
	var resvs []*model.Resv
	err := db.Where(&model.Resv{SeatID: seatID}).Find(&resvs).Error
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resvs, nil
}
