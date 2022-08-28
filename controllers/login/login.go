package login

import (
	"Project/entities"
	"database/sql"
	"fmt"
)

func Logins(db *sql.DB) {
	in := entities.Profile{}
	fmt.Println("Silahkan Masukkan Nomor Telepon Anda :")
	fmt.Scanln(&in.Nomor)
	fmt.Println("Silahkan Masukkan Password Anda :")
	fmt.Scanln(&in.Password)

	row, err := LoginUser(db, in)
	if err != nil {
		fmt.Println("Gagal Login", err)
	}

	fmt.Println("Selamat Datang", row.Name,
		"\nSaldo yang Anda Miliki adalah", row.Saldo,
		"\nAlamat saat ini", row.Address,
		"\nStatus Saat ini", row.Status)

}

func LoginUser(db *sql.DB, profile entities.Profile) (entities.Profile, error) {
	statm, err := db.Prepare("select name, address, gender, saldo, status from profile where nomor = ? and password = ?")
	if err != nil {
		// fmt.Println("Statment Error", err)
		return entities.Profile{}, err
	}

	var row entities.Profile
	errs := statm.QueryRow(profile.Nomor, profile.Password).Scan(&row.Name, &row.Address, &row.Gender, &row.Saldo, &row.Status)
	if errs != nil {
		// fmt.Println("Error", errs)
		return entities.Profile{}, errs

	}

	return row, nil
}

// Id, Saldo                                      int
// Name, Address, Gender, Status, Nomor, Password
