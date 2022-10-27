package main

import (
	"fmt"
	"io"
	"net/http"
)

func apiList(w http.ResponseWriter, r *http.Request) {
	for _, str := range logBuf {
		fmt.Fprintln(w, str)
	}
}

func apiLog(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		buf []byte
	)

	if buf, err = io.ReadAll(r.Body); err != nil {
		return
	}
	if len(logBuf) <= logCap {
		logBuf = append(logBuf, string(buf))
		return
	} else {
		logBuf = make([]string, 0, logCap)
		logBuf = append(logBuf, string(buf))
		return
	}
}
