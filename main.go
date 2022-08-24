package main

import (
	"Project/config"
	"Project/controllers/Profile"
	"Project/controllers/login"
	"Project/controllers/update"
	"Project/entities"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//test
	db := config.ConnectToDB()

	defer db.Close()

	fmt.Print("MENU:\n1. Add Account(register)\n2. Login\n3. Read Account\n4. Update Account\n5. Delete Account\n6. Top-Up\n7. Transfer\n8. History Top-Up\n9. History Transfer\n10. Liat User Lain\n0. Keluar")
	fmt.Println("Masukkan piihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
			profile.AddAccount(db)
		}
	case 2:
		{
			in := entities.Profile{}
			fmt.Println("Silahkan Masukkan Nomor Telepon Anda :")
			fmt.Scanln(&in.Nomor)
			fmt.Println("Silahkan Masukkan Password Anda :")
			fmt.Scanln(&in.Password)

			row, err := login.LoginUser(db, in)
			if err != nil {
				fmt.Println("Gagal Login", err)
			}

			fmt.Println("Selamat Datang", row.Name,
				"\n\nSaldo yang Anda Miliki adalah", row.Saldo,
				"\nAlamat saat ini", row.Address,
				"\nStatus Saat ini", row.Status)

		}
	case 3:
		{
		}
	case 4:
		{
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

					row, err := update.UpdateName(db, up)
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

					row, err := update.UpdateAddress(db, up)
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

					row, err := update.UpdateNomor(db, up)
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

					row, err := update.UpdatePassword(db, up)
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
	case 5:
		{
		}
	case 6:
		{
		}
	case 7:
		{
		}
	case 8:
		{
		}
	case 9:
		{
		}
	case 10:
		{
		}
	case 0:
		{
		}
	}
}
