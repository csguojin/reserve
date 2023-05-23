package dal

import (
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (d *dal) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	err := d.db.Select("id", "username", "email").Find(&users).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return users, nil
}

func (d *dal) CeateUser(user *model.User) (*model.User, error) {
	err := d.db.Create(user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return d.GetUser(user.ID)
}

func (d *dal) GetUser(id int) (*model.User, error) {
	user := &model.User{ID: id}
	err := d.db.Select("id", "username", "email").First(&user, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

func (d *dal) GetUserByName(username string) (*model.User, error) {
	user := &model.User{}
	err := d.db.Select("id", "username", "email").Where(&model.User{Username: username}).First(user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

func (d *dal) GetUserWithPasswordByName(username string) (*model.User, error) {
	user := &model.User{}
	err := d.db.Where(&model.User{Username: username}).First(user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

func (d *dal) UpdateUser(user *model.User) (*model.User, error) {
	err := d.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return d.GetUser(user.ID)
}

func (d *dal) DeleteUser(userID int) error {
	err := d.db.Delete(&model.User{}, userID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
