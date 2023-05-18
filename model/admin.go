package model

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"uniqueIndex;size:255;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email"  gorm:"not null"`
}

func (Admin) TableName() string {
	return "admin_table"
}
