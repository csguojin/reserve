package util

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserAuthFail = errors.New("username or password error")
)

var (
	ErrRoomNotFound = errors.New("room not found")
)
