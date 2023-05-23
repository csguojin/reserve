package service

import (
	"github.com/csguojin/reserve/dal"
	"github.com/csguojin/reserve/model"
)

type Service interface {
	CreateUser(user *model.User) (*model.User, error)
	CheckUser(username string, password string) (*model.User, error)
	GenerateToken(user *model.User) (string, error)
	GetAllUsers() ([]*model.User, error)
	GetUser(userID int) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(userID int) error

	GetAllRooms() ([]*model.Room, error)
	CreateRoom(room *model.Room) (*model.Room, error)
	GetRoom(roomID int) (*model.Room, error)
	UpdateRoom(room *model.Room) (*model.Room, error)
	DeleteRoom(roomID int) error

	GetAllSeats(roomID int) ([]*model.Seat, error)
	CreateSeat(seat *model.Seat) (*model.Seat, error)
	GetSeat(seatID int) (*model.Seat, error)
	UpdateSeat(seat *model.Seat) (*model.Seat, error)
	DeleteSeat(seatID int) error

	CreateResv(resv *model.Resv) (*model.Resv, error)
	GetResv(resvID int) (*model.Resv, error)
	UpdateResvStatus(resv *model.Resv) (*model.Resv, error)
	GetResvsByUser(userID int) ([]*model.Resv, error)
	GetResvsBySeat(seatID int) ([]*model.Resv, error)
	Signin(resvID int, userID int) (*model.Resv, error)
	Signout(resvID int, userID int) (*model.Resv, error)
	CancelResv(resvID int, userID int) (*model.Resv, error)

	CreateAdmin(admin *model.Admin) (*model.Admin, error)
	CheckAdmin(adminname string, password string) (*model.Admin, error)
	GenerateAdminToken(admin *model.Admin) (string, error)
	GetAllAdmins() ([]*model.Admin, error)
	GetAdminNoPassword(adminID int) (*model.Admin, error)
	DeleteAdmin(adminID int) error
}

func NewService(dalClient dal.Dal) Service {
	return &svc{
		dal: dalClient,
	}
}

type svc struct {
	dal dal.Dal
}
