package service

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func (s *svc) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	encryedPwd, err := encryptPassword(ctx, user.Password)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	user.Password = encryedPwd
	user, err = s.dal.CeateUser(ctx, user)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return user, nil
}

func (s *svc) CheckUser(ctx context.Context, username string, password string) (*model.User, error) {
	user, err := s.dal.GetUserWithPasswordByName(ctx, username)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	err = verifyPassword(ctx, password, user.Password)
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrUserAuthFail
	}
	return user, nil
}

func (s *svc) GenerateToken(ctx context.Context, user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),

		"ip":       ctx.Value("ip"),
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

func (s *svc) GetAllUsers(ctx context.Context, pager *model.Pager) ([]*model.User, error) {
	users, err := s.dal.GetAllUsers(ctx, pager)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return users, nil
}

func (s *svc) GetUser(ctx context.Context, userID int) (*model.User, error) {
	user, err := s.dal.GetUser(ctx, userID)
	user.Password = ""
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return user, nil
}

func (s *svc) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if user.Password != "" {
		encryedPwd, err := encryptPassword(ctx, user.Password)
		if err != nil {
			logger.L.Errorln(err)
			return nil, err
		}
		user.Password = encryedPwd
	}
	user, err := s.dal.UpdateUser(ctx, user)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return user, nil
}

func (s *svc) DeleteUser(ctx context.Context, userID int) error {
	err := s.dal.DeleteUser(ctx, userID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}

func encryptPassword(ctx context.Context, password string) (string, error) {
	passwordData, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.L.Errorln(err)
		return "", err
	}
	return string(passwordData), nil
}

func verifyPassword(ctx context.Context, password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
