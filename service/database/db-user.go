package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetDatabaseUser(userId uint32) (DatabaseUser, error) {
	dbUser := DatabaseUserDefault()

	err := db.c.QueryRow(`SELECT id, username FROM User WHERE id=?`, userId).Scan(&dbUser.Id, &dbUser.Username)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUser, ErrUserDoesNotExist
	}

	return dbUser, err
}

func (db *appdbimpl) GetDatabaseUserFromDatabaseLogin(dbLogin DatabaseLogin) (DatabaseUser, error) {
	dbUser := DatabaseUserDefault()

	err := db.c.QueryRow(`SELECT id, username FROM User WHERE username=?`, dbLogin.Username).Scan(&dbUser.Id, &dbUser.Username)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUser, ErrUserDoesNotExist
	}

	return dbUser, err
}

func (db *appdbimpl) InsertUser(dbUser *DatabaseUser) error {
	res, err := db.c.Exec("INSERT INTO User(username) VALUES (?)", dbUser.Username)

	if err != nil {
		return err
	}

	dbUserId, err := res.LastInsertId()

	if err != nil {
		return err
	}

	dbUser.Id = uint32(dbUserId)

	return nil
}

func (db *appdbimpl) UpdateUser(oldDbUser DatabaseUser, newDbUser DatabaseUser) error {
	res, err := db.c.Exec(`UPDATE User SET username=? WHERE id=? AND username=?`, newDbUser.Username, oldDbUser.Id, oldDbUser.Username)

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

func (db *appdbimpl) GetUserList(dbUser DatabaseUser, dbLogin DatabaseLogin) (DatabaseUserList, error) {
	dbUserList := DatabaseUserListDefault()

	rows, err := db.c.Query(`
		SELECT id
		FROM User
		WHERE username LIKE '%'||?||'%'
		EXCEPT 
		SELECT first_user
		FROM ban
		WHERE second_user=?
		EXCEPT
		SELECT ?
	`, dbLogin.Username, dbUser.Id, dbUser.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return dbUserList, ErrUserDoesNotExist
	}

	if err != nil {
		return dbUserList, err
	}

	for rows.Next() {
		var newDbUserId uint32

		err = rows.Scan(&newDbUserId)

		if err != nil {
			return dbUserList, err
		}

		newDbUser, err := db.GetDatabaseUser(newDbUserId)

		if err != nil {
			return dbUserList, err
		}

		dbUserList.Users = append(dbUserList.Users, newDbUser)
	}

	_ = rows.Close()

	return dbUserList, err
}
