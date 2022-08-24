package update

import (
	"Project/entities"
	"database/sql"
	// "fmt"
)

func UpdateName(db *sql.DB, profile entities.Profile) (int, error) {
	statment, err := db.Prepare("update profile set name = ? where nomor = ?")
	if err != nil {
		return -1, err
	}

	row, errex := statment.Exec(profile.Name, profile.Nomor)
	if errex != nil {
		return -1, errex
	}

	eff, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(eff), nil
}

func UpdateAddress(db *sql.DB, profile entities.Profile) (int, error) {
	statment, err := db.Prepare("update profile set address = ? where nomor = ?")
	if err != nil {
		return -1, err
	}

	row, errex := statment.Exec(profile.Address, profile.Nomor)
	if errex != nil {
		return -1, errex
	}

	eff, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(eff), nil
}

func UpdateNomor(db *sql.DB, profile entities.Profile) (int, error) {
	statment, err := db.Prepare("update profile set nomor = ? where name = ?")
	if err != nil {
		return -1, err
	}

	row, errex := statment.Exec(profile.Nomor, profile.Name)
	if errex != nil {
		return -1, errex
	}

	eff, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(eff), nil
}

func UpdatePassword(db *sql.DB, profile entities.Profile) (int, error) {
	statment, err := db.Prepare("update profile set password = ? where nomor = ?")
	if err != nil {
		return -1, err
	}

	row, errex := statment.Exec(profile.Password, profile.Nomor)
	if errex != nil {
		return -1, errex
	}

	eff, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(eff), nil
}
