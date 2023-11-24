package database

import (
	"database/sql"
)

func (db *appdbimpl) GetDatabaseUserFromUsername(userUsername string) (DatabaseUser, error) {
	var user DatabaseUser

	err := db.c.QueryRow(`SELECT id, username from USER where username=?`, userUsername).Scan(&user.Id, &user.Username)

	if err == sql.ErrNoRows {
		return user, ErrUserDoesNotExist
	}

	return user, err
}