package dal

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (d *dal) GetAllSeats(ctx context.Context, roomID int, pager *model.Pager) ([]*model.Seat, error) {
	var seats []*model.Seat
	offset := (pager.Page - 1) * pager.PerPage
	err := d.db.Offset(offset).Limit(pager.PerPage).Where(&model.Seat{RoomID: roomID}).Find(&seats).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return seats, nil
}

func (d *dal) CreateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	err := d.db.Create(seat).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	data, err := json.Marshal(seat)
	if err != nil {
		logger.L.Warnln(err)
		return seat, nil
	}
	seatCacheKey := fmt.Sprintf("%s%d", "seat:", seat.ID)
	err = d.rdb.Set(ctx, seatCacheKey, data, redisTTL).Err()
	if err != nil {
		logger.L.Warnln(err)
		return seat, nil
	}
	return seat, nil
}

func (d *dal) GetSeat(ctx context.Context, id int) (*model.Seat, error) {
	seatCacheKey := fmt.Sprintf("%s%d", "seat:", id)
	seatJSON, err := d.rdb.Get(ctx, seatCacheKey).Result()
	if err == nil {
		var seat *model.Seat
		err := json.Unmarshal([]byte(seatJSON), &seat)
		if err == nil {
			return seat, nil
		}
	}

	seat := &model.Seat{ID: id}
	err = d.db.First(&seat, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrSeatNotFound
	}

	data, err := json.Marshal(seat)
	if err != nil {
		logger.L.Warnln(err)
		return seat, nil
	}
	err = d.rdb.Set(ctx, seatCacheKey, data, redisTTL).Err()
	if err != nil {
		logger.L.Warnln(err)
		return seat, nil
	}

	return seat, nil
}

func (d *dal) UpdateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	err := d.db.Model(&model.Seat{}).Where("id = ?", seat.ID).Updates(&seat).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	seatCacheKey := fmt.Sprintf("%s%d", "seat:", seat.ID)
	err = d.rdb.Del(ctx, seatCacheKey).Err()
	if err != nil {
		logger.L.Errorln(err)
		return seat, err
	}

	return d.GetSeat(ctx, seat.ID)
}

func (d *dal) DeleteSeat(ctx context.Context, seatID int) error {
	err := d.db.Delete(&model.Seat{}, seatID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}

	seatCacheKey := fmt.Sprintf("%s%d", "seat:", seatID)
	err = d.rdb.Del(ctx, seatCacheKey).Err()
	if err != nil {
		logger.L.Errorln(err)
		return err
	}

	return nil
}
