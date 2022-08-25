package search

import (
	"Project/entities"
	"database/sql"
)

func Search(db *sql.DB, search entities.Profile) (entities.Profile, error) {
	stat, err := db.Prepare("select name, address, gender, status from profile where name = ?")
	if err != nil {
		return entities.Profile{}, err
	}

	var info entities.Profile
	errs := stat.QueryRow(search.Name).Scan(&info.Name, &info.Address, &info.Gender, &info.Status)
	if errs != nil {
		return entities.Profile{}, err
	}

	return info, nil
}
