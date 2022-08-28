package profile

import (
	"Project/entities"
	"database/sql"
	"fmt"
)

func AddAccount(db *sql.DB) {

	fmt.Println("======================")
	fmt.Println("FITUR REGISTRASI AKUN")
	fmt.Println("======================")

	daftar := entities.Profile{}
	fmt.Println("Name:")
	fmt.Scanln(&daftar.Name)
	fmt.Println("Address:")
	fmt.Scanln(&daftar.Address)
	fmt.Println("Gender:")
	fmt.Scanln(&daftar.Gender)
	fmt.Println("Status:")
	fmt.Scanln(&daftar.Status)
	fmt.Println("Nomor:")
	fmt.Scanln(&daftar.Nomor)
	fmt.Println("Password:")
	fmt.Scanln(&daftar.Password)

	var query = "Insert into Profile (Name, Address, Gender, Saldo, Status, Nomor, Password) values (?, ?, ?, ?, ?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		fmt.Println("error prepare insert", errPrepare.Error())
	}
	result, errExec := statement.Exec(daftar.Name, daftar.Address, daftar.Gender, 0, daftar.Status, daftar.Nomor, daftar.Password)
	if errExec != nil {
		fmt.Println("error exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("==========================================")
			fmt.Println("ANDA SUKSES MELAKUKAN PENDAFTARAN")
			fmt.Println("SELAMAT MENIKMATI FITUR DAN LAYANAN KAMI")
			fmt.Println("==========================================")

		} else {
			fmt.Println("gagal insert")
		}
	}
}
