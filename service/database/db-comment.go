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

func (db *appdbimpl) RemoveComment(user DatabaseUser, photo DatabasePhoto, comment DatabaseComment) error {
	res, err := db.c.Exec(`DELETE FROM Comment WHERE id=? AND user=? AND photo=? AND comment_body=?`, comment.Id, user.Id, photo.Id, comment.CommentBody)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return err
	} else if aff == 0 {
		return ErrUserNotFollowed
	}

	return nil
}