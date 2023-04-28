package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"uniqueIndex;size:255;not null"`
	Email    string `json:"email"  gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

func (User) TableName() string {
	return "user_table"
}
