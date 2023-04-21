package util

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserAuthFail = errors.New("username or password error")
)
