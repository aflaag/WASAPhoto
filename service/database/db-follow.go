package database

import (
	"database/sql"
)

func (db *appdbimpl) SetFollow(user DatabaseUser, followedUser DatabaseUser) error {
	_, err := db.c.Exec(`INSERT INTO follow (first_user, second_user) VALUES (?, ?)`, user.Id, followedUser.Id)

	if err != nil {
		return err
	}
	
	return nil
}

func (db *appdbimpl) RemoveFollow(user DatabaseUser, followedUser DatabaseUser) error {
	res, err := db.c.Exec(`DELETE FROM follow WHERE first_user=? AND second_user=?`, user.Id, followedUser.Id)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return err
	} else if aff == 0 {
		return ErrUserNotFollowed
	}

	return nil
}

// func (db *appdbimpl) GetFollowers()

func (db *appdbimpl) GetFollowersCount(user DatabaseUser) (int, error) {
	var followersCount int

	err := db.c.QueryRow(`SELECT COUNT(*) FROM follow WHERE second_user=?`, user.Id).Scan(&followersCount)

	if err != nil {
		if err == sql.ErrNoRows {
			return followersCount, ErrUserDoesNotExist
		}
		
		return followersCount, err
	}

	return followersCount, nil
}

func (db *appdbimpl) GetFollowingCount(user DatabaseUser) (int, error) {
	var followingCount int

	err := db.c.QueryRow(`SELECT COUNT(*) FROM follow WHERE first_user=?`, user.Id).Scan(&followingCount)

	if err != nil {
		if err == sql.ErrNoRows {
			return followingCount, ErrUserDoesNotExist
		}
		
		return followingCount, err
	}

	return followingCount, nil
}