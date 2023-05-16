package dal

import (
	"time"

	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func CreateResv(db *gorm.DB, resv *model.Resv) (*model.Resv, error) {
	now := time.Now()
	resv.CreateTime = &now
	err := db.Create(resv).Error
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return GetResv(db, resv.ID)
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
