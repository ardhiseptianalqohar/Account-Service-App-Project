package main

import (
	"Project/config"
	"fmt"
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
		}
	case 2:
		{
		}
	case 3:
		{
		}
	case 4:
		{
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
