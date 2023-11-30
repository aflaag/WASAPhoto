package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetUserStream(dbUser DatabaseUser) (DatabaseStream, error) {
	dbStream := DatabaseStreamDefault()

	err := db.c.QueryRow(`SELECT id, user, date, url FROM Photo WHERE id=?`, dbUser.Id).Scan(&dbStream.Photos)

	if errors.Is(err, sql.ErrNoRows) {
		return dbStream, ErrUserDoesNotExist
	}

	return dbStream, err
}

