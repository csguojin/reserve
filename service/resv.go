package service

import (
	"errors"
	"time"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (s *svc) CreateResv(resv *model.Resv) (*model.Resv, error) {
	resv, err := s.dal.CreateResv(resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) GetResv(resvID int) (*model.Resv, error) {
	resv, err := s.dal.GetResv(resvID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) UpdateResvStatus(resv *model.Resv) (*model.Resv, error) {
	resv, err := s.dal.UpdateResvStatus(resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) GetResvsByUser(userID int) ([]*model.Resv, error) {
	resv, err := s.dal.GetResvsByUser(userID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) GetResvsBySeat(seatID int) ([]*model.Resv, error) {
	resv, err := s.dal.GetResvsBySeat(seatID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) Signin(resvID int, userID int) (*model.Resv, error) {
	resv, err := s.dal.GetResv(resvID)
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
		logger.L.Errorln(util.ErrResvCanceled)
		return nil, util.ErrResvCanceled
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

	resv, err = s.dal.UpdateResvStatus(resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) Signout(resvID int, userID int) (*model.Resv, error) {
	resv, err := s.dal.GetResv(resvID)
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
		logger.L.Errorln(util.ErrResvCanceled)
		return nil, util.ErrResvCanceled
	}

	now := time.Now()
	resv.SignoutTime = &now

	resv, err = s.dal.UpdateResvStatus(resv)
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

func (s *svc) CancelResv(resvID int, userID int) (*model.Resv, error) {
	resv, err := s.dal.GetResv(resvID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	if resv.Status != 0 {
		logger.L.Errorln(util.ErrResvCanceled)
		return nil, util.ErrResvCanceled
	}

	resv.Status = 1

	resv, err = s.dal.UpdateResvStatus(resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}
