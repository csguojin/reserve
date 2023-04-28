package model

type Room struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	OpenTime    string `json:"opening_time"`
	CloseTime   string `json:"closing_time"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (Room) TableName() string {
	return "room_table"
}
