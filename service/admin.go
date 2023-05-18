package service

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func CreateAdmin(admin *model.Admin) (*model.Admin, error) {
	passwordData, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	admin.Password = string(passwordData)
	return dal.CeateAdmin(dal.GetDB(), admin)
}

func CheckAdmin(adminname string, password string) (*model.Admin, error) {
	admin, err := dal.GetAdminWithPasswordByName(dal.GetDB(), adminname)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	err = verifyAdminPassword(password, admin.Password)
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserAuthFail
	}
	return admin, nil
}

func GenerateAdminToken(admin *model.Admin) (string, error) {
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

func GetAllAdmins() ([]*model.Admin, error) {
	admins, err := dal.GetAllAdmins(dal.GetDB())
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admins, nil
}

func GetAdminNoPassword(adminID int) (*model.Admin, error) {
	admin, err := dal.GetAdmin(dal.GetDB(), adminID)
	admin.Password = ""
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admin, nil
}

func DeleteAdmin(adminID int) error {
	err := dal.DeleteAdmin(dal.GetDB(), adminID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}

func verifyAdminPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
