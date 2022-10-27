package main

import (
	"net/http"
)

const (
	port = ":1106"
)

var (
	logCap = 2
	logBuf []string
)

func main() {
	logBuf = make([]string, 0, logCap)

	http.HandleFunc("/list", apiList)
	http.HandleFunc("/log", apiLog)

	http.ListenAndServe(port, nil)
}
