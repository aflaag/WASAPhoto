package database

import (
	// "database/sql"
)

func (db *appdbimpl) SetBan(user DatabaseUser, bannedUser DatabaseUser) error {
	_, err := db.c.Exec(`INSERT INTO ban (first_user, second_user) VALUES (?, ?)`, user.Id, bannedUser.Id)

	if err != nil {
		return err
	}
	
	return nil
}