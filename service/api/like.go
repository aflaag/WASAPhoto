package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	likeUser, code, err := rt.AuthenticateUserFromParameter("like_uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	photo, code, err := rt.GetPhotoFromParameter("photo_id", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	if photo.User.Id != user.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

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
	likeUser, code, err := rt.AuthenticateUserFromParameter("like_uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	photo, code, err := rt.GetPhotoFromParameter("photo_id", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	if photo.User.Id != user.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

	err = rt.db.DeleteLike(likeUser.UserIntoDatabaseUser(), photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
