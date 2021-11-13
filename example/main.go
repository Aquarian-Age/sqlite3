package main

import (
	db "exsqlite/db"
	"time"
)

func main() {
	var t1 time.Time = time.Date(2021, time.Month(9), 3, 0, 0, 0, 0, time.Local)
	db.DBdo(t1)
}
