package api

import (
	"errors"
)

var ErrUserUnauthorized = errors.New("the user is not authorized to perform this action")