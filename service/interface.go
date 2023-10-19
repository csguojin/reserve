package service

import (
	"context"

	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
)

type Service interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	CheckUser(ctx context.Context, username string, password string) (*model.User, error)
	GenerateToken(ctx context.Context, user *model.User) (string, error)
	GetAllUsers(ctx context.Context, pager *model.Pager) ([]*model.User, error)
	GetUser(ctx context.Context, userID int) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, userID int) error

	GetAllRooms(ctx context.Context, pager *model.Pager) ([]*model.Room, error)
	CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	GetRoom(ctx context.Context, roomID int) (*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	DeleteRoom(ctx context.Context, roomID int) error

	GetAllSeats(ctx context.Context, roomID int, pager *model.Pager) ([]*model.Seat, error)
	CreateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error)
	GetSeat(ctx context.Context, seatID int) (*model.Seat, error)
	UpdateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error)
	DeleteSeat(ctx context.Context, seatID int) error

	CreateResv(ctx context.Context, resv *model.Resv) (*model.Resv, error)
	GetResv(ctx context.Context, resvID int) (*model.Resv, error)
	UpdateResvStatus(ctx context.Context, resv *model.Resv) (*model.Resv, error)
	GetResvsByUser(ctx context.Context, userID int, pager *model.Pager) ([]*model.Resv, error)
	GetResvsBySeat(ctx context.Context, seatID int, pager *model.Pager) ([]*model.Resv, error)
	Signin(ctx context.Context, resvID int, userID int) (*model.Resv, error)
	Signout(ctx context.Context, resvID int, userID int) (*model.Resv, error)
	CancelResv(ctx context.Context, resvID int, userID int) (*model.Resv, error)

	CreateAdmin(ctx context.Context, admin *model.Admin) (*model.Admin, error)
	CheckAdmin(ctx context.Context, adminname string, password string) (*model.Admin, error)
	GenerateAdminToken(ctx context.Context, admin *model.Admin) (string, error)
	GetAllAdmins(ctx context.Context, pager *model.Pager) ([]*model.Admin, error)
	GetAdminNoPassword(ctx context.Context, adminID int) (*model.Admin, error)
	DeleteAdmin(ctx context.Context, adminID int) error
}

func NewService(dalClient dal.Dal) Service {
	return &svc{
		dal: dalClient,
	}
}

type svc struct {
	dal dal.Dal
}
