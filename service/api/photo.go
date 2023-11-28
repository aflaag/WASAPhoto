package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userUsername := ps.ByName("uname")
	userLogin := LoginFromUsername(userUsername)

	user, err := rt.GetUserFromLogin(userLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = CheckAuthorization(user, r.Header.Get("Authorization"))
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	photo := PhotoDefault()

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
	userUsername := ps.ByName("uname")
	userLogin := LoginFromUsername(userUsername)

	user, err := rt.GetUserFromLogin(userLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = CheckAuthorization(user, r.Header.Get("Authorization"))
	
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

	photo, err := rt.GetPhotoFromPhotoId(uint32(photoId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rt.db.DeletePhoto(photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(photo)
}