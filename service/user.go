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

func (s *svc) CreateUser(user *model.User) (*model.User, error) {
	encryedPwd, err := encryptPassword(user.Password)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	user.Password = encryedPwd
	user, err = s.dal.CeateUser(user)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return user, nil
}

func (s *svc) CheckUser(username string, password string) (*model.User, error) {
	user, err := s.dal.GetUserWithPasswordByName(username)
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

func (s *svc) GenerateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),

		"userid":   strconv.Itoa(user.ID),
		"username": user.Username,
	})
	tokenStr, err := token.SignedString([]byte(viper.GetString("jwt.userkey")))
	if err != nil {
		logger.L.Errorln(err)
		return "", err
	}
	return tokenStr, nil
}

func (s *svc) GetAllUsers() ([]*model.User, error) {
	users, err := s.dal.GetAllUsers()
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return users, nil
}

func (s *svc) GetUser(userID int) (*model.User, error) {
	user, err := s.dal.GetUser(userID)
	user.Password = ""
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return user, nil
}

func (s *svc) UpdateUser(user *model.User) (*model.User, error) {
	if user.Password != "" {
		encryedPwd, err := encryptPassword(user.Password)
		if err != nil {
			logger.L.Errorln(err)
			return nil, err
		}
		user.Password = encryedPwd
	}
	user, err := s.dal.UpdateUser(user)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return user, nil
}

func (s *svc) DeleteUser(userID int) error {
	err := s.dal.DeleteUser(userID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}

func encryptPassword(password string) (string, error) {
	passwordData, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.L.Errorln(err)
		return "", err
	}
	return string(passwordData), nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
