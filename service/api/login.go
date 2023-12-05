package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) session(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	login := LoginDefault()

	// get the new user's username
	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create a new user
	user := UserDefault()
	dbUser := user.UserIntoDatabaseUser()

	// update the new user's username
	dbUser.Username = login.Username

	// insert the new user into the database
	err = rt.db.InsertUser(&dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the user id from the database
	user = UserFromDatabaseUser(dbUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	// return the newly created user
	_ = json.NewEncoder(w).Encode(user)
}
