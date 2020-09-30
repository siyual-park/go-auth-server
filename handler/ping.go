package handler

import "net/http"

func Ping(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	res.WriteHeader(http.StatusOK)
	data := []byte("pong")
	res.Write(data)
}
