package main

import (
	"flag"
	"net/http"
)

const (
	port      = ":1106"
	msgFormat = "[%s] %s"
	modKey    = "Module"
)

var (
	logCap *uint
	logBuf []string
)

func main() {
	// Assign and parse flags
	logCap = flag.Uint("c", 256, "Specify buffer size")
	flag.Parse()

	// Bind handlers
	http.HandleFunc("/list", apiList)
	http.HandleFunc("/log", apiLog)

	http.ListenAndServe(port, nil)
}
