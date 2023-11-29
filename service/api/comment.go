package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	comment := CommentDefault()

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	login := LoginDefault()
	login.Username = comment.User.Username

	user, err := rt.GetUserFromLogin(login)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if user.Id != comment.User.Id {
		http.Error(w, ErrUserDoesNotExist.Error(), http.StatusUnauthorized)
		return
	}

	err = CheckAuthorization(comment.User, r.Header.Get("Authorization"))
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	photoIdString := ps.ByName("photo_id")

	photoId, err := strconv.ParseUint(photoIdString, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	photo := PhotoDefault()
	photo.Id = uint32(photoId)

	dbComment := comment.CommentIntoDatabaseComment()

	err = rt.db.InsertComment(&dbComment, photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment.Id = dbComment.Id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(comment)
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	commentIdString := ps.ByName("comment_id")

	commentId, err := strconv.ParseUint(commentIdString, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment, err := rt.GetCommentFromCommentId(uint32(commentId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = CheckAuthorization(comment.User, r.Header.Get("Authorization"))
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	photoIdString := ps.ByName("photo_id")

	photoId, err := strconv.ParseUint(photoIdString, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo := PhotoDefault()
	photo.Id = uint32(photoId)

	err = rt.db.RemoveComment(comment.CommentIntoDatabaseComment(), photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(comment)
}