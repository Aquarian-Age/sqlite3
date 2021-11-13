package main

import (
	"exsqlite/record"
	"fmt"
	"html/template"
	"net/http"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	out := ""
	if r.Method == "POST" {
		r.ParseForm()
		for _, v := range r.Form {
			// fmt.Println("--", k, v)
			if v[0] != "" {
				fmt.Println("os: ", runtime.GOOS)
				go record.DeleteDBinfo()
				break
			}
		}
	}
	page := `<!DOCTYPE html>
	<html>
	
	<head>
		<title>删除陌生人信息</title>
		<style>
			div {
				border: 1px solid black;
				padding: 5px;
				width: 120px;
				background-color: #808080;
				margin-top: 20%;
				margin-left: auto;
				margin-right: auto;
			}
		</style>
	</head>
	
	<body bgcolor="#1a1a1a">
		  <div>
			<form action="/" method="POST">
			  <input type="text" name="cmd" style="width: 110px;height: 30px;" autofocus>
			  <input type="submit" value="键入任意键 点我" style="width: 120px;height: 120px;">
			</form>
			</div>
	</body>
	
	</html>`
	t := template.New("page")
	t, _ = t.Parse(page)
	t.Execute(w, out)
}
