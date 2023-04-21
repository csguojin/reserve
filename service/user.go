package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
)

func CreateUser(user *model.User) (*model.User, error) {
	passwordData, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordData)
	return dal.CeateUser(dal.GetDB(), user)
}
