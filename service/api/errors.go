package api

import (
	"errors"
)

// User
var ErrUserDoesNotExist = errors.New("the requested user does not exist")
var ErrUserUnauthorized = errors.New("the requested user is not authorized to perform this action")
