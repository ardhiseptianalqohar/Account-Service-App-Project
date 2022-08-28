package main

import (
	"Project/config"
	profile "Project/controllers/Profile"
	"Project/controllers/ReadAccount"
	"Project/controllers/deleteaccount"
	"Project/controllers/login"
	"Project/controllers/search"
	"Project/controllers/topup"
	"Project/controllers/topuphistory"
	"Project/controllers/transfer"
	"Project/controllers/update"

	//"database/sql"
	"fmt"
)

func main() {
	//test
	db := config.ConnectToDB()

	defer db.Close()
	run := true
	for run {
		fmt.Print("MENU:\n1. Add Account(register)\n2. Login\n0. Keluar\n")
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
				login.Logins(db)
				run2 := true
				for run2 {
					fmt.Print("MENU:\n3. Read Account\n4. Update Account\n5. Delete Account\n6. Top-Up\n7. Transfer\n8. History Top-Up\n9. History Transfer\n10. Liat User Lain\n0. Keluar\n")
					fmt.Println("Masukkan piihan anda:")
					var pilihan2 int
					fmt.Scanln(&pilihan2)

					switch pilihan2 {
					case 3:
						{
							ReadAccount.ReadAccount(db)
						}
					case 4:
						{
							update.UpdateAll(db)
						}
					case 5:
						{
							deleteaccount.DeleteAkun(db)
						}
					case 6:
						{
							topup.TopUp(db)
						}
					case 7:
						{
							transfer.Transfer(db)
						}
					case 8:
						{
							topuphistory.Historys(db)
						}
					case 9:
						{
						}
					case 10:
						{
							search.Other(db)
						}
					case 0:
						{
							run2 = false
						}
					}
				}
			}
		case 0:
			{
				run = false
			}
		}
	}
}
