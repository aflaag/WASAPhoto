package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	likeUserUsername := ps.ByName("like_uname")
	likeUserLogin := LoginFromUsername(likeUserUsername)

	likeUser, err := rt.GetUserFromLogin(likeUserLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = CheckAuthorization(likeUser, r.Header.Get("Authorization"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	photoIdString := ps.ByName("photo_id")

	photoId, err := strconv.ParseUint(photoIdString, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	photo := PhotoDefault()
	photo.Id = uint32(photoId)

	err = rt.db.InsertLike(likeUser.UserIntoDatabaseUser(), photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(likeUser)
}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	likeUserUsername := ps.ByName("like_uname")
	likeUserLogin := LoginFromUsername(likeUserUsername)

	likeUser, err := rt.GetUserFromLogin(likeUserLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = CheckAuthorization(likeUser, r.Header.Get("Authorization"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	photoIdString := ps.ByName("photo_id")

	photoId, err := strconv.ParseUint(photoIdString, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	photo := PhotoDefault()
	photo.Id = uint32(photoId)

	err = rt.db.DeleteLike(likeUser.UserIntoDatabaseUser(), photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
