package main

import (
	"net/http"
	"strconv"
	"sync"
)

var count int
var m sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.ListenAndServe(":18080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	message := "Request URL Path is " + r.URL.Path + "\n"
	m.Lock()
	count++
	m.Unlock()
	w.Write([]byte(message))
}

func counter(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	sCount := strconv.Itoa(count)
	m.Unlock()
	message := "Counter : " + sCount + "\n"
	w.Write([]byte(message))
}