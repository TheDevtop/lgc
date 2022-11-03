package main

import (
	"net/http"

	"github.com/TheDevtop/lgc/src/lgccfg/dict"
)

var kvs *dict.Dictionary

func main() {
	// Allocate new dictionary (key/value store)
	kvs = dict.NewDictionary()

	// Bind handlers
	http.HandleFunc("/create", apiCreate)
	http.HandleFunc("/read", apiRead)
	http.HandleFunc("/update", apiUpdate)
	http.HandleFunc("/delete", apiDelete)

	http.ListenAndServe(":1064", nil)

}
