package dal

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (d *dal) GetAllRooms(ctx context.Context, pager *model.Pager) ([]*model.Room, error) {
	var rooms []*model.Room
	offset := (pager.Page - 1) * pager.PerPage
	err := d.db.Offset(offset).Limit(pager.PerPage).Find(&rooms).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return rooms, nil
}

func (d *dal) CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	err := d.db.Create(room).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	data, err := json.Marshal(room)
	if err != nil {
		logger.L.Warnln(err)
		return room, nil
	}
	roomCacheKey := fmt.Sprintf("%s%d", "room:", room.ID)
	err = d.rdb.Set(ctx, roomCacheKey, data, redisTTL).Err()
	if err != nil {
		logger.L.Warnln(err)
	}
	return room, nil
}

func (d *dal) GetRoom(ctx context.Context, id int) (*model.Room, error) {
	roomCacheKey := fmt.Sprintf("%s%d", "room:", id)
	roomJSON, err := d.rdb.Get(ctx, roomCacheKey).Result()
	if err == nil {
		var room *model.Room
		err := json.Unmarshal([]byte(roomJSON), &room)
		if err == nil {
			return room, nil
		}
	}

	room := &model.Room{ID: id}
	err = d.db.First(&room, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrRoomNotFound
	}

	data, err := json.Marshal(room)
	if err != nil {
		logger.L.Warnln(err)
		return room, nil
	}
	err = d.rdb.Set(ctx, roomCacheKey, data, redisTTL).Err()
	if err != nil {
		logger.L.Warnln(err)
	}

	return room, nil
}

func (d *dal) UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	err := d.db.Model(&model.Room{}).Where("id = ?", room.ID).Updates(&room).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	roomCacheKey := fmt.Sprintf("%s%d", "room:", room.ID)
	err = d.rdb.Del(ctx, roomCacheKey).Err()
	if err != nil {
		logger.L.Errorln(err)
		return room, err
	}
	return room, nil
}

func (d *dal) DeleteRoom(ctx context.Context, roomID int) error {
	err := d.db.Delete(&model.Room{}, roomID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}

	roomCacheKey := fmt.Sprintf("%s%d", "room:", roomID)
	err = d.rdb.Del(ctx, roomCacheKey).Err()
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
