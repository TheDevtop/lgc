package main

import (
	"fmt"
	"os"

	"github.com/TheDevtop/lgc/clu/cluadm/cmd"
)

func usage() int {
	fmt.Println("Usage: cluadm (add/del/start/stop/list) [options...]")
	return 2
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(usage())
	}
	switch os.Args[1] {
	case "add":
		os.Args = os.Args[1:]
		os.Exit(cmd.AddMain())
	case "del":
		os.Args = os.Args[1:]
		os.Exit(cmd.DeleteMain())
	case "start":
		os.Args = os.Args[1:]
		os.Exit(cmd.StartMain())
	case "stop":
		os.Args = os.Args[1:]
		os.Exit(cmd.StopMain())
	case "list":
		os.Args = os.Args[1:]
		os.Exit(cmd.ListMain())
	default:
		os.Exit(usage())
	}
}
