package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetDatabaseComment(commentId uint32) (DatabaseComment, error) {
	dbComment := DatabaseCommentDefault()

	err := db.c.QueryRow(`
		SELECT id, user, date, comment_body
		FROM Comment
		WHERE id=?
	`, commentId).Scan(&dbComment.Id, &dbComment.User.Id, &dbComment.Date, &dbComment.CommentBody)

	if errors.Is(err, sql.ErrNoRows) {
		return dbComment, ErrCommentDoesNotExist
	}

	dbUser, err := db.GetDatabaseUser(dbComment.User.Id)

	if err != nil {
		return dbComment, err
	}

	dbComment.User.Username = dbUser.Username

	return dbComment, err
}

func (db *appdbimpl) InsertComment(dbComment *DatabaseComment) error {
	res, err := db.c.Exec(`
		INSERT INTO Comment(user, photo, date, comment_body)
		VALUES (?, ?, ?, ?)
	`, dbComment.User.Id, dbComment.Photo.Id, dbComment.Date, dbComment.CommentBody)

	if err != nil {
		return err
	}

	dbCommentId, err := res.LastInsertId()

	if err != nil {
		return err
	}

	dbComment.Id = uint32(dbCommentId)

	return nil
}

func (db *appdbimpl) DeleteComment(dbComment DatabaseComment) error {
	res, err := db.c.Exec(`
		DELETE FROM Comment
		WHERE id=?
	`, dbComment.Id)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if aff == 0 {
		return ErrPhotoNotCommented
	}

	return err
}

func (db *appdbimpl) GetCommentList(dbUser DatabaseUser, dbPhoto DatabasePhoto) (DatabaseCommentList, error) {
	dbCommentList := DatabaseCommentListDefault()

	rows, err := db.c.Query(`
		SELECT id, user, photo, date, comment_body
		FROM Comment
		WHERE photo=?
		AND user NOT IN (
			SELECT first_user
			FROM ban
			WHERE second_user=?
		)
	`, dbPhoto.Id, dbUser.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbCommentList, ErrPhotoDoesNotExist
	}

	if err != nil {
		return dbCommentList, err
	}

	for rows.Next() {
		dbComment := DatabaseCommentDefault()

		// TODO: MANCANO DA FILLARE UTENTI E FOTO
		err = rows.Scan(&dbComment.Id, &dbComment.User.Id, &dbComment.Photo.Id, &dbComment.Date, &dbComment.CommentBody)

		if err != nil {
			return dbCommentList, err
		}

		dbCommentList.Comments = append(dbCommentList.Comments, dbComment)
	}

	if rows.Err() != nil {
		return dbCommentList, err
	}

	_ = rows.Close()

	return dbCommentList, err
}
