package service

import (
	"context"
	"errors"
	"time"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (s *svc) CreateResv(ctx context.Context, resv *model.Resv) (*model.Resv, error) {
	now := time.Now()
	resv.CreateTime = &now
	resv.Status = 0

	resv, err := s.dal.CreateResv(ctx, resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) GetResv(ctx context.Context, resvID int) (*model.Resv, error) {
	resv, err := s.dal.GetResv(ctx, resvID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) UpdateResvStatus(ctx context.Context, resv *model.Resv) (*model.Resv, error) {
	resv, err := s.dal.UpdateResvStatus(ctx, resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) GetResvsByUser(ctx context.Context, userID int, pager *model.Pager) ([]*model.Resv, error) {
	resv, err := s.dal.GetResvsByUser(ctx, userID, pager)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) GetResvsBySeat(ctx context.Context, seatID int, pager *model.Pager) ([]*model.Resv, error) {
	resv, err := s.dal.GetResvsBySeat(ctx, seatID, pager)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) Signin(ctx context.Context, resvID int, userID int) (*model.Resv, error) {
	resv, err := s.dal.GetResv(ctx, resvID)
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

	resv, err = s.dal.UpdateResvStatus(ctx, resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) Signout(ctx context.Context, resvID int, userID int) (*model.Resv, error) {
	resv, err := s.dal.GetResv(ctx, resvID)
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

	resv, err = s.dal.UpdateResvStatus(ctx, resv)
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

func (s *svc) CancelResv(ctx context.Context, resvID int, userID int) (*model.Resv, error) {
	resv, err := s.dal.GetResv(ctx, resvID)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	if resv.Status != 0 {
		logger.L.Errorln(util.ErrResvCanceled)
		return nil, util.ErrResvCanceled
	}

	resv.Status = 1

	resv, err = s.dal.UpdateResvStatus(ctx, resv)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}
