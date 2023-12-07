package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) InsertLike(dbUser DatabaseUser, dbPhoto DatabasePhoto) error {
	// insert the like into the database
	_, err := db.c.Exec(`
		INSERT OR IGNORE INTO like(user, photo)
		VALUES (?, ?)
	`, dbUser.Id, dbPhoto.Id)

	return err
}

func (db *appdbimpl) DeleteLike(dbUser DatabaseUser, dbPhoto DatabasePhoto) error {
	res, err := db.c.Exec(`
		DELETE FROM like
		WHERE user=?
		AND photo=?
	`, dbUser.Id, dbPhoto.Id)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return err
	}

	// if there are no affected rows
	// then the photo was not liked
	if aff == 0 {
		return ErrPhotoNotLiked
	}

	return err
}

func (db *appdbimpl) GetLikeList(dbPhoto DatabasePhoto, dbUser DatabaseUser) (DatabaseUserList, error) {
	dbUserList := DatabaseUserListDefault()

	// get the table of the users who liked the photo
	// without the users who banned the user performing the action
	rows, err := db.c.Query(`
		SELECT id, username
		FROM User
		WHERE id IN (
			SELECT user
			FROM like
			WHERE photo=?
		)
		AND id NOT IN (
			SELECT first_user
			FROM ban
			WHERE second_user=?
		)
	`, dbPhoto.Id, dbUser.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUserList, ErrPhotoDoesNotExist
	}

	if err != nil {
		return dbUserList, err
	}

	// build the like list
	for rows.Next() {
		tableDbUser := DatabaseUserDefault()

		err = rows.Scan(&tableDbUser.Id, &tableDbUser.Username)

		if err != nil {
			return dbUserList, err
		}

		dbUserList.Users = append(dbUserList.Users, tableDbUser)
	}

	if rows.Err() != nil {
		return dbUserList, err
	}

	_ = rows.Close()

	return dbUserList, err
}
