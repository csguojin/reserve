package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func CreateUser(user *model.User) (*model.User, error) {
	passwordData, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	user.Password = string(passwordData)
	return dal.CeateUser(dal.GetDB(), user)
}

func CheckUser(username string, password string) (*model.User, error) {
	user, err := dal.GetUserByName(dal.GetDB(), username)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	err = verifyPassword(password, user.Password)
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserAuthFail
	}
	return user, nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(viper.GetString("jwt.key")))
	if err != nil {
		logger.L.Errorln(err)
		return "", err
	}

	return tokenStr, nil
}

func GetAllUsers() ([]*model.User, error) {
	users, err := dal.GetAllUsers(dal.GetDB())
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return users, nil
}

func GetUserNoPassword(userID int) (*model.User, error) {
	user, err := dal.GetUser(dal.GetDB(), userID)
	user.Password = ""
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return user, nil
}

func DeleteUser(userID int) error {
	err := dal.DeleteUser(dal.GetDB(), userID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}
