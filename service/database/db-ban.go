package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) InsertBan(dbUser DatabaseUser, bannedDbUser DatabaseUser) error {
	// insert the ban into the database
	_, err := db.c.Exec(`
		INSERT OR IGNORE INTO ban(first_user, second_user)
		VALUES (?, ?)
	`, dbUser.Id, bannedDbUser.Id)

	return err
}

func (db *appdbimpl) DeleteBan(dbUser DatabaseUser, bannedDbUser DatabaseUser) error {
	// remove the ban from the database
	res, err := db.c.Exec(`
		DELETE FROM ban
		WHERE first_user=?
		AND second_user=?
	`, dbUser.Id, bannedDbUser.Id)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return err
	}

	// if there are no affected rows
	// then the user was not banned
	if aff == 0 {
		return ErrUserNotBanned
	}

	return nil
}

func (db *appdbimpl) CheckBan(firstDbUser DatabaseUser, secondDbUser DatabaseUser) (bool, error) {
	checkBan := true

	// check wether the first user has banned the second user
	err := db.c.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM ban
			WHERE first_user=?
			AND second_user=?
		)
	`, firstDbUser.Id, secondDbUser.Id).Scan(&checkBan)

	// if no table rows are found, then there is no row
	// containing the ban, hence the first user has not
	// banned the second user
	if errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	return checkBan, err
}

func (db *appdbimpl) GetBanList(dbUser DatabaseUser) (DatabaseUserList, error) {
	dbUserList := DatabaseUserListDefault()

	// get the table of the banned users
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

	// build the banned users list
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
