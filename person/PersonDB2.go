package person

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// PersonDB表的字段
type PersonDBStruct struct {
	PersonIndex   int64     `json:"person_index"`
	PersonID      string    `json:"person_id"`
	PersonName    string    `json:"person_name"`
	PersonPicture string    `json:"person_picture"`
	CreateTime    time.Time `json:"create_time"`
}

// PersonDB2.db 返回值 k:PersonIndex v:PersonDBStruct结构体指针
func NewPersonDB() map[int64]*PersonDBStruct {
	var dbFile string
	arch := runtime.GOOS
	if arch == "windows" {
		dbFile = "E:\\SQLiteStudio\\data\\persondb-test\\PersonDB2.db"
		fmt.Println(arch + " " + dbFile)
	} else {
		dbFile = "/home/data/sqlite/PersonDB2.db"
		fmt.Println(arch + " " + dbFile)
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select  PersonIndex, PersonID , PersonName , PersonPicture , CreateTime  from PersonDB")
	if err != nil {
		fmt.Println("dbQuery:", err)
	}

	var (
		personIndex   int64
		personID      string
		PersonName    string
		personPicture string
		CreateTime    int64
	)

	objMap := make(map[int64]*PersonDBStruct) //k:PersonIndex v:结构体
	for rows.Next() {
		err = rows.Scan(&personIndex, &personID, &PersonName, &personPicture, &CreateTime)
		if err != nil {
			log.Fatal(err)
		}

		//-----时间处理-----
		//1970-01-01 08:00:00 见接口文档
		coreEpoch := time.Date(1970, time.January, 1, 8, 0, 0, 0, time.UTC)
		timestamp := coreEpoch.Add(time.Duration(CreateTime * int64(time.Second))) //s

		//-----图片路径处理-----
		picturePath, pictureFile := path.Split(personPicture)
		picturePath = strings.Trim(picturePath, "//") //去掉“//”
		personPicture = path.Join(picturePath, pictureFile)
		personPicture = "/" + personPicture //绝对路径

		//-----structMap-----
		objMap[personIndex] = &PersonDBStruct{
			PersonIndex:   personIndex,
			PersonID:      personID,
			PersonName:    PersonName,
			PersonPicture: personPicture,
			CreateTime:    timestamp,
		}

	}
	//
	rows.Close()
	db.Close()

	return objMap
}

//删除t1时间点之前的所有人员信息  精确到日
func DeletePersonInfo(t1 time.Time, objMap map[int64]*PersonDBStruct) {
	fmt.Printf("删除 %v时间之前的数据\n", t1)
	// dbFile := "E:\\SQLiteStudio\\data\\persondb-test\\PersonDB2.db"

	var dbFile string
	arch := runtime.GOOS
	if arch == "windows" {
		dbFile = "E:\\SQLiteStudio\\data\\persondb-test\\PersonDB2.db"
		fmt.Println(arch + " " + dbFile)
	} else {
		dbFile = "/home/data/sqlite/PersonDB2.db"
		fmt.Println(arch + " " + dbFile)
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Println(err)
	}
	for pid, obj := range objMap {
		t := obj.CreateTime
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local) //精确到日
		if t.Before(t1) {
			stmt, err := db.Prepare("DELETE FROM PersonDB where PersonID=?")
			if err != nil {
				log.Fatal("### Delete:", err)
			} else {
				log.Printf("delete %s from PersonDB\n", obj.PersonID)
			}

			if pid == (obj.PersonIndex) {
				//删除数据库内容
				_, err = stmt.Exec(obj.PersonID)
				if err != nil {
					log.Fatal("### Exec:", err)
				}
				//删除图片
				err = os.Remove(obj.PersonPicture)
				if err != nil {
					log.Println("### delete img", err)
				} else {
					fmt.Printf("remove %s OK\n", obj.PersonPicture)
				}
			}

		}
	}
}
