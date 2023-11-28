package database

import "errors"

var ErrUserDoesNotExist = errors.New("the requested user does not exist")
var ErrPhotoDoesNotExist = errors.New("the requested photo does not exist")
var ErrUserNotFollowed = errors.New("the second user was not followed by the first user")
var ErrUserNotBanned = errors.New("the second user was not banned by the first user")