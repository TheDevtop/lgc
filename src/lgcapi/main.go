package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	certPath = "lgcapi.crt"
	keyPath  = "lgcapi.key"
	cfgPath  = "lgcapi.json"
	mgrPort  = ":8080"
	pubPort  = ":443"
)

var nsTable map[string]string

func loadConfig(path string) map[string]string {
	var (
		buf   []byte
		err   error
		table = make(map[string]string)
	)

	if buf, err = os.ReadFile(path); err != nil {
		log.Println(err)
		return table
	}
	if err = json.Unmarshal(buf, &table); err != nil {
		log.Println(err)
	}
	return table
}

func main() {
	nsTable = loadConfig(cfgPath)
	mgrMux = http.NewServeMux()
	pubMux = http.NewServeMux()

	mgrMux.HandleFunc("/publish", mgrPublish)
	mgrMux.HandleFunc("/remove", mgrRemove)
	pubMux.HandleFunc("/", pubHandle)

	go http.ListenAndServe(mgrPort, mgrMux)
	http.ListenAndServeTLS(pubPort, certPath, keyPath, pubMux)
}
