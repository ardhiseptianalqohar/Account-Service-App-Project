package profile

import (
	"Project/entities"
	"database/sql"
	"fmt"
)

func AddAccount(db *sql.DB) {
	daftar := entities.Profile{}
	fmt.Println("Name:")
	fmt.Scanln(&daftar.Name)
	fmt.Println("Address:")
	fmt.Scanln(&daftar.Address)
	fmt.Println("Gender:")
	fmt.Scanln(&daftar.Gender)
	fmt.Println("Nomor:")
	fmt.Scanln(&daftar.Nomor)
	fmt.Println("Password:")
	fmt.Scanln(&daftar.Password)

	var query = "Insert into Profile (Name, Address, Gender, Nomor, Password) values (?, ?, ?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		fmt.Println("error prepare insert", errPrepare.Error())
	}
	result, errExec := statement.Exec(daftar.Name, daftar.Address, daftar.Gender, daftar.Nomor, daftar.Password)
	if errExec != nil {
		fmt.Println("error exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Succes Insert Data")
		} else {
			fmt.Println("gagal insert")
		}
	}
}