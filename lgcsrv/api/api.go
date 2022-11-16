package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	RouteApi       = "/"
	RouteMgr       = "/api"
	TagRoute       = "route"
	TagDestination = "dest"
)

var (
	ApiTable map[string]string
)

// Manage the ApiTable
func HandleMgmt(w http.ResponseWriter, r *http.Request) {
	var (
		src = r.Header.Get(TagRoute)
		dst = r.Header.Get(TagDestination)
		err error
	)

	if r.Method == http.MethodPut || r.Method == http.MethodPost {
		if dst == "" {
			delete(ApiTable, src)
			log.Printf("Removed route \"%s\"\n", src)
			return
		}
		if _, err = url.Parse(dst); err != nil {
			log.Println(err)
			return
		}
		ApiTable[src] = dst
		log.Printf("Mapped route \"%s\" to \"%s\"\n", src, dst)
		return
	}
	for src, dst = range ApiTable {
		fmt.Fprintf(w, "%s => %s\n", src, dst)
	}
}

// Handle the incoming requests
func HandleApi(w http.ResponseWriter, r *http.Request) {
	var (
		route string
		ok    bool
		err   error
		pu    *url.URL
		pr    *http.Response
	)

	if route, ok = ApiTable[r.URL.Path]; !ok {
		http.NotFound(w, r)
		return
	}
	if pu, err = url.Parse(route); err != nil {
		log.Println(err)
		return
	}

	r.Host = pu.Host
	r.URL.Host = pu.Host
	r.URL.Path = pu.Path
	r.URL.Scheme = pu.Scheme
	//r.URL.RawQuery = pu.RawQuery
	r.RequestURI = ""

	if pr, err = http.DefaultClient.Do(r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Copy(w, pr.Body)
}
