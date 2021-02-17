package httphandler

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}
// Books 圖書
type Books map[string]dollars

// func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	for item, price := range db {
// 		fmt.Fprintf(w, "%s: %s\n", item, price)
// 	}
// }
type context struct{
	Request *http.Request
	
}

// List 獲取列表
func (db Books) List(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// Price 獲取價格
func (db Books) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q", item)
		return
	}
	fmt.Fprintf(w, "%s price: %s\n", item, price)
}
