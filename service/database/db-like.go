package database

import (
	"database/sql"
)

func (db *appdbimpl) SetLike(user DatabaseUser, photo DatabasePhoto) error {
	_, err := db.c.Exec(`INSERT INTO like (user, photo) VALUES (?, ?)`, user.Id, photo.Id)

	if err != nil {
		return err
	}
	
	return nil
}

func (db *appdbimpl) RemoveLike(user DatabaseUser, photo DatabasePhoto) error {
	res, err := db.c.Exec(`DELETE FROM follow WHERE user=? AND photo=?`, user.Id, photo.Id)

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

func (db *appdbimpl) GetLikesCount(photo DatabasePhoto) (int, error) {
	var likesCount int

	err := db.c.QueryRow(`SELECT COUNT(*) FROM like WHERE photo=?`, photo.Id).Scan(&likesCount)

	if err != nil {
		if err == sql.ErrNoRows {
			return likesCount, ErrUserDoesNotExist
		}
		
		return likesCount, err
	}

	return likesCount, nil
}