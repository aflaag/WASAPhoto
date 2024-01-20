package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get the user performing the action from the resource parameter
	user, code, err := rt.AuthenticateUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	dbUser := user.UserIntoDatabaseUser()

	// get the stream of the user performing the action
	dbStream, err := rt.db.GetDatabaseStream(dbUser)

	dbStream.User = dbUser

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stream := StreamFromDatabaseStream(dbStream)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the user's stream
	_ = json.NewEncoder(w).Encode(stream)
}
