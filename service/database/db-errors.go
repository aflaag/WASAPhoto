package database

import (
	"errors"
)

var ErrUserDoesNotExist = errors.New("The requested user does not exist.")
var ErrUserNotFollowed = errors.New("The second user was not followed be the first user.")