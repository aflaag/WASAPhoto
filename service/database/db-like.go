package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) InsertLike(dbUser DatabaseUser, dbPhoto DatabasePhoto) error {
	_, err := db.c.Exec(`INSERT INTO like(user, photo) VALUES (?, ?)`, dbUser.Id, dbPhoto.Id)

	return err
}

func (db *appdbimpl) DeleteLike(dbUser DatabaseUser, dbPhoto DatabasePhoto) error {
	res, err := db.c.Exec(`DELETE FROM like WHERE user=? AND photo=?`, dbUser.Id, dbPhoto.Id)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if aff == 0 {
		return ErrPhotoNotLiked
	}

	return err
}

func (db *appdbimpl) GetLikesCount(dbPhoto DatabasePhoto) (int, error) {
	var likesCount int

	err := db.c.QueryRow(`SELECT COUNT(*) FROM like WHERE photo=?`, dbPhoto.Id).Scan(&likesCount)

	if errors.Is(err, sql.ErrNoRows) {
		return likesCount, ErrUserDoesNotExist
	}

	return likesCount, err
}
