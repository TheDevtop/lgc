package main

import "net/http"

var mgrMux *http.ServeMux

func mgrPublish(w http.ResponseWriter, r *http.Request) {
	var (
		key string = r.URL.Query().Get("route")
		val string = r.URL.Query().Get("resolve")
	)
	nsTable[key] = val
}

func mgrRemove(w http.ResponseWriter, r *http.Request) {
	var key string = r.URL.Query().Get("route")
	delete(nsTable, key)
}
