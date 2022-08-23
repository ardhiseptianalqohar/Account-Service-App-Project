package entities

import "time"

type TransferHistory struct {
	Id, Jumlah, Send_Id, Accept_Id int
	Date                           time.Time
}
