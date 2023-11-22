package database

import (
	// "database/sql"
)

func (db *appdbimpl) SetComment(user DatabaseUser, photo DatabasePhoto, comment DatabaseComment) error {
	_, err := db.c.Exec(`INSERT INTO Comment (id, user, photo, comment_body) VALUES (?, ?, ?, ?)`, comment.Id, user.Id, photo.Id, comment.CommentBody)

	if err != nil {
		return err
	}
	
	return nil
}