package dal

import (
	"context"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (d *dal) GetAllAdmins(ctx context.Context, pager *model.Pager) ([]*model.Admin, error) {
	var admins []*model.Admin
	offset := (pager.Page - 1) * pager.PerPage
	err := d.db.Offset(offset).Limit(pager.PerPage).Select("id", "name", "email").Find(&admins).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admins, nil
}

func (d *dal) CeateAdmin(ctx context.Context, admin *model.Admin) (*model.Admin, error) {
	err := d.db.Create(admin).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return d.GetAdmin(ctx, admin.ID)
}

func (d *dal) GetAdmin(ctx context.Context, id int) (*model.Admin, error) {
	admin := &model.Admin{ID: id}
	err := d.db.Select("id", "name", "email").First(&admin, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return admin, nil
}

func (d *dal) GetAdminByName(ctx context.Context, name string) (*model.Admin, error) {
	admin := &model.Admin{}
	err := d.db.Select("id", "name", "email").Where(&model.Admin{Name: name}).First(admin).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return admin, nil
}

func (d *dal) GetAdminWithPasswordByName(ctx context.Context, name string) (*model.Admin, error) {
	admin := &model.Admin{}
	err := d.db.Where(&model.Admin{Name: name}).First(admin).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return admin, nil
}

func (d *dal) DeleteAdmin(ctx context.Context, adminID int) error {
	err := d.db.Delete(&model.Admin{}, adminID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
