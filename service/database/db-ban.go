package database

import (
	// "database/sql"
)

func (db *appdbimpl) InsertBan(dbUser DatabaseUser, bannedDbUser DatabaseUser) error {
	_, err := db.c.Exec(`INSERT INTO ban(first_user, second_user) VALUES (?, ?)`, dbUser.Id, bannedDbUser.Id)

	return err
}

func (db *appdbimpl) DeleteBan(dbUser DatabaseUser, bannedDbUser DatabaseUser) error {
	res, err := db.c.Exec(`DELETE FROM ban WHERE first_user=? AND second_user=?`, dbUser.Id, bannedDbUser.Id)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if aff == 0 {
		return ErrUserNotBanned
	}

	return nil
}