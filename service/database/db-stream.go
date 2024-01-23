package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetDatabaseStream(dbUser DatabaseUser) (DatabaseStream, error) {
	dbStream := DatabaseStreamDefault()

	// get the user's stream table
	rows, err := db.c.Query(`
		SELECT id, user, url, date
		FROM Photo
		WHERE user IN (
			SELECT second_user
			FROM follow
			WHERE first_user=?
			  AND second_user NOT IN (
				SELECT first_user
				FROM ban
				WHERE second_user=?
			)
		)
		ORDER BY date DESC
	`, dbUser.Id, dbUser.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbStream, ErrUserDoesNotExist
	}

	if err != nil {
		return dbStream, err
	}

	dbPhotoUser := DatabaseUserDefault()

	// build the user's stream
	for rows.Next() {
		dbPhoto := DatabasePhotoDefault()

		err = rows.Scan(&dbPhoto.Id, &dbPhoto.User.Id, &dbPhoto.Url, &dbPhoto.Date)

		if err != nil {
			return dbStream, err
		}

		if dbPhotoUser.Id == 0 {
			dbPhotoUser, err = db.GetDatabaseUser(dbPhoto.User.Id)

			if err != nil {
				return dbStream, err
			}
		}

		dbPhoto.User = dbPhotoUser

		err = db.GetPhotoLikeCount(&dbPhoto, dbUser)

		if err != nil {
			return dbStream, err
		}

		err = db.GetPhotoCommentCount(&dbPhoto, dbUser)

		if err != nil {
			return dbStream, err
		}

		err = db.GetPhotoLikeStatus(&dbPhoto, dbUser)

		if err != nil {
			return dbStream, err
		}

		dbStream.Photos = append(dbStream.Photos, dbPhoto)
	}

	if rows.Err() != nil {
		return dbStream, err
	}

	_ = rows.Close()

	return dbStream, err
}
