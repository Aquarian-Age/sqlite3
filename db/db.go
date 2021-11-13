package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// RecordData 字段
type RecordDataStruct struct {
	GauzeResult     int     `json:"GauzeResult"`
	HealthCodeColor string  `json:"HealthCodeColor"`
	ICCardNO        string  `json:"ICCardNO"`
	Latitude        int     `json:"Latitude"`
	Longitude       int     `json:"Longitude"`
	PersonName      string  `json:"PersonName"`
	Price           int     `json:"Price"`
	PriceID         string  `json:"PriceID"`
	QRCode          string  `json:"QRCode"`
	SatetyHatResult int     `json:"SatetyHatResult"`
	Similarity      float32 `json:"Similarity"`
}

// RecordDB2.db
func DBdo(t1 time.Time) {
	//var t1 time.Time = time.Date(2021, time.Month(9), 3, 0, 0, 0, 0, time.Local)
	fmt.Printf("删除 %v时间之前所有符合的数据\n", t1)

	dbFile := "E:\\SQLiteStudio\\data\\RecordDB2.db"
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select RecordID, RecordTime from RecordDB2")
	if err != nil {
		fmt.Println("dbQuery:", err)
	}

	var (
		RecordID   int
		RecordTime int64
	)

	var tMap = make(map[int]time.Time) //RecordTime
	for rows.Next() {
		err = rows.Scan(&RecordID, &RecordTime)
		if err != nil {
			log.Fatal(err)
		}
		// 1970-01-01 08:00:00 见接口文档
		coreEpoch := time.Date(1970, time.January, 1, 8, 0, 0, 0, time.UTC)
		timestamp := coreEpoch.Add(time.Duration(RecordTime * int64(time.Second))) //s
		tMap[RecordID] = timestamp
	}

	for id, t := range tMap {
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local) //精确到日
		//找出在t1时间之前符合的所有时间和id
		if t.Before(t1) {
			fmt.Printf("%d %s符合删除条件\n", id, t)
			//执行删除
			stmt, err := db.Prepare("DELETE FROM RecordDB2 where RecordID=?")
			if err != nil {
				log.Fatal("### Delete:", err)
			}
			_, err = stmt.Exec(id)
			if err != nil {
				log.Fatal("### Exec:", err)
			}
		}
	}
	//
	rows.Close()
	db.Close()

}
