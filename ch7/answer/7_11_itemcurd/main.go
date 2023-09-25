package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

/*
在这个例子中，我们创建了一个名为 Inventory 的结构体，其中包含一个互斥锁 mu 和一个记录数据库记录的映射 records。我们为 Inventory 结构体实现了 ServeHTTP 方法，以便它可以作为一个处理程序来处理不同的请求。

根据请求的路径，我们在 ServeHTTP 方法中调用相应的处理程序。createHandler 处理创建记录的请求，readHandler 处理读取记录的请求，updateHandler 处理更新记录的请求，deleteHandler 处理删除记录的请求。

在每个处理程序中，我们从请求中获取相应的参数，并根据需要进行验证和操作。如果出现错误，我们使用 http.Error 函数返回相应的错误响应。

在 main 函数中，我们创建了一个 Inventory 实例，并将其作为处理程序注册到根路径。然后，我们使用 http.ListenAndServe 函数启动服务器，监听端口8080。

通过发送不同的请求，如 /create?item=socks&price=6 来创建记录，/read?item=socks 来读取记录，/update?item=socks&price=8 来更新记录，/delete?item=socks 来删除记录，我们可以测试这些处理程序的功能。
*/

type Inventory struct {
	mu      sync.Mutex //加锁
	records map[string]float64
}

func (inv *Inventory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/create":
		inv.createHandler(w, r)
	case "/read":
		inv.readHandler(w, r)
	case "/update":
		inv.updateHandler(w, r)
	case "/delete":
		inv.deleteHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (inv *Inventory) createHandler(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	priceStr := r.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	inv.mu.Lock()
	defer inv.mu.Unlock()

	if _, ok := inv.records[item]; ok {
		http.Error(w, "Item already exists", http.StatusBadRequest)
		return
	}

	inv.records[item] = price
	fmt.Fprintf(w, "Item created: %s, Price: %.2f\n", item, price)
}

func (inv *Inventory) readHandler(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")

	inv.mu.Lock()
	defer inv.mu.Unlock()

	price, ok := inv.records[item]
	if !ok {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Item: %s, Price: %.2f\n", item, price)
}

func (inv *Inventory) updateHandler(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	priceStr := r.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	inv.mu.Lock()
	defer inv.mu.Unlock()

	if _, ok := inv.records[item]; !ok {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	inv.records[item] = price
	fmt.Fprintf(w, "Item updated: %s, New price: %.2f\n", item, price)
}

func (inv *Inventory) deleteHandler(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")

	inv.mu.Lock()
	defer inv.mu.Unlock()

	if _, ok := inv.records[item]; !ok {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	delete(inv.records, item)
	fmt.Fprintf(w, "Item deleted: %s\n", item)
}

func main() {
	inv := &Inventory{
		records: make(map[string]float64),
	}

	http.Handle("/", inv)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
