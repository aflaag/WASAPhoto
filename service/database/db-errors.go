package database

import (
	"errors"
)

var ErrUserDoesNotExist = errors.New("The requested user does not exist.")