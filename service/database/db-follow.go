package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) InsertFollow(dbUser DatabaseUser, followedDbUser DatabaseUser) error {
	_, err := db.c.Exec(`INSERT INTO follow(first_user, second_user) VALUES (?, ?)`, dbUser.Id, followedDbUser.Id)

	return err
}

func (db *appdbimpl) DeleteFollow(dbUser DatabaseUser, followedDbUser DatabaseUser) error {
	res, err := db.c.Exec(`DELETE FROM follow WHERE first_user=? AND second_user=?`, dbUser.Id, followedDbUser.Id)

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

// TODO: func (db *appdbimpl) GetFollowers(dbUser DatabaseUser) ([]DatabaseUser, error)

// TODO: func (db *appdbimpl) GetFollowing(dbUser DatabaseUser) ([]DatabaseUser, error)

func (db *appdbimpl) GetFollowersCount(dbUser DatabaseUser) (int, error) {
	var followersCount int

	err := db.c.QueryRow(`SELECT COUNT(*) FROM follow WHERE second_user=?`, dbUser.Id).Scan(&followersCount)

	if errors.Is(err, sql.ErrNoRows) {
		return followersCount, ErrUserDoesNotExist
	}

	return followersCount, err
}

func (db *appdbimpl) GetFollowingCount(dbUser DatabaseUser) (int, error) {
	var followingCount int

	err := db.c.QueryRow(`SELECT COUNT(*) FROM follow WHERE first_user=?`, dbUser.Id).Scan(&followingCount)

	if errors.Is(err, sql.ErrNoRows) {
		return followingCount, ErrUserDoesNotExist
	}

	return followingCount, err
}
