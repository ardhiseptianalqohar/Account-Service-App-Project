package topuphistory

import (
	"Project/entities"
	"database/sql"
	"fmt"
)

func Historys(db *sql.DB) {
	hist := entities.Profile{}
	fmt.Println("Masukkan Nomor Anda :")
	fmt.Scanln(&hist.Nomor)

	history, name, err := CheckHistory(db, hist)
	if err != nil {
		fmt.Println("Gagal Memuat history", err)
	}

	fmt.Println("History Top Up untuk User", name)
	for _, v := range history {
		fmt.Println("Jumlah :", v.Jumlah, "Dilakukan Pada :", v.Date)
	}
}

func CheckHistory(db *sql.DB, check entities.Profile) ([]entities.Top_UpHistory, string, error) {
	stat, err := db.Prepare("select name, id from profile where nomor = ?")
	if err != nil {
		return []entities.Top_UpHistory{}, "", err
	}

	var test entities.Profile
	errs := stat.QueryRow(check.Nomor).Scan(&test.Name, &test.Id)
	if err != nil {
		return []entities.Top_UpHistory{}, "", errs
	}

	query := fmt.Sprintf("select h.jumlah, h.date from profile p inner join topup_history h on p.id = h.user_id where p.id = %d", test.Id)

	results, errsel := db.Query(query)
	if errsel != nil {
		return []entities.Top_UpHistory{}, "", errsel
	}

	var rows []entities.Top_UpHistory
	for results.Next() {

		var row entities.Top_UpHistory
		errScan := results.Scan(&row.Jumlah, &row.Date)
		if errScan != nil {
			return []entities.Top_UpHistory{}, "", errScan
		}
		rows = append(rows, row)
	}

	return rows, test.Name, nil
}
