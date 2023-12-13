package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get the bearer token
	token, err := GetBearerToken(r.Header.Get("Authorization"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// authenticate the user performing the action
	dbUser, err := rt.db.GetDatabaseUser(uint32(token))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the user of the photo from the resource parameter
	photoUser, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check whether the user of the photo
	// has banned the user performing the action
	checkBan, err := rt.db.CheckBan(photoUser.UserIntoDatabaseUser(), dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if checkBan {
		http.Error(w, ErrBannedUser.Error(), http.StatusUnauthorized)
		return
	}

	// get the photo from the resource parameter
	photo, code, err := rt.GetPhotoFromParameter("photo_id", UserFromDatabaseUser(dbUser), r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check if the resource is consistent
	if photo.User.Id != photoUser.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

	// get the like list from the database
	dbLikeList, err := rt.db.GetLikeList(photo.PhotoIntoDatabasePhoto(), dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	likeList := UserListFromDatabaseUserList(dbLikeList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the like list
	_ = json.NewEncoder(w).Encode(likeList)
}

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// authenticate the user performing the action
	likeUser, code, err := rt.AuthenticateUserFromParameter("like_uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// get the user of the photo from the resource parameter
	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// get the photo from the resource parameter
	photo, code, err := rt.GetPhotoFromParameter("photo_id", likeUser, r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check if the resource is consistent
	if photo.User.Id != user.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

	// insert the like into the databse
	err = rt.db.InsertLike(likeUser.UserIntoDatabaseUser(), photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbPhoto := photo.PhotoIntoDatabasePhoto()

	// update the number of likes to the photo
	err = rt.db.GetPhotoLikeCount(&dbPhoto, likeUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo.LikeCount = dbPhoto.LikeCount

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the photo that was liked
	_ = json.NewEncoder(w).Encode(photo)
}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// authenticate the user performing the action
	likeUser, code, err := rt.AuthenticateUserFromParameter("like_uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// get the user of the photo from the resource parameter
	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// get the photo from the resource parameter
	photo, code, err := rt.GetPhotoFromParameter("photo_id", likeUser, r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check if the resource is consistent
	if photo.User.Id != user.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

	// remove the like from the database
	err = rt.db.DeleteLike(likeUser.UserIntoDatabaseUser(), photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // 204
}
