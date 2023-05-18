package util

import "errors"

var (
	ErrRequestBodyFormat = errors.New("body is missing or illegal")
	ErrTokenInvalid      = errors.New("invalid token")
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserAuthFail = errors.New("username or password error")
	ErrUserIDNil    = errors.New("user id is nil")
)

var (
	ErrRoomNotFound = errors.New("room not found")
)

var (
	ErrSeatNotFound = errors.New("seat not found")
)
