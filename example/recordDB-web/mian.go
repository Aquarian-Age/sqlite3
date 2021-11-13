package main

import (
	"flag"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage imgdel -p 88989")
		flag.PrintDefaults()
	}
	p := flag.Int("p", 8989, "")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	fmt.Println("#############################")
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}
