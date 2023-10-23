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

func (s *svc) CreateAdmin(ctx context.Context, admin *model.Admin) (*model.Admin, error) {
	encryedPwd, err := encryptPassword(ctx, admin.Password)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	admin.Password = encryedPwd
	admin, err = s.dal.CeateAdmin(ctx, admin)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admin, nil
}

func (s *svc) CheckAdmin(ctx context.Context, adminname string, password string) (*model.Admin, error) {
	admin, err := s.dal.GetAdminWithPasswordByName(ctx, adminname)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	err = verifyAdminPassword(ctx, password, admin.Password)
	if err != nil {
		logger.L.Errorln(err)
		return nil, util.ErrAdminAuthFail
	}
	return admin, nil
}

func (s *svc) GenerateAdminToken(ctx context.Context, admin *model.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),

		"ip":        ctx.Value("ip"),
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

func (s *svc) GetAllAdmins(ctx context.Context, pager *model.Pager) ([]*model.Admin, error) {
	admins, err := s.dal.GetAllAdmins(ctx, pager)
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admins, nil
}

func (s *svc) GetAdminNoPassword(ctx context.Context, adminID int) (*model.Admin, error) {
	admin, err := s.dal.GetAdmin(ctx, adminID)
	admin.Password = ""
	if err != nil {
		logger.L.Errorln(err)
		return nil, err
	}
	return admin, nil
}

func (s *svc) DeleteAdmin(ctx context.Context, adminID int) error {
	err := s.dal.DeleteAdmin(ctx, adminID)
	if err != nil {
		logger.L.Errorln(err)
		return err
	}
	return nil
}

func verifyAdminPassword(ctx context.Context, password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
