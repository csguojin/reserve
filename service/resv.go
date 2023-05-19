package service

import (
	"errors"
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

func UpdateResvStatus(resv *model.Resv) (*model.Resv, error) {
	resv, err := dal.UpdateResvStatus(dal.GetDB(), resv)
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

	if resv.SigninTime != nil {
		err := errors.New("repeat signin error")
		logger.L.Errorln(err)
		return nil, err
	}

	if resv.Status != 0 {
		err := errors.New("reservation status error")
		logger.L.Errorln(err)
		return nil, err
	}

	now := time.Now()
	if now.Add(15 * time.Minute).Before(*resv.StartTime) {
		err := errors.New("sign in too early")
		logger.L.Errorln(err)
		return nil, err
	}

	if now.After((*resv.EndTime).Add(15 * time.Minute)) {
		err := errors.New("sign in too late")
		logger.L.Errorln(err)
		return nil, err
	}

	resv.SigninTime = &now

	resv, err = UpdateResvStatus(resv)
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

	if resv.SignoutTime != nil {
		err := errors.New("repeat signout error")
		logger.L.Errorln(err)
		return nil, err
	}

	if resv.Status != 0 {
		err := errors.New("reservation status error")
		logger.L.Errorln(err)
		return nil, err
	}

	now := time.Now()
	resv.SignoutTime = &now

	resv, err = UpdateResvStatus(resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	if now.After((*resv.EndTime)) {
		err := errors.New("the time is up, it ends automatically")
		logger.L.Errorln(err)
		return resv, err
	}

	return resv, nil
}

func CancelResv(resvID int, userID int) (*model.Resv, error) {
	resv, err := GetResv(resvID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	if resv.Status != 0 {
		err := errors.New("reservation status error")
		logger.L.Errorln(err)
		return nil, err
	}

	resv.Status = 1

	resv, err = UpdateResvStatus(resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}
