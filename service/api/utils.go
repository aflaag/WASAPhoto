package api

import (
	// "net/http"
	// "strconv"
	"errors"

	// "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	// "github.com/julienschmidt/httprouter"
)

func (rt *_router) GetUserFromUsername(userUsername string) (User, error) {
	dbUser, err := rt.db.GetUserFromUsername(userUsername)

	if err != nil {
		return UserDefault(), errors.New("TODO")
	}

	user := UserFromDatabaseUser(dbUser)

	return user, nil
}