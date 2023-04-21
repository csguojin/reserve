package service

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
)

func CreateUser(user *model.User) (*model.User, error) {
	passwordData, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordData)
	return dal.CeateUser(dal.GetDB(), user)
}

func CheckUser(username string, password string) (*model.User, error) {
	user, err := dal.GetUserByName(dal.GetDB(), username)
	if err != nil {
		return nil, err
	}
	err = verifyPassword(password, user.Password)
	if err != nil {
		log.Println(err)
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
		return "", err
	}

	return tokenStr, nil
}
