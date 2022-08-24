package deleteaccount

import (
	"Project/entities"
	"database/sql"
	"fmt"
)

func DeleteAkun(db *sql.DB) {
	delete := entities.Profile {}
	fmt.Println("Delete Account:")
	fmt.Scanln(&delete.Id)
	fmt.Println("Password:")
	fmt.Scanln(&delete.Password)
	
	var query = "Delete from profile where Id = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		fmt.Println("error prepare delete", errPrepare.Error())
	}
	result, errExec := statement.Exec(delete.Id)
	if errExec != nil {
		fmt.Println("error exec delete", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Succes delete account")
		} else {
			fmt.Println("Gagal menghapus akun, dan coba lagi..!!")
		}
	}
	}