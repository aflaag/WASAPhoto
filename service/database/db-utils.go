package database

func (db *appdbimpl) CheckIfAvailablePhotoId(newId uint64) (bool, error) {
	res, err := db.c.Exec(`SELECT id FROM Photo WHERE id=?`, newId)

	if err != nil {
		return false, err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return false, err
	}

	return aff == 0, err
}