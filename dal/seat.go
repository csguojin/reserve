package dal

import (
	"context"

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
	return d.GetSeat(ctx, seat.ID)
}

func (d *dal) GetSeat(ctx context.Context, id int) (*model.Seat, error) {
	seat := &model.Seat{ID: id}
	err := d.db.First(&seat, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrSeatNotFound
	}
	return seat, nil
}

func (d *dal) UpdateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	err := d.db.Model(&model.Seat{}).Where("id = ?", seat.ID).Updates(&seat).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return d.GetSeat(ctx, seat.ID)
}

func (d *dal) DeleteSeat(ctx context.Context, seatID int) error {
	err := d.db.Delete(&model.Seat{}, seatID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
