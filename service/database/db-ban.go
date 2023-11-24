package database

import (
	// "database/sql"
)

func (db *appdbimpl) SetBan(dbUser DatabaseUser, bannedDbUser DatabaseUser) error {
	_, err := db.c.Exec(`INSERT INTO ban (first_user, second_user) VALUES (?, ?)`, dbUser.Id, bannedDbUser.Id)

	return err
}