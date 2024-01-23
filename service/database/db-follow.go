package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) InsertFollow(dbUser DatabaseUser, followedDbUser DatabaseUser) error {
	// insert the following into the database
	_, err := db.c.Exec(`
		INSERT OR IGNORE INTO follow(first_user, second_user)
		VALUES (?, ?)
	`, dbUser.Id, followedDbUser.Id)

	return err
}

func (db *appdbimpl) DeleteFollow(dbUser DatabaseUser, followedDbUser DatabaseUser) error {
	// remove the following from the database
	res, err := db.c.Exec(`
		DELETE FROM follow
		WHERE first_user=?
		AND second_user=?
	`, dbUser.Id, followedDbUser.Id)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return err
	}

	// if there are no affected rows
	// then the user was not followed
	if aff == 0 {
		return ErrUserNotFollowed
	}

	return nil
}

func (db *appdbimpl) GetFollowersCount(profileDbUser DatabaseUser, dbUser DatabaseUser) (int, error) {
	var followersCount int

	// get the number of user following
	// the user performing the action
	err := db.c.QueryRow(`
		SELECT COUNT(*)
		FROM follow
		WHERE second_user=?
		AND first_user NOT IN (
			SELECT first_user
			FROM ban
			WHERE second_user=?
		)
	`, profileDbUser.Id, dbUser.Id).Scan(&followersCount)

	if errors.Is(err, sql.ErrNoRows) {
		return followersCount, ErrUserDoesNotExist
	}

	return followersCount, err
}

func (db *appdbimpl) GetFollowingCount(profileDbUser DatabaseUser, dbUser DatabaseUser) (int, error) {
	var followingCount int

	// get the number of users followed by
	// the user performing the action
	err := db.c.QueryRow(`
		SELECT COUNT(*)
		FROM follow
		WHERE first_user=?
		AND second_user NOT IN (
			SELECT first_user
			FROM ban
			WHERE second_user=?
		)
	`, profileDbUser.Id, dbUser.Id).Scan(&followingCount)

	if errors.Is(err, sql.ErrNoRows) {
		return followingCount, ErrUserDoesNotExist
	}

	return followingCount, err
}

func (db *appdbimpl) GetFollowersList(followersDbUser DatabaseUser, dbUser DatabaseUser) (DatabaseUserList, error) {
	dbUserList := DatabaseUserListDefault()

	// get the table of the followers
	// without the users who banned the user performing the action
	rows, err := db.c.Query(`
		SELECT id, username
		FROM User
		WHERE id IN (
			SELECT first_user
			FROM follow
			WHERE second_user=?
		)
		AND id NOT IN (
			SELECT first_user
			FROM ban
			WHERE second_user=?
		)
	`, followersDbUser.Id, dbUser.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUserList, ErrUserDoesNotExist
	}

	if err != nil {
		return dbUserList, err
	}

	// build the followers list
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

func (db *appdbimpl) GetFollowingList(followingDbUser DatabaseUser, dbUser DatabaseUser) (DatabaseUserList, error) {
	dbUserList := DatabaseUserListDefault()

	// get the table of the followed
	rows, err := db.c.Query(`
		SELECT id, username
		FROM User
		WHERE id IN (
			SELECT second_user
			FROM follow
			WHERE first_user=?
		)
		AND id NOT IN (
			SELECT first_user
			FROM ban
			WHERE second_user=?
		)
	`, followingDbUser.Id, dbUser.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUserList, ErrUserDoesNotExist
	}

	if err != nil {
		return dbUserList, err
	}

	// build the following list
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

func (db *appdbimpl) GetFollowStatus(firstDbUser DatabaseUser, secondDbUser DatabaseUser) (bool, error) {
	followStatus := false

	// check whether the first user follows the second user
	err := db.c.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM follow
			WHERE first_user=?
			AND second_user=?
		)
	`, firstDbUser.Id, secondDbUser.Id).Scan(&followStatus)

	// if no table rows are found, then there is no row
	// containing the ban, hence the first user has not
	// banned the second user
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return followStatus, err
}
