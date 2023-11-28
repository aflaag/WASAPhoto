package database

import (
	"database/sql"
)

func (db *appdbimpl) GetDatabaseUser(userId uint64) (DatabaseUser, error) {
	dbUser := DatabaseUserDefault()

	err := db.c.QueryRow(`SELECT id, username FROM User WHERE id=?`, userId).Scan(&dbUser.Id, &dbUser.Username)

	if err == sql.ErrNoRows {
		return dbUser, ErrUserDoesNotExist
	}

	return dbUser, err
}

func (db *appdbimpl) GetDatabaseUserFromDatabaseLogin(dbLogin DatabaseLogin) (DatabaseUser, error) {
	dbUser := DatabaseUserDefault()

	err := db.c.QueryRow(`SELECT id, username FROM User WHERE username=?`, dbLogin.Username).Scan(&dbUser.Id, &dbUser.Username)

	if err == sql.ErrNoRows {
		return dbUser, ErrUserDoesNotExist
	}

	return dbUser, err
}

func (db *appdbimpl) InsertUser(dbUser *DatabaseUser) error {
	res, err := db.c.Exec("INSERT INTO User(username) VALUES (?)", dbUser.Username)

	if err != nil {
		return err
	}

	dbUserId, err := res.LastInsertId()

	if err != nil {
		return err
	}

	dbUser.Id = uint64(dbUserId)

	return nil
}