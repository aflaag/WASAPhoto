package database

import (
	"database/sql"
)

func (db *appdbimpl) GetDatabaseUser(dbLogin DatabaseLogin) (DatabaseUser, error) {
	var dbUser DatabaseUser

	err := db.c.QueryRow(`SELECT id, username from USER where username=?`, dbLogin.Username).Scan(&dbUser.Id, &dbUser.Username)

	if err == sql.ErrNoRows {
		return dbUser, ErrUserDoesNotExist
	}

	return dbUser, err
}

func (db *appdbimpl) CreateDatabaseUser(dbLogin DatabaseLogin) (DatabaseUser, error) {
	res, err := db.c.Exec("INSERT INTO User(username) VALUES (?)", dbLogin.Username)

	dbUser := DatabaseUserDefault()

	if err != nil {
		err := db.c.QueryRow(`SELECT id, username FROM User WHERE username=?`, dbLogin.Username).Scan(&dbUser.Id, &dbUser.Username)

		if err == sql.ErrNoRows {
			return dbUser, ErrUserDoesNotExist
		}

		return dbUser, err
	}

	dbUserId, err := res.LastInsertId()

	if err != nil {
		return dbUser, err
	}

	dbUser.Id = uint64(dbUserId)
	dbUser.Username = dbLogin.Username

	return dbUser, nil
}