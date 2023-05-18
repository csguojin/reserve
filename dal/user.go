package dal

import (
	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func GetAllUsers(db *gorm.DB) ([]*model.User, error) {
	var users []*model.User
	err := db.Find(&users).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return users, nil
}

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

func DeleteUser(db *gorm.DB, userID int) error {
	err := db.Delete(&model.User{}, userID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
