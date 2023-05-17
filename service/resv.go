package service

import (
	"time"

	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util/logger"
)

func CreateResv(resv *model.Resv) (*model.Resv, error) {
	resv, err := dal.CreateResv(dal.GetDB(), resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func GetResv(resvID int) (*model.Resv, error) {
	resv, err := dal.GetResv(dal.GetDB(), resvID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func UpdateResv(resv *model.Resv) (*model.Resv, error) {
	resv, err := dal.UpdateResv(dal.GetDB(), resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func GetResvsByUser(userID int) ([]*model.Resv, error) {
	resv, err := dal.GetResvsByUser(dal.GetDB(), userID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func GetResvsBySeat(seatID int) ([]*model.Resv, error) {
	resv, err := dal.GetResvsBySeat(dal.GetDB(), seatID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func Signin(resvID int, userID int) (*model.Resv, error) {
	resv, err := GetResv(resvID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	now := time.Now()
	resv.SigninTime = &now

	resv, err = UpdateResv(resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func Signout(resvID int, userID int) (*model.Resv, error) {
	resv, err := GetResv(resvID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	now := time.Now()
	resv.SignoutTime = &now
	resv.Status = 1

	resv, err = UpdateResv(resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func CancelResv(resvID int, userID int) (*model.Resv, error) {
	resv, err := GetResv(resvID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	resv.Status = 2

	resv, err = UpdateResv(resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}
