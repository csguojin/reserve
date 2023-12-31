package model

type Seat struct {
	ID          int    `json:"id" gorm:"primaryKey; autoIncrement"`
	RoomID      int    `json:"room_id" gorm:"index; not null"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (Seat) TableName() string {
	return "seat_table"
}
