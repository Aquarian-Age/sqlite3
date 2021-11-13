package main

import (
	"exsqlite/person"
	"flag"
	"fmt"
	"strconv"
	"time"
)

var t1 string

// 删除指定时间之前的数据
func main() {

	flag.StringVar(&t1, "t1", "20210101", "")
	flag.Parse()

	objMap := person.NewPersonDB()

	if len(t1) == 8 {
		ys := t1[:4]
		ms := t1[4:6]
		ds := t1[6:8]
		fmt.Println(ys, ms, ds)
		y, _ := strconv.Atoi(ys)
		m, _ := strconv.Atoi(ms)
		d, _ := strconv.Atoi(ds)
		t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
		person.DeletePersonInfo(t, objMap)
	} else {
		fmt.Println(time.Now().Local(), "输入时间格式不对 eg:20210101")
	}
	// var t1 time.Time = time.Date(2021, time.Month(9), 3, 0, 0, 0, 0, time.Local)
	// objMap := person.NewPersonDB()
	// person.DeletePersonInfo(t1, objMap)
}
