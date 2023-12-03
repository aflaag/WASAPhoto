package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetDatabaseComment(commentId uint32) (DatabaseComment, error) {
	dbComment := DatabaseCommentDefault()

	err := db.c.QueryRow(`SELECT id, user, comment_body FROM Comment WHERE id=?`, commentId).Scan(&dbComment.Id, &dbComment.User.Id, &dbComment.CommentBody)

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

func (db *appdbimpl) InsertComment(dbComment *DatabaseComment, dbPhoto DatabasePhoto) error {
	res, err := db.c.Exec(`INSERT INTO Comment(user, photo, comment_body) VALUES (?, ?, ?)`, dbComment.User.Id, dbPhoto.Id, dbComment.CommentBody)

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

func (db *appdbimpl) RemoveComment(dbComment DatabaseComment, dbPhoto DatabasePhoto) error {
	res, err := db.c.Exec(`DELETE FROM Comment WHERE id=? AND user=? AND photo=? AND comment_body=?`, dbComment.Id, dbComment.User.Id, dbPhoto.Id, dbComment.CommentBody)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if aff == 0 {
		return ErrPhotoNotCommented
	}

	return err
}

func (db *appdbimpl) GetCommentList(dbPhoto DatabasePhoto) (DatabaseCommentList, error) {
	dbCommentList := DatabaseCommentListDefault()

	rows, err := db.c.Query(`
		SELECT id, user, comment_body, photo
		FROM Comment
		WHERE photo=?
	`, dbPhoto.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbCommentList, ErrUserDoesNotExist
	}

	if err != nil {
		return dbCommentList, err
	}

	for rows.Next() {
		dbComment := DatabaseCommentDefault()

		var photoId int

		err = rows.Scan(&dbComment.Id, &dbComment.User.Id, &dbComment.CommentBody, &photoId)

		if err != nil {
			return dbCommentList, err
		}

		dbCommentList.Comments = append(dbCommentList.Comments, dbComment)
	}

	_ = rows.Close()

	return dbCommentList, err
}
