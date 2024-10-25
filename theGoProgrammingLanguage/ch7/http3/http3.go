package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// /create?item=hat&price=10
// func (db database) create(w http.ResponseWriter, req *http.Request) {
// todo
// }

// /update?item=socks&price=6
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(404)
		fmt.Fprintf(w, "item or price empty")
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not expected price:%v\n", err)
		return
	}
	if _, ok := db[item]; !ok {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s does not exist in db", item)
		return
	}
	db[item] = dollars(p)
	fmt.Fprintf(w, "update %s: %s/n", item, price)

}

func main() {

	db := database{"shoes": 50, "socks": 5}
	// mux := http.NewServeMux()
	// // mux.Handle("/list", http.HandlerFunc(db.list))
	// // mux.Handle("/price", http.HandlerFunc(db.price))
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.list)

	// http.ListenAndServe("localhost:8888", mux)
	//http.DefaultServeMux
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.ListenAndServe("localhost:8888", nil)
}
