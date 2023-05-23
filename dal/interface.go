package dal

import (
	"gorm.io/gorm"

	"github.com/csguojin/reserve/model"
)

type Dal interface {
	GetAllUsers(pager *model.Pager) ([]*model.User, error)
	CeateUser(user *model.User) (*model.User, error)
	GetUser(id int) (*model.User, error)
	GetUserByName(username string) (*model.User, error)
	GetUserWithPasswordByName(username string) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(userID int) error

	GetAllRooms(pager *model.Pager) ([]*model.Room, error)
	CreateRoom(room *model.Room) (*model.Room, error)
	GetRoom(id int) (*model.Room, error)
	UpdateRoom(room *model.Room) (*model.Room, error)
	DeleteRoom(roomID int) error

	GetAllSeats(roomID int, pager *model.Pager) ([]*model.Seat, error)
	CreateSeat(seat *model.Seat) (*model.Seat, error)
	GetSeat(id int) (*model.Seat, error)
	UpdateSeat(seat *model.Seat) (*model.Seat, error)
	DeleteSeat(seatID int) error

	CreateResv(resv *model.Resv) (*model.Resv, error)
	GetResv(resvID int) (*model.Resv, error)
	UpdateResvStatus(resv *model.Resv) (*model.Resv, error)
	GetResvsByUser(userID int, pager *model.Pager) ([]*model.Resv, error)
	GetResvsBySeat(seatID int, pager *model.Pager) ([]*model.Resv, error)

	GetAllAdmins(pager *model.Pager) ([]*model.Admin, error)
	CeateAdmin(admin *model.Admin) (*model.Admin, error)
	GetAdmin(id int) (*model.Admin, error)
	GetAdminByName(name string) (*model.Admin, error)
	GetAdminWithPasswordByName(name string) (*model.Admin, error)
	DeleteAdmin(adminID int) error
}

func GetDal(db *gorm.DB) Dal {
	return &dal{db: db}
}

type dal struct {
	db *gorm.DB
}
