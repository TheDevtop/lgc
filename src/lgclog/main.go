package main

import (
	"net/http"
)

const (
	port      = ":1106"
	msgFormat = "[%s] %s"
	modKey    = "Module"
)

var (
	logCap = 256
	logBuf []string
)

func main() {
	http.HandleFunc("/list", apiList)
	http.HandleFunc("/log", apiLog)

	http.ListenAndServe(port, nil)
}
