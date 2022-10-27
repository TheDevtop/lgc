package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/TheDevtop/lgc/clu/lib"
)

const (
	verStr     = "LGC Cluster Engine Mk%d\n"
	verInt     = 3
	configPath = "/etc/cluster.json"
	configPerm = 0o644
)

// Global job scheduler
var jSched JobScheduler

// Load job descriptors fromt config file
func loadConfig() {
	var (
		err error
		jl  []lib.JobDesc
		buf []byte
	)

	if buf, err = os.ReadFile(configPath); err != nil {
		log.Println(err)
		return
	}
	if err = json.Unmarshal(buf, &jl); err != nil {
		log.Println(err)
		return
	}

	jSched = make(JobScheduler, len(jl))
	for _, jd := range jl {
		if err = jSched.Enqueue(jd); err != nil {
			log.Println(err)
		}
	}
}

// Store job descriptors into config file
func storeConfig() {
	var (
		err error
		jl  []lib.JobDesc
		buf []byte
	)

	for _, job := range jSched {
		jl = append(jl, job.Desc)
	}

	if buf, err = json.Marshal(jl); err != nil {
		log.Println(err)
		return
	}
	if err = os.WriteFile(configPath, buf, configPerm); err != nil {
		log.Println(err)
	}
}

func restart() {
	for _, job := range jSched {
		if job.Desc.Enabled {
			if jSched.Restart(job.Desc.Name) != nil {
				log.Printf("restart: Failed to restart %s\n", job.Desc.Name)
			}
		}
	}
}

func main() {
	// Welcome message
	log.Printf(verStr, verInt)

	// Allocate job scheduler
	jSched = make(JobScheduler)

	// Load config and restart enabled jobs
	loadConfig()
	restart()

	// Attatch api functions
	http.HandleFunc(lib.RouteAdd, apiAdd)
	http.HandleFunc(lib.RouteDelete, apiDelete)
	http.HandleFunc(lib.RouteStart, apiStart)
	http.HandleFunc(lib.RouteStop, apiStop)
	http.HandleFunc(lib.RouteList, apiList)

	// Start http server
	if err := http.ListenAndServe(lib.PortSignature, nil); err != nil {
		log.Println(err)
	}
}
