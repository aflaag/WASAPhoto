package api

import (
	"net/http"
	// "strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userUsername := ps.ByName("uid")

	user, err := rt.GetUserFromUsername(userUsername)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followedUserUsername := ps.ByName("followuid")

	followedUser, err := rt.GetUserFromUsername(followedUserUsername)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rt.db.SetFollow(user.UserIntoDatabaseUser(), followedUser.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}