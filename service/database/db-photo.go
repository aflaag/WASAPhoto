package database

import (
	// "database/sql"
)

func (db *appdbimpl) SetPhoto(dbUser DatabaseUser, dbPhoto DatabasePhoto) error {
	_, err := db.c.Exec(`INSERT INTO Photo (id, user, url, date) VALUES (?, ?, ?, ?)`, dbPhoto.Id, dbUser.Id, dbPhoto.Url, dbPhoto.Date)

	return err
}

func (db *appdbimpl) RemovePhoto(dbPhoto DatabasePhoto) error {
	var err error

	_, err = db.c.Exec(`DELETE FROM like WHERE photo=?`, dbPhoto.Id)

	if err != nil {
		return err
	}

	_, err = db.c.Exec(`DELETE FROM comment WHERE photo=?`, dbPhoto.Id)

	if err != nil {
		return err
	}

	_, err = db.c.Exec(`DELETE FROM Photo WHERE photo=?`, dbPhoto.Id)

	if err != nil {
		return err
	}

	return nil
}