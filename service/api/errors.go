package api

import (
	"errors"
)

// User
var ErrUserDoesNotExist = errors.New("the requested user does not exist")
var ErrUserUnauthorized = errors.New("the requested user is not authorized to perform this action")

// Ban
var ErrBannedUser = errors.New("the requested user has banned the user performing the action")
var ErrSelfBan = errors.New("the user performing the ban and the user to be banned are the same user")

// Follow
var ErrSelfFollow = errors.New("the user performing the following and the user to be followed are the same user")

// Others
var ErrPageNotFound = errors.New("the requested resource does not exist")
