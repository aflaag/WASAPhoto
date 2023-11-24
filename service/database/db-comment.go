package database

import (
	// "database/sql"
)

func (db *appdbimpl) SetComment(dbUser DatabaseUser, dbPhoto DatabasePhoto, dbComment DatabaseComment) error {
	_, err := db.c.Exec(`INSERT INTO Comment (id, user, photo, comment_body) VALUES (?, ?, ?, ?)`, dbComment.Id, dbUser.Id, dbPhoto.Id, dbComment.CommentBody)

	return err
}

func (db *appdbimpl) RemoveComment(dbUser DatabaseUser, dbPhoto DatabasePhoto, dbComment DatabaseComment) error {
	res, err := db.c.Exec(`DELETE FROM Comment WHERE id=? AND user=? AND photo=? AND comment_body=?`, dbComment.Id, dbUser.Id, dbPhoto.Id, dbComment.CommentBody)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if aff == 0 {
		return ErrUserNotFollowed
	}

	return err
}