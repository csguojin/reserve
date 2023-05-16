package model

import (
	"time"
)

type Resv struct {
	ID          int        `json:"id" gorm:"primaryKey; autoIncrement"`
	UserID      int        `json:"user_id" gorm:"index"`
	SeatID      int        `json:"seat_id" gorm:"index"`
	CreateTime  *time.Time `json:"create_time"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	SigninTime  *time.Time `json:"signin_time"`
	SignoutTime *time.Time `json:"signout_time"`
	Status      int        `json:"status"`
}

func (Resv) TableName() string {
	return "reservation_table"
}
