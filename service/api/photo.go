package api

import (
	"encoding/json"
	"net/http"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	user, code, err := rt.AuthenticateUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	photo := PhotoDefault() // TODO: IN TEORIA DEVI PRENDERLO DAL REQUEST BODY PER PRENDERE ANCHE L'URL

	currentTime := time.Now()

	photo.User = user
	photo.Date = currentTime.Format("2006-01-02 15:04:05")

	// TODO: MANCA DA CAPIRE COSA SIA L'URL

	dbPhoto := photo.PhotoIntoDatabasePhoto()

	err = rt.db.InsertPhoto(&dbPhoto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo.Id = dbPhoto.Id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(photo)
}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	_, code, err := rt.AuthenticateUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	photo, code, err := rt.GetPhotoFromParameter("photo_id", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	err = rt.db.DeletePhoto(photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(photo)
}
