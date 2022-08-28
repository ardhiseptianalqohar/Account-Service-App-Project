package search

import (
	"Project/entities"
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

func Other(db *sql.DB) {
	cari := entities.Profile{}
	fmt.Println("Masukkan Nama User :")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	cari.Name = line

	row, err := Search(db, cari)
	if err != nil {
		fmt.Println("Tidak di temukan", err)
	}

	fmt.Printf("Name : %sAddress : %s\nGender : %s\nStasus : %s\n", row.Name, row.Address, row.Gender, row.Status)
}

func Search(db *sql.DB, search entities.Profile) (entities.Profile, error) {
	stat, err := db.Prepare("select name, address, gender, status from profile where name = ?")
	if err != nil {
		return entities.Profile{}, err
	}

	var info entities.Profile
	errs := stat.QueryRow(search.Name).Scan(&info.Name, &info.Address, &info.Gender, &info.Status)
	if errs != nil {
		return entities.Profile{}, err
	}

	return info, nil
}
