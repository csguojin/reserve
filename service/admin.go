package service

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (s *svc) CreateAdmin(admin *model.Admin) (*model.Admin, error) {
	encryedPwd, err := encryptPassword(admin.Password)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	admin.Password = encryedPwd
	admin, err = s.dal.CeateAdmin(admin)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admin, nil
}

func (s *svc) CheckAdmin(adminname string, password string) (*model.Admin, error) {
	admin, err := s.dal.GetAdminWithPasswordByName(adminname)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	err = verifyAdminPassword(password, admin.Password)
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrAdminAuthFail
	}
	return admin, nil
}

func (s *svc) GenerateAdminToken(admin *model.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),

		"adminid":   strconv.Itoa(admin.ID),
		"adminname": admin.Name,
	})
	tokenStr, err := token.SignedString([]byte(viper.GetString("jwt.adminkey")))
	if err != nil {
		logger.L.Errorln(err)
		return "", err
	}

	return tokenStr, nil
}

func (s *svc) GetAllAdmins(pager *model.Pager) ([]*model.Admin, error) {
	admins, err := s.dal.GetAllAdmins(pager)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admins, nil
}

func (s *svc) GetAdminNoPassword(adminID int) (*model.Admin, error) {
	admin, err := s.dal.GetAdmin(adminID)
	admin.Password = ""
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admin, nil
}

func (s *svc) DeleteAdmin(adminID int) error {
	err := s.dal.DeleteAdmin(adminID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}

func verifyAdminPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
