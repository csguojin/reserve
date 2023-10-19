package dal

import (
	"context"

	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
)

type Dal interface {
	GetAllUsers(ctx context.Context, pager *model.Pager) ([]*model.User, error)
	CeateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserByName(ctx context.Context, username string) (*model.User, error)
	GetUserWithPasswordByName(ctx context.Context, username string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, userID int) error

	GetAllRooms(ctx context.Context, pager *model.Pager) ([]*model.Room, error)
	CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	GetRoom(ctx context.Context, id int) (*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	DeleteRoom(ctx context.Context, roomID int) error

	GetAllSeats(ctx context.Context, roomID int, pager *model.Pager) ([]*model.Seat, error)
	CreateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error)
	GetSeat(ctx context.Context, id int) (*model.Seat, error)
	UpdateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error)
	DeleteSeat(ctx context.Context, seatID int) error

	CreateResv(ctx context.Context, resv *model.Resv) (*model.Resv, error)
	GetResv(ctx context.Context, resvID int) (*model.Resv, error)
	UpdateResvStatus(ctx context.Context, resv *model.Resv) (*model.Resv, error)
	GetResvsByUser(ctx context.Context, userID int, pager *model.Pager) ([]*model.Resv, error)
	GetResvsBySeat(ctx context.Context, seatID int, pager *model.Pager) ([]*model.Resv, error)

	GetAllAdmins(ctx context.Context, pager *model.Pager) ([]*model.Admin, error)
	CeateAdmin(ctx context.Context, admin *model.Admin) (*model.Admin, error)
	GetAdmin(ctx context.Context, id int) (*model.Admin, error)
	GetAdminByName(ctx context.Context, name string) (*model.Admin, error)
	GetAdminWithPasswordByName(ctx context.Context, name string) (*model.Admin, error)
	DeleteAdmin(ctx context.Context, adminID int) error
}

func GetDal(db *gorm.DB) Dal {
	return &dal{db: db}
}

type dal struct {
	db *gorm.DB
}
