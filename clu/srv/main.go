package main

import (
	"clu/srv/lib"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var jobMap map[string]*exec.Cmd

// Gets no of processes
func handleCount(w http.ResponseWriter, r *http.Request) {
	var jn uint
	for range jobMap {
		jn++
	}
	fmt.Fprintf(w, "%d", jn)
}

// Gets list of processes
func handleList(w http.ResponseWriter, r *http.Request) {
	var (
		jl   []lib.JobForm
		err  error
		name string
		cmd  *exec.Cmd
	)

	for name, cmd = range jobMap {
		jl = append(jl, lib.JobForm{
			Name: name,
			Wdir: cmd.Dir,
			Prog: cmd.Path,
			Args: cmd.Args,
			Envs: cmd.Env,
		})
	}

	if err = lib.WriteList(w, jl); err != nil {
		log.Println(err)
		return
	}
	log.Println("Send job list to client")
}

// Starts job
func handleStart(w http.ResponseWriter, r *http.Request) {
	var (
		jf  lib.JobForm
		err error
		ok  bool
		cmd *exec.Cmd
	)

	if jf, err = lib.ReadForm(r.Body); err != nil {
		log.Println(err)
		return
	}
	if _, ok = jobMap[jf.Name]; ok {
		log.Printf("Job %s already enqueued\n", jf.Name)
		return
	}

	// Create and configure command
	cmd = exec.Command(jf.Prog, jf.Args...)
	cmd.Env = jf.Envs
	cmd.Dir = jf.Wdir
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	jobMap[jf.Name] = cmd

	// Start job
	if err = jobMap[jf.Name].Start(); err != nil {
		log.Println(err)
	}
	log.Printf("Started job: %s\n", jf.Name)
}

// Stop job
func handleStop(w http.ResponseWriter, r *http.Request) {
	var (
		jf  lib.JobForm
		err error
		ok  bool
	)

	if jf, err = lib.ReadForm(r.Body); err != nil {
		log.Println(err)
		return
	}
	if _, ok = jobMap[jf.Name]; !ok {
		log.Printf("Job %s already dequeued\n", jf.Name)
		return
	}

	jobMap[jf.Name].Process.Kill()
	jobMap[jf.Name].Process.Wait()
	delete(jobMap, jf.Name)
	log.Printf("Stopped job: %s\n", jf.Name)
}

func main() {
	// Allocate jobMap
	jobMap = make(map[string]*exec.Cmd)

	// Bind routes
	http.HandleFunc(lib.PathCount, handleCount)
	http.HandleFunc(lib.PathList, handleList)
	http.HandleFunc(lib.PathStart, handleStart)
	http.HandleFunc(lib.PathStop, handleStop)
	log.Println("Initialized routes")

	go unmain()
	http.ListenAndServe(lib.Port, nil)
}

func unmain() {
	var (
		sigch = make(chan os.Signal, 1)
		cmd   *exec.Cmd
		name  string
	)

	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Catched signal %s!\n", <-sigch)

	for name, cmd = range jobMap {
		cmd.Process.Kill()
		cmd.Process.Wait()
		log.Printf("Stopped job: %s\n", name)
	}
	os.Exit(0)
}
