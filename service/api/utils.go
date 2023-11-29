package api

import (
	"regexp"
	"strconv"
)

func (rt *_router) GetUserFromLogin(login Login) (User, error) {
	dbUser, err := rt.db.GetDatabaseUserFromDatabaseLogin(login.LoginIntoDatabaseLogin())

	if err != nil {
		return UserDefault(), err
	}

	user := UserFromDatabaseUser(dbUser)

	return user, nil
}

func CheckAuthorization(user User, authRaw string) error {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	tokenString := re.FindAllString(authRaw, -1)

	if len(tokenString) == 0 {
		return ErrUserUnauthorized
	}

	token, _ := strconv.Atoi(tokenString[0])

	if int(user.Id) != token {
		return ErrUserUnauthorized
	}

	return nil
}

func (rt *_router) GetPhotoFromPhotoId(photoId uint32) (Photo, error) {
	dbPhoto, err := rt.db.GetDatabasePhoto(photoId)

	if err != nil {
		return PhotoDefault(), err
	}

	photo := PhotoFromDatabasePhoto(dbPhoto)

	return photo, nil
}

func (rt *_router) GetCommentFromCommentId(commentId uint32) (Comment, error) {
	dbComment, err := rt.db.GetDatabaseComment(commentId)

	if err != nil {
		return CommentDefault(), err
	}

	comment := CommentFromDatabaseComment(dbComment)

	return comment, nil
}
