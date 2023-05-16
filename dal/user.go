package dal

import (
	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func CeateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Create(user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return GetUser(db, user.ID)
}

func GetUser(db *gorm.DB, id int) (*model.User, error) {
	user := &model.User{ID: id}
	err := db.First(&user, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

func GetUserByName(db *gorm.DB, username string) (*model.User, error) {
	user := &model.User{}
	err := db.Where(&model.User{Username: username}).First(user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return user, nil
}
