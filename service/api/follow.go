package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userUsername := ps.ByName("uname")
	userLogin := LoginFromUsername(userUsername)

	followedUserUsername := ps.ByName("follow_uname")
	followedUserLogin := LoginFromUsername(followedUserUsername)

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

	followedUser, err := rt.GetUserFromLogin(followedUserLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rt.db.InsertFollow(user.UserIntoDatabaseUser(), followedUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(followedUser)
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userUsername := ps.ByName("uname")
	userLogin := LoginFromUsername(userUsername)

	followedUserUsername := ps.ByName("follow_uname")
	followedUserLogin := LoginFromUsername(followedUserUsername)

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

	followedUser, err := rt.GetUserFromLogin(followedUserLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rt.db.DeleteFollow(user.UserIntoDatabaseUser(), followedUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) getFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}