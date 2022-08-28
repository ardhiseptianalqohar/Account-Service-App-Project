package update

import (
	"Project/entities"
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

func UpdateAll(db *sql.DB) {

	fmt.Println("Bagian Yang Di Update\n1. Nama\n2. Address\n3. Nomor\n4. Password")
	fmt.Print("Masukkan pilihan:")
	var pilih int
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		{
			up := entities.Profile{}
			fmt.Println("Masukkan Nama Terbaru Anda :")
			in := bufio.NewReader(os.Stdin)
			line, _ := in.ReadString('\n')
			up.Name = line
			// fmt.Scanln(&up.Name)
			fmt.Println("Masukkan Nomor Anda :")
			fmt.Scanln(&up.Nomor)

			row, err := UpdateName(db, up)
			if err != nil {
				fmt.Println("Terjadi error", err)
			} else if row == 0 {
				fmt.Println("gagal update, row affected = 0")
			} else {
				fmt.Println("sukses update, row affected", row)
			}
		}
	case 2:
		{
			up := entities.Profile{}
			fmt.Println("Masukkan Alamat Terbaru Anda :")
			fmt.Scanln(&up.Address)
			fmt.Println("Masukkan Nomor Anda :")
			fmt.Scanln(&up.Nomor)

			row, err := UpdateAddress(db, up)
			if err != nil {
				fmt.Println("Terjadi error", err)
			} else if row == 0 {
				fmt.Println("gagal update, row affected = 0")
			} else {
				fmt.Println("sukses update, row affected", row)
			}
		}
	case 3:
		{
			up := entities.Profile{}
			fmt.Println("Masukkan Nomor Terbaru Anda :")
			fmt.Scanln(&up.Nomor)
			fmt.Println("Masukkan Nama Anda :")
			in := bufio.NewReader(os.Stdin)
			line, _ := in.ReadString('\n')
			up.Name = line

			row, err := UpdateNomor(db, up)
			if err != nil {
				fmt.Println("Terjadi error", err)
			} else if row == 0 {
				fmt.Println("gagal update, row affected = 0")
			} else {
				fmt.Println("sukses update, row affected", row)
			}
		}
	case 4:
		{
			up := entities.Profile{}
			fmt.Println("Masukkan Password Terbaru Anda :")
			fmt.Scanln(&up.Password)
			fmt.Println("Masukkan Nomor Anda :")
			fmt.Scanln(&up.Nomor)

			row, err := UpdatePassword(db, up)
			if err != nil {
				fmt.Println("Terjadi error", err)
			} else if row == 0 {
				fmt.Println("gagal update, row affected = 0")
			} else {
				fmt.Println("sukses update, row affected", row)
			}
		}
	}
}

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
