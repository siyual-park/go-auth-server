package main

import (
	"fmt"
	"log"
	"net/http"
)

func Ping(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	res.WriteHeader(http.StatusOK)
	data := []byte("pong")
	res.Write(data)
}

func main() {
	port := 8080

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Ping)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal("Listen server", err)
	}
}
