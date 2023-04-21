package dal

import (
	"log"

	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
)

func CeateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return GetUser(db, user.ID)
}

func GetUser(db *gorm.DB, id int64) (*model.User, error) {
	user := &model.User{ID: id}
	err := db.First(&user, id).Error
	if err != nil {
		log.Println(err)
		return nil, util.ErrUserNotFound
	}
	return user, nil
}
