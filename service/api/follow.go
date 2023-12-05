package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// authenticate the user performing the action
	user, code, err := rt.AuthenticateUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// get the user to be followed from the resource parameter
	followedUser, code, err := rt.GetUserFromParameter("followed_uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// insert the following into the database
	err = rt.db.InsertFollow(user.UserIntoDatabaseUser(), followedUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the followed user
	_ = json.NewEncoder(w).Encode(followedUser)
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// authenticate the user performing the action
	user, code, err := rt.AuthenticateUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// get the followed user from the resource parameter
	followedUser, code, err := rt.GetUserFromParameter("followed_uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// remove the following from the database
	err = rt.db.DeleteFollow(user.UserIntoDatabaseUser(), followedUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // 204
}

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// get the user of the list from the resource parameter
	followersUser, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check whether the user of the list
	// has banned the user performing the action
	checkBan, err := rt.db.CheckBan(followersUser.UserIntoDatabaseUser(), dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if checkBan {
		http.Error(w, ErrBannedUser.Error(), http.StatusUnauthorized)
		return
	}

	// get the followers list from the database
	dbFollowersList, err := rt.db.GetFollowersList(followersUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followersList := UserListFromDatabaseUserList(dbFollowersList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the followers list
	_ = json.NewEncoder(w).Encode(followersList)
}

func (rt *_router) getFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// get the user of the list from the resource parameter
	followingUser, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check whether the user of the list
	// has banned the user performing the action
	checkBan, err := rt.db.CheckBan(followingUser.UserIntoDatabaseUser(), dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if checkBan {
		http.Error(w, ErrBannedUser.Error(), http.StatusUnauthorized)
		return
	}

	// get the following list from the database
	dbFollowingList, err := rt.db.GetFollowingList(followingUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followingList := UserListFromDatabaseUserList(dbFollowingList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the following list
	_ = json.NewEncoder(w).Encode(followingList)
}
