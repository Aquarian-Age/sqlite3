package main

import (
	"exsqlite/record"
	"fmt"
)

func main() {
	fmt.Println("delete unkonw person from DB")
	record.DeleteDBinfo()
}
