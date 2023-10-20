package dal

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm/clause"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (d *dal) CreateResv(ctx context.Context, resv *model.Resv) (*model.Resv, error) {
	if d.rdb == nil {
		return d.createResvOnlyMySQL(ctx, resv)
	}

	lockKeyUser := fmt.Sprintf("lock:user:%d", resv.UserID)
	lockValue := generateRandomString(16)
	lockAcquired, err := d.rdb.SetNX(ctx, lockKeyUser, lockValue, time.Minute).Result()
	if err != nil {
		logger.L.Errorln("Error acquiring lock:", err)
		return nil, err
	}
	if !lockAcquired {
		fmt.Println("Failed to acquire user lock. Cur user is performing an operation.")
		return nil, err
	}

	lockKeySeat := fmt.Sprintf("lock:seat:%d", resv.SeatID)
	lockAcquired, err = d.rdb.SetNX(ctx, lockKeySeat, lockValue, time.Minute).Result()
	if err != nil {
		logger.L.Errorln("Error acquiring lock:", err)
		return nil, err
	}
	if !lockAcquired {
		fmt.Println("Failed to acquire seat lock. Cur seat is performing an operation.")
		return nil, err
	}

	defer func() {
		script := `
			if redis.call("get", KEYS[1]) == ARGV[1] then
				return redis.call("del", KEYS[1])
			else
				return 0
			end
		`

		_, err := d.rdb.Eval(ctx, script, []string{lockKeySeat}, lockValue).Result()
		if err != nil {
			logger.L.Errorln("Error releasing lock:", err)
		}

		_, err = d.rdb.Eval(ctx, script, []string{lockKeyUser}, lockValue).Result()
		if err != nil {
			logger.L.Errorln("Error releasing lock:", err)
		}
	}()

	seatBitKey := fmt.Sprintf("seat:%d:%s", resv.SeatID, resv.StartTime.Format("20060102"))
	userKey := fmt.Sprintf("user:%d:%s", resv.UserID, resv.StartTime.Format("20060102"))

	luaScript := `
		local seatKey = KEYS[1]
		local userKey = KEYS[2]
		local startOffset = tonumber(ARGV[1])
		local endOffset = tonumber(ARGV[2])
		
		local seatBitCount = redis.call('BITCOUNT', seatKey, startOffset, endOffset, 'BIT')
		local userBitCount = redis.call('BITCOUNT', userKey, startOffset, endOffset, 'BIT')
		
		if seatBitCount == 0 and userBitCount == 0 then
			local bitfieldArgs = {}

			for i = startOffset, endOffset do
				table.insert(bitfieldArgs, "SET")
				table.insert(bitfieldArgs, "u1")
				table.insert(bitfieldArgs, i)
				table.insert(bitfieldArgs, 1)
			end

			redis.call('BITFIELD', seatKey, unpack(bitfieldArgs))
			redis.call('BITFIELD', userKey, unpack(bitfieldArgs))
		
			return 'OK'
		else
			return 'BITCOUNT not equal to 0'
		end
	`

	startOffset, endOffset := resv.CalculateTimeBits(*resv.StartTime, *resv.EndTime)

	result, err := d.rdb.Eval(ctx, luaScript, []string{seatBitKey, userKey}, startOffset, endOffset).Result()
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}

	if result.(string) != "OK" {
		err = errors.New(result.(string))
		logger.L.Errorln(err)
		return nil, err
	}

	err = d.db.Create(resv).Error
	if err != nil {
		logger.L.Errorln(err)

		// Handle rollback Redis operations
		luaScript = `
			local seatKey = KEYS[1]
			local userKey = KEYS[2]
			local startOffset = tonumber(ARGV[1])
			local endOffset = tonumber(ARGV[2])

			local bitfieldArgs = {}

			for i = startOffset, endOffset do
				table.insert(bitfieldArgs, "SET")
				table.insert(bitfieldArgs, "u1")
				table.insert(bitfieldArgs, i)
				table.insert(bitfieldArgs, 0)
			end

			redis.call('BITFIELD', seatKey, unpack(bitfieldArgs))
			redis.call('BITFIELD', userKey, unpack(bitfieldArgs))

			return 'OK'
		`

		_, err2 := d.rdb.Eval(ctx, luaScript, []string{seatBitKey, userKey}, startOffset, endOffset).Result()
		if err2 != nil {
			logger.L.Errorln(err)
			err = errors.New(err.Error() + err2.Error())
		}

		return nil, err
	}

	return resv, nil
}

func (d *dal) createResvOnlyMySQL(ctx context.Context, resv *model.Resv) (*model.Resv, error) {
	tx := d.db.Begin()
	if tx.Error != nil {
		logger.L.Errorln(tx.Error)
		return nil, tx.Error
	}

	var seat model.Seat
	if err := tx.
		Clauses(clause.Locking{Strength: "SHARE"}).
		Where("id = ?", resv.SeatID).
		Where("status = ?", 0).
		First(&seat).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if seat.Status != 0 {
		tx.Rollback()
		err := errors.New("seat is not available")
		logger.L.Errorln(err)
		return nil, err
	}

	var room *model.Room
	if err := tx.
		Clauses(clause.Locking{Strength: "SHARE"}).
		Where("id = ?", seat.RoomID).
		Where("status = ?", 0).
		First(&room).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if ok := checkRoomResv(ctx, room, resv); !ok {
		tx.Rollback()
		err := errors.New("room is not available")
		logger.L.Errorln(err)
		return nil, err
	}

	var existingResvCount int64
	if err := tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&model.Resv{}).
		Where("seat_id = ?", resv.SeatID).
		Where("status = ?", 0).
		Where("start_time < ? AND end_time > ?", resv.EndTime, resv.StartTime).
		Count(&existingResvCount).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if existingResvCount > 0 {
		tx.Rollback()
		logger.L.Errorln(util.ErrResvSeatTimeConflict)
		return nil, util.ErrResvSeatTimeConflict
	}

	existingResvCount = 0
	if err := tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&model.Resv{}).
		Where("user_id = ?", resv.UserID).
		Where("status = ?", 0).
		Where("start_time < ? AND end_time > ?", resv.EndTime, resv.StartTime).
		Count(&existingResvCount).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if existingResvCount > 0 {
		tx.Rollback()
		logger.L.Errorln(util.ErrResvUserTimeConflict)
		return nil, util.ErrResvUserTimeConflict
	}

	if err := tx.Create(resv).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	return resv, nil
}

func checkRoomResv(ctx context.Context, room *model.Room, resv *model.Resv) bool {
	openTime, err := time.Parse(time.RFC3339, resv.StartTime.Format("2006-01-02T")+room.OpenTime+"+08:00")
	if err != nil {
		logger.L.Errorln("room opentime format error", room.OpenTime, err)
		return false
	}

	closeTime, err := time.Parse(time.RFC3339, resv.StartTime.Format("2006-01-02T")+room.CloseTime+"+08:00")
	if err != nil {
		logger.L.Errorln("room closetime format error", room.CloseTime, err)
		return false
	}
	return resv.StartTime.Before(*resv.EndTime) && resv.StartTime.After(openTime) && resv.EndTime.Before(closeTime)
}

func (d *dal) GetResv(ctx context.Context, resvID int) (*model.Resv, error) {
	resv := &model.Resv{ID: resvID}
	err := d.db.First(&resv, resvID).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (d *dal) UpdateResvStatus(ctx context.Context, resv *model.Resv) (*model.Resv, error) {
	err := d.db.Model(&model.Resv{}).Where("id = ?", resv.ID).Updates(
		&model.Resv{
			SigninTime:  resv.SigninTime,
			SignoutTime: resv.SignoutTime,
			Status:      resv.Status,
		}).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return d.GetResv(ctx, resv.ID)
}

func (d *dal) UpdateResvStartEndTime(ctx context.Context, newResv *model.Resv) (*model.Resv, error) {
	tx := d.db.Begin()
	if tx.Error != nil {
		logger.L.Errorln(tx.Error)
		return nil, tx.Error
	}

	oldResv := &model.Resv{ID: newResv.ID}
	err := tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&oldResv, newResv.ID).Error
	if err != nil {
		logger.L.Errorln(err)
		tx.Rollback()
		return nil, err
	}

	if newResv.SeatID != oldResv.SeatID {
		err = errors.New("resv seat cannot update")
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if newResv.UserID != oldResv.UserID {
		err = errors.New("resv user cannot update")
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	var existingResvCount int64
	if err := tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&model.Resv{}).
		Where("seat_id = ?", newResv.SeatID).
		Where("id != ?", newResv.ID).
		Where("status = ?", 0).
		Where("start_time < ? AND end_time > ?", newResv.EndTime, newResv.StartTime).
		Count(&existingResvCount).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if existingResvCount > 0 {
		tx.Rollback()
		logger.L.Errorln(util.ErrResvSeatTimeConflict)
		return nil, util.ErrResvSeatTimeConflict
	}

	existingResvCount = 0
	if err := tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&model.Resv{}).
		Where("user_id = ?", newResv.UserID).
		Where("id != ?", newResv.ID).
		Where("status = ?", 0).
		Where("start_time < ? AND end_time > ?", newResv.EndTime, newResv.StartTime).
		Count(&existingResvCount).Error; err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	if existingResvCount > 0 {
		tx.Rollback()
		logger.L.Errorln(util.ErrResvUserTimeConflict)
		return nil, util.ErrResvUserTimeConflict
	}

	err = tx.Model(&model.Resv{}).Where("id = ?", newResv.ID).Updates(
		&model.Resv{
			StartTime: newResv.StartTime,
			EndTime:   newResv.EndTime,
		}).Error
	if err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		logger.L.Errorln(err)
		return nil, err
	}

	return d.GetResv(ctx, newResv.ID)
}

func (d *dal) GetResvsByUser(ctx context.Context, userID int, pager *model.Pager) ([]*model.Resv, error) {
	var resvs []*model.Resv
	offset := (pager.Page - 1) * pager.PerPage
	err := d.db.Offset(offset).Limit(pager.PerPage).Where(&model.Resv{UserID: userID}).Find(&resvs).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resvs, nil
}

func (d *dal) GetResvsBySeat(ctx context.Context, seatID int, pager *model.Pager) ([]*model.Resv, error) {
	var resvs []*model.Resv
	offset := (pager.Page - 1) * pager.PerPage
	err := d.db.Offset(offset).Limit(pager.PerPage).Where(&model.Resv{SeatID: seatID}).Find(&resvs).Offset(offset).Limit(pager.PerPage).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return resvs, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
