package record

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// /home/data/sqlite/RecordDB2.db
func DeleteDBinfo() {
	fmt.Println("delete DB info ")
	var dbFile string
	arch := runtime.GOOS

	if arch == "windows" {
		dbFile = "E:\\SQLiteStudio\\data\\RecordDB2.db"
		fmt.Println(arch + " " + dbFile)
	} else {
		dbFile = "/home/data/sqlite/RecordDB2.db"
		fmt.Println(arch + " " + dbFile)
	}
	//dbFile = "/mnt/e/SQLiteStudio/data/RecordDB2.db" //wsl

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Println(err)
	}
	//defer db.Close()

	rows, err := db.Query("select * from RecordDB2")
	if err != nil {
		fmt.Println("dbQuery:", err)
	}
	///	defer rows.Close()

	var (
		RecordID int

		PersoIndex    int
		PersonID      string
		RecordPicture string

		RecordTime  int
		RecordType  int
		RecordPass  int
		Temperature float64
		RecordData  string
	)

	var dataMap = make(map[int]RecordDataStruct) //RecordData字段
	var imgMap = make(map[int]string)            //RecordPicture字段 图片路径
	for rows.Next() {
		err = rows.Scan(&RecordID, &PersoIndex, &PersonID, &RecordPicture, &RecordTime, &RecordType, &RecordPass, &Temperature, &RecordData)
		if err != nil {
			log.Fatal(err)
		}
		//----------
		var RecordDataobj RecordDataStruct
		jsonData := []byte(RecordData)
		err = json.Unmarshal(jsonData, &RecordDataobj)
		if err != nil {
			fmt.Println("### Unmarshal: ", err)
		}
		dataMap[RecordID] = RecordDataobj

		imgMap[RecordID] = RecordPicture
	}

	//------ 删除陌生人的数据库ID ------
	deleteName := "陌生人"
	for id, obj := range dataMap {
		if strings.EqualFold(deleteName, obj.PersonName) {
			//执行删除
			stmt, err := db.Prepare("delete from RecordDB2 where RecordID=?")
			if err != nil {
				log.Fatal("### Delete:", err)
			}
			_, err = stmt.Exec(id)
			if err != nil {
				log.Fatal("### Exec:", err)
			} else {
				log.Printf("delete id:%v %v %v\n", RecordID, RecordTime, RecordData)
			}
		}
	}

	//------删除id对应的图片------ 删除图片之后 web页面总是弹框 没找到弹框原因在哪儿
	// for _, v := range imgMap {
	// 	//执行删除图片
	// 	err := os.Remove(v)
	// 	if err != nil {
	// 		fmt.Println("file remove Error!")
	// 		fmt.Printf("%s\n", err)
	// 	} else {
	// 		fmt.Println("file remove OK!")
	// 	}
	// }

	//
	rows.Close()
	db.Close()

}
