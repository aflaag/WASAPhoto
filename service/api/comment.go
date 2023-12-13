package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// get the user of the photo from the resource parameter
	photoUser, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check whether the user of the photo
	// has banned the user performing the action
	checkBan, err := rt.db.CheckBan(photoUser.UserIntoDatabaseUser(), dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if checkBan {
		http.Error(w, ErrBannedUser.Error(), http.StatusUnauthorized)
		return
	}

	// get the photo from the resource parameter
	photo, code, err := rt.GetPhotoFromParameter("photo_id", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check if the resource is consistent
	if photo.User.Id != photoUser.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

	// get the comment list from the database
	dbCommentList, err := rt.db.GetCommentList(photo.PhotoIntoDatabasePhoto(), dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commentList := CommentListFromDatabaseCommentList(dbCommentList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the comment list
	_ = json.NewEncoder(w).Encode(commentList)
}

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	comment := CommentDefault()

	// get the comment information from the request body
	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commentLogin := LoginDefault()

	commentLogin.Username = comment.User.Username

	// get the user writing the comment from the database
	commentUser, err := rt.GetUserFromLogin(commentLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check whether the user id specified
	// in the request body matches the real user id
	if comment.User.Id != commentUser.Id {
		http.Error(w, ErrUserDoesNotExist.Error(), http.StatusUnauthorized)
		return
	}

	// check if the bearer token matches the
	// user specified in the request body
	err = CheckAuthorization(comment.User, r.Header.Get("Authorization"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// get the user of the photo from the resource parameter
	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// get the photo from the resource parameter
	photo, code, err := rt.GetPhotoFromParameter("photo_id", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check if the resource is consistent
	if photo.User.Id != user.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

	comment.Photo = photo

	comment.Date = time.Now().Format("2006-01-02 15:04:05")

	dbComment := comment.CommentIntoDatabaseComment()

	// insert the comment into the database
	err = rt.db.InsertComment(&dbComment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the comment id from the database
	comment.Id = dbComment.Id

	dbPhoto := photo.PhotoIntoDatabasePhoto()

	// update the number of comments under the photo
	err = rt.db.GetPhotoCommentCount(&dbPhoto, user.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment.Photo = PhotoFromDatabasePhoto(dbPhoto)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	// return the newly created comment
	_ = json.NewEncoder(w).Encode(comment)
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get the comment id from the resource parameter
	commentIdString := ps.ByName("comment_id")
	commentId, err := strconv.ParseUint(commentIdString, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the comment from the database
	comment, err := rt.GetCommentFromCommentId(uint32(commentId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if the user in the bearer token
	// matches the comment user
	err = CheckAuthorization(comment.User, r.Header.Get("Authorization"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// get the user of the photo from the resource parameter
	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// get the photo from the resource parameter
	photo, code, err := rt.GetPhotoFromParameter("photo_id", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// check if the resource is consistent
	if photo.User.Id != user.Id || photo.Id != comment.Photo.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

	// remove the comment from the database
	err = rt.db.DeleteComment(comment.CommentIntoDatabaseComment())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbPhoto := photo.PhotoIntoDatabasePhoto()

	// update the number of comments under the photo
	err = rt.db.GetPhotoCommentCount(&dbPhoto, user.UserIntoDatabaseUser())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment.Photo = PhotoFromDatabasePhoto(dbPhoto)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the removed comment
	_ = json.NewEncoder(w).Encode(comment)
}
