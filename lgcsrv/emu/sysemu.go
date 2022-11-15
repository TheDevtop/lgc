package emu

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	RouteLog  = "/sys/log"
	RoutePipe = "/sys/pipe"
	RouteTime = "/sys/time"

	TagModule = "module"
)

var (
	LogSize int
	pipeBuf []byte
	logBuf  []string
)

// Manage the system logger
func HandleLog(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut || r.Method == http.MethodPost {
		if len(logBuf) > LogSize {
			logBuf = nil
		}
		buf, _ := io.ReadAll(r.Body)
		mod := r.URL.Query().Get(TagModule)
		logBuf = append(logBuf, fmt.Sprintf("[%s] %s", mod, buf))
		return
	}
	for _, v := range logBuf {
		fmt.Fprintln(w, v)
	}
}

// Manage the system pipe
func HandlePipe(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == http.MethodPut || r.Method == http.MethodPost {
		if pipeBuf, err = io.ReadAll(r.Body); err != nil {
			log.Println(err)
			pipeBuf = []byte("")
		}
		return
	}
	fmt.Fprintln(w, string(pipeBuf))
}

// Print the current time
func HandleTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().String())
}
