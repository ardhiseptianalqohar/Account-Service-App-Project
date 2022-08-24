package ReadAccount

import (
	"Project/entities"
	"database/sql"
	"fmt"
)

func ReadAccount(db *sql.DB) {
	results, errselect := db.Query("select Id, Name, Address, Gender, Status, Nomor, Password from Profile")
	if errselect != nil {
		fmt.Println("error select", errselect.Error())
	}

	var dataUser []entities.Profile
	for results.Next() {
		var rowdata entities.Profile
		errScan := results.Scan(&rowdata.Id, &rowdata.Name, &rowdata.Address, &rowdata.Gender, &rowdata.Status, &rowdata.Nomor, &rowdata.Password)

		if errScan != nil {
			fmt.Println("error scan", errScan.Error())
		}
		dataUser = append(dataUser, rowdata)
	}
	for _, v := range dataUser {
		fmt.Println("Id", v.Id, "Name", v.Name, "Address", v.Address, "Gender", v.Gender,"Status", v.Status, "Nomor", v.Nomor, "Password", v.Password)
	}
}