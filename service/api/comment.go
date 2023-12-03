package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	photo, code, err := rt.GetPhotoFromParameter("photo_id", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	if photo.User.Id != user.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

	dbCommentList, err := rt.db.GetCommentList(photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commentList := CommentListFromDatabaseCommentList(dbCommentList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(commentList)
}

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	comment := CommentDefault()

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commentLogin := LoginDefault()
	commentLogin.Username = comment.User.Username

	commentUser, err := rt.GetUserFromLogin(commentLogin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if comment.User.Id != commentUser.Id {
		http.Error(w, ErrUserDoesNotExist.Error(), http.StatusUnauthorized)
		return
	}

	err = CheckAuthorization(comment.User, r.Header.Get("Authorization"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	photo, code, err := rt.GetPhotoFromParameter("photo_id", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	if photo.User.Id != user.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

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

	user, code, err := rt.GetUserFromParameter("uname", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	photo, code, err := rt.GetPhotoFromParameter("photo_id", r, ps)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	if photo.User.Id != user.Id {
		http.Error(w, ErrPageNotFound.Error(), http.StatusNotFound)
		return
	}

	err = rt.db.RemoveComment(comment.CommentIntoDatabaseComment(), photo.PhotoIntoDatabasePhoto())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(comment)
}
