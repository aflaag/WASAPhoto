package api

import (
	"errors"
)

// User
var ErrUserDoesNotExist = errors.New("the requested user does not exist")
var ErrUserUnauthorized = errors.New("the requested user is not authorized to perform this action")

// Ban
var ErrBannedUser = errors.New("the requested has banned the user performing the action")

// Others
var ErrPageNotFound = errors.New("the requested resource does not exist")
