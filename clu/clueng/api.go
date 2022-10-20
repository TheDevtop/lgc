package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TheDevtop/lgc/clu/lib"
)

func apiAdd(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		jd  lib.JobDesc
	)

	if jd, err = lib.ReadJobDesc(r.Body); err != nil {
		log.Printf("apiAdd: %s\n", err)
		return
	}
	if err = jSched.Enqueue(jd); err != nil {
		log.Printf("apiAdd: %s\n", err)
		return
	}

	storeConfig()
	log.Printf("apiAdd: Job %s queued\n", jd.Name)
}

func apiDelete(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		jd  lib.JobDesc
	)

	if jd, err = lib.ReadJobDesc(r.Body); err != nil {
		log.Printf("apiDelete: %s\n", err)
		return
	}
	if err = jSched.Dequeue(jd.Name); err != nil {
		log.Printf("apiDelete: %s\n", err)
		return
	}

	storeConfig()
	log.Printf("apiDelete: Job %s removed from queue\n", jd.Name)
}

func apiStart(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		jd  lib.JobDesc
	)

	if jd, err = lib.ReadJobDesc(r.Body); err != nil {
		log.Printf("apiStart: %s\n", err)
		return
	}
	if err = jSched.Start(jd.Name); err != nil {
		log.Printf("apiStart: %s\n", err)
		return
	}

	storeConfig()
	log.Printf("apiStart: Job %s started\n", jd.Name)
}

func apiStop(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		jd  lib.JobDesc
	)

	if jd, err = lib.ReadJobDesc(r.Body); err != nil {
		log.Printf("apiStop: %s\n", err)
		return
	}
	if err = jSched.Stop(jd.Name); err != nil {
		log.Printf("apiStop: %s\n", err)
		return
	}

	storeConfig()
	log.Printf("apiStop: Job %s stopped\n", jd.Name)
}

func apiList(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		jm  = make(map[string]lib.JobDesc, len(jSched))
		job lib.Job
		buf []byte
	)

	for _, job = range jSched {
		jm[job.Desc.Name] = job.Desc
	}
	if buf, err = json.Marshal(jm); err != nil {
		log.Printf("apiList: %s\n", err)
		return
	}

	fmt.Fprint(w, string(buf))
	log.Println("apiList: Wrote list to client")
}
