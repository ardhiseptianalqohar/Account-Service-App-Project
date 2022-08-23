package entities

import "time"

type Top_UpHistory struct {
	Id, Jumlah, User_Id int
	Date                time.Time
}
