package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

var pubMux *http.ServeMux

func pubHandle(w http.ResponseWriter, r *http.Request) {
	var (
		val string
		ok  bool
		err error
		px  *url.URL
		pr  *http.Response
	)

	if val, ok = nsTable[r.URL.Path]; !ok {
		http.NotFound(w, r)
		return
	}
	if px, err = url.Parse(val); err != nil {
		log.Println(err)
		return
	}

	r.Host = px.Host
	r.URL.Host = px.Host
	r.URL.Path = px.Path
	r.URL.Scheme = px.Scheme
	r.RequestURI = ""

	if pr, err = http.DefaultClient.Do(r); err != nil {
		log.Println(err)
		return
	}
	io.Copy(w, pr.Body)
}
