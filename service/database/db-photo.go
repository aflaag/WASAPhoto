package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetDatabasePhoto(photoId uint32) (DatabasePhoto, error) {
	dbPhoto := DatabasePhotoDefault()

	err := db.c.QueryRow(`SELECT id, user, date, url FROM Photo WHERE id=?`, photoId).Scan(&dbPhoto.Id, &dbPhoto.User.Id, &dbPhoto.Date, &dbPhoto.Url)

	if errors.Is(err, sql.ErrNoRows) {
		return dbPhoto, ErrPhotoDoesNotExist
	}

	dbUser, err := db.GetDatabaseUser(dbPhoto.User.Id)

	if err != nil {
		return dbPhoto, err
	}

	dbPhoto.User.Username = dbUser.Username

	return dbPhoto, err
}

func (db *appdbimpl) InsertPhoto(dbPhoto *DatabasePhoto) error {
	res, err := db.c.Exec(`
		INSERT INTO Photo(user, url, date)
		VALUES (?, ?, ?)
	`, dbPhoto.User.Id, dbPhoto.Url, dbPhoto.Date)

	if err != nil {
		return err
	}

	// get the photo id
	dbPhotoId, err := res.LastInsertId()

	if err != nil {
		return err
	}

	dbPhoto.Id = uint32(dbPhotoId)

	return nil
}

func (db *appdbimpl) DeletePhoto(dbPhoto DatabasePhoto) error {
	// remove every like to the photo from the database
	_, err := db.c.Exec(`
		DELETE FROM like
		WHERE photo=?
	`, dbPhoto.Id)

	if err != nil {
		return err
	}

	// remove every comment under the photo from the database
	_, err = db.c.Exec(`
		DELETE FROM Comment
		WHERE photo=?
	`, dbPhoto.Id)

	if err != nil {
		return err
	}

	// remove the photo from the database
	_, err = db.c.Exec(`
		DELETE FROM Photo
		WHERE id=?
	`, dbPhoto.Id)

	return err
}

func (db *appdbimpl) GetPhotoLikeCount(dbUser DatabaseUser, dbPhoto *DatabasePhoto) error {
	err := db.c.QueryRow(`
		SELECT COUNT(*)
		FROM like
		WHERE photo=?
		AND user NOT IN (
			SELECT first_user
			FROM ban
			WHERE second_user=?
		)
	`, dbPhoto.Id, dbUser.Id).Scan(&dbPhoto.LikeCount)

	if errors.Is(err, sql.ErrNoRows) {
		return ErrPhotoDoesNotExist
	}

	return err
}

func (db *appdbimpl) GetPhotoCommentCount(dbUser DatabaseUser, dbPhoto *DatabasePhoto) error {
	err := db.c.QueryRow(`
		SELECT COUNT(*)
		FROM Comment
		WHERE photo=?
		AND user NOT IN (
			SELECT first_user
			FROM ban
			WHERE second_user?
		)`, dbPhoto.Id, dbUser.Id).Scan(&dbPhoto.CommentCount)

	if errors.Is(err, sql.ErrNoRows) {
		return ErrPhotoDoesNotExist
	}

	return err
}

func (db *appdbimpl) GetPhotoCount(dbUser DatabaseUser) (int, error) {
	var photoCount int

	err := db.c.QueryRow(`SELECT COUNT(*) FROM Photo WHERE user=?`, dbUser.Id).Scan(&photoCount)

	if errors.Is(err, sql.ErrNoRows) {
		return photoCount, ErrPhotoDoesNotExist
	}

	return photoCount, err
}
