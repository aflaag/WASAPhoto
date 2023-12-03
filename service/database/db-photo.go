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
	res, err := db.c.Exec(`INSERT INTO Photo(user, url, date) VALUES (?, ?, ?)`, dbPhoto.User.Id, dbPhoto.Url, dbPhoto.Date)

	if err != nil {
		return err
	}

	dbPhotoId, err := res.LastInsertId()

	if err != nil {
		return err
	}

	dbPhoto.Id = uint32(dbPhotoId)

	return nil
}

func (db *appdbimpl) DeletePhoto(dbPhoto DatabasePhoto) error {
	var err error

	_, err = db.c.Exec(`DELETE FROM like WHERE photo=?`, dbPhoto.Id)

	if err != nil {
		return err
	}

	_, err = db.c.Exec(`DELETE FROM comment WHERE photo=?`, dbPhoto.Id)

	if err != nil {
		return err
	}

	_, err = db.c.Exec(`DELETE FROM Photo WHERE id=?`, dbPhoto.Id)

	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) GetPhotoLikeCount(dbPhoto *DatabasePhoto) error {
	err := db.c.QueryRow(`SELECT COUNT(*) FROM like WHERE photo=?`, dbPhoto.Id).Scan(&dbPhoto.LikeCount)

	if errors.Is(err, sql.ErrNoRows) {
		return ErrPhotoDoesNotExist
	}

	return err
}

func (db *appdbimpl) GetPhotoCommentCount(dbPhoto *DatabasePhoto) error {
	err := db.c.QueryRow(`SELECT COUNT(*) FROM Comment WHERE photo=?`, dbPhoto.Id).Scan(&dbPhoto.CommentCount)

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
