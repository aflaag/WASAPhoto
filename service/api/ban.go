package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getBanList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {	
}

// TODO: CAMBIA I NOMI ALLE VARIABILI
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userUsername := ps.ByName("uname")
	userLogin := LoginFromUsername(userUsername)

	bannedUserUsername := ps.ByName("banned_uname")
	bannedUserLogin := LoginFromUsername(bannedUserUsername)

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

	bannedUser, err := rt.GetUserFromLogin(bannedUserLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rt.db.InsertBan(user.UserIntoDatabaseUser(), bannedUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(bannedUser)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userUsername := ps.ByName("uname")
	userLogin := LoginFromUsername(userUsername)

	bannedUserUsername := ps.ByName("banned_uname")
	bannedUserLogin := LoginFromUsername(bannedUserUsername)

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

	bannedUser, err := rt.GetUserFromLogin(bannedUserLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rt.db.DeleteBan(user.UserIntoDatabaseUser(), bannedUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}