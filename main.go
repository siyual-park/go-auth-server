package main

import (
	"log"
	"net/http"

	"./core"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Ping)

	runner := core.NewServerRunner(8080, mux)

	err := runner.Run()
	if err != nil {
		log.Fatal("Listen server", err)
	}
}
