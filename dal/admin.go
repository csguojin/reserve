package dal

import (
	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func GetAllAdmins(db *gorm.DB) ([]*model.Admin, error) {
	var admins []*model.Admin
	err := db.Select("id", "name", "email").Find(&admins).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admins, nil
}

func CeateAdmin(db *gorm.DB, admin *model.Admin) (*model.Admin, error) {
	err := db.Create(admin).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return GetAdmin(db, admin.ID)
}

func GetAdmin(db *gorm.DB, id int) (*model.Admin, error) {
	admin := &model.Admin{ID: id}
	err := db.Select("id", "name", "email").First(&admin, id).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return admin, nil
}

func GetAdminByName(db *gorm.DB, name string) (*model.Admin, error) {
	admin := &model.Admin{}
	err := db.Select("id", "name", "email").Where(&model.Admin{Name: name}).First(admin).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return admin, nil
}

func GetAdminWithPasswordByName(db *gorm.DB, name string) (*model.Admin, error) {
	admin := &model.Admin{}
	err := db.Where(&model.Admin{Name: name}).First(admin).Error
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserNotFound
	}
	return admin, nil
}

func DeleteAdmin(db *gorm.DB, adminID int) error {
	err := db.Delete(&model.Admin{}, adminID).Error
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
