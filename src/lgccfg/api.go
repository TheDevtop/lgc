package main

import (
	"fmt"
	"net/http"
)

func apiCreate(w http.ResponseWriter, r *http.Request) {
	var (
		key string = r.URL.Query().Get("key")
		val string = r.URL.Query().Get("val")
	)

	kvs.Create(key, val)
}

func apiRead(w http.ResponseWriter, r *http.Request) {
	var (
		key string = r.URL.Query().Get("key")
		val string
		ok  bool
	)

	if val, ok = kvs.Read(key); !ok {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "%s", val)
}

func apiUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		key string = r.URL.Query().Get("key")
		val string = r.URL.Query().Get("val")
	)

	kvs.Update(key, val)
}

func apiDelete(w http.ResponseWriter, r *http.Request) {
	var (
		key string = r.URL.Query().Get("key")
		val string = r.URL.Query().Get("val")
	)

	kvs.Delete(key, val)
}
