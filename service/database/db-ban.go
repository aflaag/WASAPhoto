package database

import (
	"database/sql"
	"errors"
)

// "database/sql"

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

func (db *appdbimpl) CheckBan(firstDbUser DatabaseUser, secondDbUser DatabaseUser) (bool, error) {
	checkBan := true

	err := db.c.QueryRow(`SELECT EXISTS(SELECT 1 FROM ban WHERE first_user=? AND second_user=?)`, firstDbUser.Id, secondDbUser.Id).Scan(&checkBan)

	if errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	return checkBan, err
}

func (db *appdbimpl) GetBanList(dbUser DatabaseUser) (DatabaseUserList, error) {
	dbUserList := DatabaseUserListDefault()

	rows, err := db.c.Query(`
		SELECT id, username
		FROM User
		WHERE id IN (
			SELECT second_user
			FROM ban
			WHERE first_user=?
		)
	`, dbUser.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUserList, ErrUserDoesNotExist
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
