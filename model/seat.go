package model

type Seat struct {
	ID          int64  `json:"id" gorm:"primaryKey; autoIncrement"`
	RoomID      int    `json:"room_id" gorm:"index; not null"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"URL"`
	Status      int    `json:"status"`
}

func (Seat) TableName() string {
	return "seat_table"
}
