package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetDatabaseStream(dbUser DatabaseUser) (DatabaseStream, error) {
	dbStream := DatabaseStreamDefault()

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
	`, dbUser.Id, dbUser.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbStream, ErrUserDoesNotExist
	}

	if err != nil {
		return dbStream, err
	}

	for rows.Next() {
		dbPhoto := DatabasePhotoDefault()

		err = rows.Scan(&dbPhoto.Id, &dbPhoto.User.Id, &dbPhoto.Url, &dbPhoto.Date)

		if err != nil {
			return dbStream, err
		}

		dbUser, err := db.GetDatabaseUser(dbPhoto.User.Id)

		if err != nil {
			return dbStream, err
		}

		dbPhoto.User = dbUser

		err = db.GetPhotoLikeCount(&dbPhoto)

		if err != nil {
			return dbStream, err
		}

		err = db.GetPhotoCommentCount(&dbPhoto)

		if err != nil {
			return dbStream, err
		}

		dbStream.Photos = append(dbStream.Photos, dbPhoto)
	}

	_ = rows.Close()

	return dbStream, err
}
