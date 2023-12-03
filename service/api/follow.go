package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	user, code, err := rt.AuthenticateUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	followedUser, code, err := rt.GetUserFromParameter("followed_uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
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
	user, code, err := rt.AuthenticateUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	followedUser, code, err := rt.GetUserFromParameter("followed_uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
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
	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	dbFollowersList, err := rt.db.GetFollowersList(user.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followersList := UserListFromDatabaseUserList(dbFollowersList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(followersList)
}

func (rt *_router) getFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	dbFollowersList, err := rt.db.GetFollowingList(user.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followersList := UserListFromDatabaseUserList(dbFollowersList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(followersList)
}
