package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token, err := GetBearerToken(r.Header.Get("Authorization"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dbUser, err := rt.db.GetDatabaseUser(uint32(token))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profileUser, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	checkBan, err := rt.db.CheckBan(profileUser.UserIntoDatabaseUser(), dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if checkBan {
		http.Error(w, ErrBannedUser.Error(), http.StatusUnauthorized)
		return
	}

	profile := ProfileDefault()

	profile.User = profileUser

	profile.PhotosCount, err = rt.db.GetPhotosCount(profileUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile.FollowersCount, err = rt.db.GetFollowersCount(profileUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile.FollowingCount, err = rt.db.GetFollowingCount(profileUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(profile)
}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	oldUser, code, err := rt.AuthenticateUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	newUserLogin := LoginDefault()

	err = json.NewDecoder(r.Body).Decode(&newUserLogin)

	newUser := UserDefault()

	newUser.Id = oldUser.Id
	newUser.Username = newUserLogin.Username

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rt.db.UpdateUser(oldUser.UserIntoDatabaseUser(), newUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(newUser)
}
