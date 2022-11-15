package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheDevtop/lgc/lgcsrv/api"
	"github.com/TheDevtop/lgc/lgcsrv/emu"
	"github.com/TheDevtop/lgc/lgcsrv/env"
)

var (
	mgrMux *http.ServeMux
	pubMux *http.ServeMux
	cfg    *Config
)

func main() {
	var (
		err   error
		sigch = make(chan os.Signal, 1)
	)

	// Fetch and parse configuration
	if len(os.Args) < 2 {
		log.Fatalln("Need to specify path to settings file!")
	}
	if cfg, err = loadConfig(os.Args[1]); err != nil {
		log.Fatalln(err)
	}

	// Assign and allocate
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	mgrMux = http.NewServeMux()
	pubMux = http.NewServeMux()
	env.EnvMap = cfg.EnvMap
	api.ApiTable = cfg.ApiTable
	emu.LogSize = cfg.LogSize
	log.Println("Applied configuration")

	// Configure routes
	mgrMux.HandleFunc(emu.RouteLog, emu.HandleLog)
	mgrMux.HandleFunc(emu.RoutePipe, emu.HandlePipe)
	mgrMux.HandleFunc(emu.RouteTime, emu.HandleTime)
	mgrMux.HandleFunc(env.RouteEnv, env.HandleEnv)
	mgrMux.HandleFunc(api.RouteMgr, api.HandleMgmt)
	log.Println("Configured management routes")
	pubMux.HandleFunc(api.RouteApi, api.HandleApi)
	log.Println("Configured public routes")

	// Start servers
	go func() {
		log.Println(http.ListenAndServe(cfg.MgrAddr, mgrMux))
	}()
	go func() {
		log.Println(http.ListenAndServeTLS(cfg.PubAddr, cfg.CertFile, cfg.KeyFile, pubMux))
	}()
	log.Println("Started system...")

	// Wait for stop signal
	log.Printf("Catched signal %s, stopping system...\n", <-sigch)

	// Write configuration and stop system
	cfg.EnvMap = env.EnvMap
	cfg.ApiTable = api.ApiTable

	if err = storeConfig(os.Args[1], cfg); err != nil {
		log.Fatalln(err)
	}
	log.Println("Stopped system!")
	os.Exit(0)
}
