package database

import "errors"

// User
var ErrUserDoesNotExist = errors.New("the requested user does not exist")

// Follow
var ErrUserNotFollowed = errors.New("the second user was not followed by the first user")

// Ban
var ErrUserNotBanned = errors.New("the second user was not banned by the first user")

// Photo
var ErrPhotoDoesNotExist = errors.New("the requested photo does not exist")

// Like
var ErrPhotoNotLiked = errors.New("the requested photo was not liked by the given user")

// Comment
var ErrCommentDoesNotExist = errors.New("the requested comment does not exist")
var ErrPhotoNotCommented = errors.New("the requested photo was not commented by the given user")
