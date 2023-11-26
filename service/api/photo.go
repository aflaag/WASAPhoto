package api

import (
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userUsername := ps.ByName("uid")
	photoIdString := ps.ByName("photoid")

	userLogin := LoginFromUsername(userUsername)

	user, err := rt.GetUserFromLogin(userLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photoId, err := strconv.ParseUint(photoIdString, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo := PhotoDefault()
	photo.Id = photoId

	err = rt.db.SetPhoto(user.UserIntoDatabaseUser(), photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userUsername := ps.ByName("uid")
	photoIdString := ps.ByName("photoid")

	userLogin := LoginFromUsername(userUsername)

	_, err := rt.GetUserFromLogin(userLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photoId, err := strconv.ParseUint(photoIdString, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo := PhotoDefault() // TODO: VANNO RIEMPITI GLI ALTRI CAMPI

	photo.Id = photoId

	// TODO: CONTROLLARE CHE L'UTENTE SIA LO STESSO

	err = rt.db.RemovePhoto(photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}