package dal

import (
	"context"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (d *dal) GetAllUsers(ctx context.Context, pager *model.Pager) ([]*model.User, error) {
	var users []*model.User
	offset := (pager.Page - 1) * pager.PerPage
	err := d.db.Offset(offset).Limit(pager.PerPage).Select("id", "username", "email").Find(&users).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return users, nil
}

func (d *dal) CeateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := d.db.Create(user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return d.GetUser(ctx, user.ID)
}

func (d *dal) GetUser(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{ID: id}
	err := d.db.Select("id", "username", "email").First(&user, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

func (d *dal) GetUserByName(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := d.db.Select("id", "username", "email").Where(&model.User{Username: username}).First(user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

func (d *dal) GetUserWithPasswordByName(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := d.db.Where(&model.User{Username: username}).First(user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

func (d *dal) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := d.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return d.GetUser(ctx, user.ID)
}

func (d *dal) DeleteUser(ctx context.Context, userID int) error {
	err := d.db.Delete(&model.User{}, userID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
