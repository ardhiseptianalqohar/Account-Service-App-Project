package topup

import (
	"Project/entities"
	"database/sql"
	"fmt"
)

func TopUpAcc(db *sql.DB, top entities.Profile, topup int) (string, error) {
	// mencari id
	statid, erra := db.Prepare("select id, saldo from profile where nomor = ?")
	if erra != nil {
		return "Gagal", erra
	}

	var id entities.Profile
	errorq := statid.QueryRow(top.Nomor).Scan(&id.Id, &id.Saldo)
	if errorq != nil {
		return "Gagal", errorq
	}

	//memasukkan ke history
	stat, errh := db.Prepare("insert into topup_history(Jumlah, User_Id) values (?, ?)")
	if errh != nil {
		return "gagal terhubung ke history", errh
	}
	result, erri := stat.Exec(topup, id.Id)
	if erri != nil {
		return "gagal menghubungkan ke history", erri
	}

	// menghitung jumlah total saldo yang ditambahkan
	total := id.Saldo + topup

	//memasukkan jumlah saldo pada database
	stat, err := db.Prepare("update profile set saldo = ? where nomor = ?")
	if err != nil {
		return "Silahkan periksa kembali no anda", err
	}

	result, errs := stat.Exec(total, top.Nomor)
	if errs != nil {
		return "Terjadi Kesalahan", errs
	}
	effect, errf := result.RowsAffected()
	if errf != nil {
		return "Terjadi Kesalahan", errf
	} else if effect == 0 {
		return "Gagal Menambahkan saldo", nil
	}

	suksec := fmt.Sprintf("Sukses Melakukan Top up sebesar %d", topup)

	return suksec, nil
}
