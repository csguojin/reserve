package service

import (
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
