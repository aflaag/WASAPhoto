package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetDatabasePhoto(photoId uint32, dbUser DatabaseUser) (DatabasePhoto, error) {
	dbPhoto := DatabasePhotoDefault()

	err := db.c.QueryRow(`
		SELECT id, user, date, url
		FROM Photo
		WHERE id=?
	`, photoId).Scan(&dbPhoto.Id, &dbPhoto.User.Id, &dbPhoto.Date, &dbPhoto.Url)

	if errors.Is(err, sql.ErrNoRows) {
		return dbPhoto, ErrPhotoDoesNotExist
	}

	// get the user information
	dbPhotoUser, err := db.GetDatabaseUser(dbPhoto.User.Id)

	if err != nil {
		return dbPhoto, err
	}

	dbPhoto.User.Username = dbPhotoUser.Username

	// get the like count
	err = db.GetPhotoLikeCount(&dbPhoto, dbUser)

	if err != nil {
		return dbPhoto, err
	}

	// get the comment count
	err = db.GetPhotoCommentCount(&dbPhoto, dbUser)

	if err != nil {
		return dbPhoto, err
	}

	// get the like status
	err = db.GetPhotoLikeStatus(&dbPhoto, dbUser)

	return dbPhoto, err
}

func (db *appdbimpl) GetPhotoLikeStatus(dbPhoto *DatabasePhoto, dbUser DatabaseUser) error {
	// check whether the first user has banned the second user
	err := db.c.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM like
			WHERE user=?
			AND photo=?
		)
	`, dbUser.Id, dbPhoto.Id).Scan(&dbPhoto.LikeStatus)

	// if no table rows are found, then there is no row
	// containing the like, hence the user has not liked the photo
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	return err
}

func (db *appdbimpl) InsertPhoto(dbPhoto *DatabasePhoto) error {
	// insert the photo into the database
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

func (db *appdbimpl) GetPhotoLikeCount(dbPhoto *DatabasePhoto, dbUser DatabaseUser) error {
	// return the number of likes to the photo
	// without counting the likes of users who banned
	// the user performing the action
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

func (db *appdbimpl) GetPhotoCommentCount(dbPhoto *DatabasePhoto, dbUser DatabaseUser) error {
	// return the number of likes to the photo
	// without counting the likes of users who banned
	// the user performing the action
	err := db.c.QueryRow(`
		SELECT COUNT(*)
		FROM Comment
		WHERE photo=?
		AND user NOT IN (
			SELECT first_user
			FROM ban
			WHERE second_user=?
		)`, dbPhoto.Id, dbUser.Id).Scan(&dbPhoto.CommentCount)

	if errors.Is(err, sql.ErrNoRows) {
		return ErrPhotoDoesNotExist
	}

	return err
}

func (db *appdbimpl) GetPhotos(dbProfile *DatabaseProfile, dbUser DatabaseUser) error {
	rows, err := db.c.Query(`
		SELECT id
		FROM photo
		WHERE user=?
	`, dbProfile.User.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserDoesNotExist
		}

		return err
	}

	// build the results list
	for rows.Next() {
		newDbPhoto := DatabasePhotoDefault()

		err = rows.Scan(&newDbPhoto.Id)

		if err != nil {
			return err
		}

		newDbPhoto, err = db.GetDatabasePhoto(newDbPhoto.Id, dbUser)

		if err != nil {
			return err
		}

		dbProfile.Photos = append(dbProfile.Photos, newDbPhoto)
	}

	if rows.Err() != nil {
		return err
	}

	_ = rows.Close()

	return err
}

func (db *appdbimpl) GetPhotoCount(dbUser DatabaseUser) (int, error) {
	var photoCount int

	// get the number of photos the user has posted
	err := db.c.QueryRow(`
		SELECT COUNT(*)
		FROM Photo
		WHERE user=?
	`, dbUser.Id).Scan(&photoCount)

	if errors.Is(err, sql.ErrNoRows) {
		return photoCount, ErrPhotoDoesNotExist
	}

	return photoCount, err
}
