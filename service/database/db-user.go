package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetDatabaseUser(userId uint32) (DatabaseUser, error) {
	dbUser := DatabaseUserDefault()

	// get the user having the given user id
	err := db.c.QueryRow(`
		SELECT id, username
		FROM User
		WHERE id=?
	`, userId).Scan(&dbUser.Id, &dbUser.Username)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUser, ErrUserDoesNotExist
	}

	return dbUser, err
}

func (db *appdbimpl) GetDatabaseUserFromDatabaseLogin(dbLogin DatabaseLogin) (DatabaseUser, error) {
	dbUser := DatabaseUserDefault()

	// get the user from the given login instance
	err := db.c.QueryRow(`
		SELECT id, username
		FROM User
		WHERE username=?
	`, dbLogin.Username).Scan(&dbUser.Id, &dbUser.Username)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUser, ErrUserDoesNotExist
	}

	return dbUser, err
}

func (db *appdbimpl) InsertUser(dbUser *DatabaseUser) error {
	// check if the user is already registered
	err := db.c.QueryRow(`
		SELECT id
		FROM User
		WHERE username=?
	`, dbUser.Username).Scan(&dbUser.Id)

	if err != nil {
		// if there are no rows, the user was not registered
		// hence it must be inserted into the database
		if errors.Is(err, sql.ErrNoRows) {
			// insert the new user into the database
			res, err := db.c.Exec(`
				INSERT INTO User(username)
				VALUES (?)
			`, dbUser.Username)

			if err != nil {
				return err
			}

			// get the user id
			dbUserId, err := res.LastInsertId()

			if err != nil {
				return err
			}

			dbUser.Id = uint32(dbUserId)

			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}

func (db *appdbimpl) UpdateUser(oldDbUser DatabaseUser, newDbUser DatabaseUser) error {
	res, err := db.c.Exec(`
		UPDATE User
		SET username=?
		WHERE id=?
		AND username=?
	`, newDbUser.Username, oldDbUser.Id, oldDbUser.Username)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if aff == 0 {
		return ErrUserDoesNotExist
	}

	return nil
}

func (db *appdbimpl) GetUserList(dbUser DatabaseUser, dbLogin DatabaseLogin) (DatabaseUserList, error) {
	dbUserList := DatabaseUserListDefault()

	// get the table of the users matching the query
	rows, err := db.c.Query(`
		SELECT id, username
		FROM User
		WHERE id IN (
			SELECT id
			FROM User
			WHERE username LIKE '%'||?||'%'
			EXCEPT 
			SELECT first_user
			FROM ban
			WHERE second_user=?
			EXCEPT
			SELECT ?
		)
	`, dbLogin.Username, dbUser.Id, dbUser.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUserList, ErrUserDoesNotExist
	}

	if err != nil {
		return dbUserList, err
	}

	// build the results list
	for rows.Next() {
		newDbUser := DatabaseUserDefault()

		err = rows.Scan(&newDbUser.Id, &newDbUser.Username)

		if err != nil {
			return dbUserList, err
		}

		dbUserList.Users = append(dbUserList.Users, newDbUser)
	}

	if rows.Err() != nil {
		return dbUserList, err
	}

	_ = rows.Close()

	return dbUserList, err
}
