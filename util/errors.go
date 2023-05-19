package util

import "errors"

var (
	ErrTokenInvalid = errors.New("invalid token")
)

var (
	ErrUserIDNil    = errors.New("user id is nil")
	ErrUserNotFound = errors.New("user not found")
	ErrUserAuthFail = errors.New("username or password error")
)

var (
	ErrRoomIDNil    = errors.New("room id is nil")
	ErrRoomNotFound = errors.New("room not found")
)

var (
	ErrSeatIDNil    = errors.New("seat id is nil")
	ErrSeatNotFound = errors.New("seat not found")
)

var (
	ErrResvIDNil            = errors.New("reservation id is nil")
	ErrResvTimeIllegal      = errors.New("reservation time is illegal")
	ErrResvCanceled         = errors.New("reservation has been cancel")
	ErrResvSeatTimeConflict = errors.New("seat reservation time conflict")
	ErrResvUserTimeConflict = errors.New("user reservation time conflict")
)

var (
	ErrAdminIDNil    = errors.New("admin id is nil")
	ErrAdminNotFound = errors.New("admin not found")
	ErrAdminAuthFail = errors.New("admin name or password error")
)
