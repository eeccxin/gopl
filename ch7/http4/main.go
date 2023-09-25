// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	//当调用 http.ListenAndServe 函数时，如果传递 nil 作为handler，它会默认使用 DefaultServeMux 作为服务器的主 handler。
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	// 表单展示
	tmpl := template.Must(template.New("list").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Inventory List</title>
		</head>
		<body>
			<table>
				<tr>
					<th>Item</th>
					<th>Price</th>
				</tr>
				{{range $item, $price := .}}
				<tr>
					<td>{{$item}}</td>
					<td>{{$price}}</td>
				</tr>
				{{end}}
			</table>
		</body>
		</html>
	`))

	err := tmpl.Execute(w, db)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 直接输出
	//for item, price := range db {
	//	fmt.Fprintf(w, "%s: %s\n", item, price)
	//}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
