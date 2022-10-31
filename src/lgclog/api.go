package main

import (
	"fmt"
	"io"
	"net/http"
)

// Print the log
func apiList(w http.ResponseWriter, r *http.Request) {
	for _, str := range logBuf {
		fmt.Fprintln(w, str)
	}
}

// Send a log message to the system
func apiLog(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		buf []byte
		mod []string
		ok  bool
		msg string
	)

	if mod, ok = r.Header[modKey]; !ok || len(mod) < 1 {
		return
	}
	if buf, err = io.ReadAll(r.Body); err != nil {
		return
	}

	msg = fmt.Sprintf(msgFormat, mod[0], string(buf))

	if len(logBuf) <= logCap {
		logBuf = append(logBuf, msg)
		return
	} else {
		logBuf = make([]string, 0)
		logBuf = append(logBuf, msg)
		return
	}
}
