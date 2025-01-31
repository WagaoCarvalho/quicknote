package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("básic server http"))
}

func main() {
	fmt.Println("Server in http://127.0.0.1:5000")

	handler := MyHandler{}

	http.ListenAndServe(":5000", handler)
}
