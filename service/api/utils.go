package api

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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

func (rt *_router) GetUserFromLogin(login Login) (User, error) {
	dbUser, err := rt.db.GetDatabaseUserFromDatabaseLogin(login.LoginIntoDatabaseLogin())

	if err != nil {
		return UserDefault(), err
	}

	user := UserFromDatabaseUser(dbUser)

	return user, nil
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

func (rt *_router) GetUserFromParameter(parameter string, r *http.Request, ps httprouter.Params) (User, int, error) {
	userUsername := ps.ByName(parameter)
	userLogin := LoginFromUsername(userUsername)

	user, err := rt.GetUserFromLogin(userLogin)

	code := -1

	if err != nil {
		code = http.StatusInternalServerError
	}

	return user, code, err
}

func (rt *_router) GetPhotoFromParameter(parameter string, r *http.Request, ps httprouter.Params) (Photo, int, error) {
	photo := PhotoDefault()

	photoIdString := ps.ByName(parameter)
	photoId, err := strconv.ParseUint(photoIdString, 10, 64)

	if err != nil {
		return photo, http.StatusInternalServerError, err
	}

	photo, err = rt.GetPhotoFromPhotoId(uint32(photoId))

	if err != nil {
		return photo, http.StatusInternalServerError, err
	}

	return photo, -1, nil
}

func (rt *_router) AuthenticateUserFromParameter(parameter string, r *http.Request, ps httprouter.Params) (User, int, error) {
	user, code, err := rt.GetUserFromParameter(parameter, r, ps)

	if err != nil {
		return user, code, err
	}

	err = CheckAuthorization(user, r.Header.Get("Authorization"))

	if err != nil {
		code = http.StatusUnauthorized
	}

	return user, code, err
}
