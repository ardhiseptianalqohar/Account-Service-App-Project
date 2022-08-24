package login

import (
	"Project/entities"
	"database/sql"
	"fmt"
)

func LoginUser(db *sql.DB, profile entities.Profile) (entities.Profile, error) {
	statm, err := db.Prepare("select name, address, gender, saldo, status from profile where nomor = ? and password = ?")
	if err != nil {
		fmt.Println("Statment Error", err)
	}

	var row entities.Profile
	errs := statm.QueryRow(profile.Nomor, profile.Password).Scan(&row.Name, &row.Address, &row.Gender, &row.Saldo, &row.Status)
	if errs != nil {
		fmt.Println("Error", errs)
	}

	return row, nil
}

// Id, Saldo                                      int
// Name, Address, Gender, Status, Nomor, Password
