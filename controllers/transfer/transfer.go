package transfer

import (
	"Project/entities"
	"database/sql"
	"fmt"
)

func Transfer(db *sql.DB) {

	fmt.Println("==============")
	fmt.Println("FITUR TRANSFER")
	fmt.Println("==============")

	// membaca data saldo pengirim dan penerima
	oke := entities.Profile{}
	penerima := entities.Profile{}

	fmt.Println("Masukan Nomor Akun Anda :")
	fmt.Scanln(&oke.Nomor)
	fmt.Println("Masukan Nomor Penerima :")
	fmt.Scanln(&penerima.Nomor)

	statm, err := db.Prepare("select id, saldo from profile where nomor = ?")
	if err != nil {
		fmt.Println("SILAHKAN MELANJUTKAN UNTUK PROSES TRANSFER NYA YO MASSEE", statm)
	}
	var in entities.Profile
	errs := statm.QueryRow(oke.Nomor).Scan(&in.Id, &in.Saldo)
	if errs != nil {
		fmt.Println("Gagal bang bro, ulang lagi donk..!!", errs.Error())
	}

	//==LANJUT KE TRANSFER KE NOMOR PENERIMA==

	var transfer int
	fmt.Println("Masukan Nominal Transfer :")
	fmt.Scanln(&transfer)

	if in.Saldo < transfer {
		fmt.Println("Saldo Anda Tidak Mencukupi")
	}

	stat, err := db.Prepare("select id, saldo from profile where nomor = ?")
	if err != nil {
		fmt.Println("Penerima Tidak Ditemukan", err.Error())
	}
	var terima entities.Profile
	er := stat.QueryRow(penerima.Nomor).Scan(&terima.Id, &terima.Saldo)
	if er != nil {
		fmt.Println("Penerima Tidak Ditemukan", er.Error())
	}

	stat, err = db.Prepare("insert into transfer_history (jumlah, send_id, accept_id) values (?, ?, ?) ")
	if err != nil {
		fmt.Println("Gagal Mengirim", err.Error())
	}
	_, error := stat.Exec(transfer, in.Id, terima.Id)
	if error != nil {
		fmt.Println("Gagal Ke Menu Transfer History", error.Error())
	}

	stat, err = db.Prepare("update profile set saldo = ? where nomor = ?")
	if err != nil {
		fmt.Println("Gagal update", err.Error())
	}
	saldo := in.Saldo - transfer
	saldo2 := terima.Saldo + transfer

	_, error = stat.Exec(saldo, oke.Nomor)
	if error != nil {
		fmt.Println("Gagal", error.Error())
	}
	_, error = stat.Exec(saldo2, penerima.Nomor)
	if error != nil {
		fmt.Println("Gagal", error.Error())
	}

	fmt.Println("Sukses Melakukan Transfer Sebesar : Rp.", transfer)

}
