package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) InsertLike(dbUser DatabaseUser, dbPhoto DatabasePhoto) error {
	_, err := db.c.Exec(`
		INSERT OR IGNORE INTO like(user, photo)
		VALUES (?, ?)
	`, dbUser.Id, dbPhoto.Id)

	return err
}

func (db *appdbimpl) DeleteLike(dbUser DatabaseUser, dbPhoto DatabasePhoto) error {
	res, err := db.c.Exec(`DELETE FROM like WHERE user=? AND photo=?`, dbUser.Id, dbPhoto.Id)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if aff == 0 {
		return ErrPhotoNotLiked
	}

	return err
}

func (db *appdbimpl) GetLikeList(dbPhoto DatabasePhoto) (DatabaseUserList, error) {
	dbUserList := DatabaseUserListDefault()

	rows, err := db.c.Query(`
		SELECT id, username
		FROM User
		WHERE id IN (
			SELECT id
			FROM like
			WHERE photo=?
		)
	`, dbPhoto.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUserList, ErrPhotoDoesNotExist
	}

	if err != nil {
		return dbUserList, err
	}

	for rows.Next() {
		dbUser := DatabaseUserDefault()

		err = rows.Scan(&dbUser.Id, &dbUser.Username)

		if err != nil {
			return dbUserList, err
		}

		dbUserList.Users = append(dbUserList.Users, dbUser)
	}

	_ = rows.Close()

	return dbUserList, err
}
