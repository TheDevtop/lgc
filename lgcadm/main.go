package main

import (
	"fmt"
	"os"

	"github.com/TheDevtop/lgc/lgcadm/cmd"
	"github.com/TheDevtop/lgc/lgcadm/lib"
)

func usage() int {
	fmt.Println("Usage: lgcadm [command] [options...]")
	fmt.Println("Read the manpage for more information")
	return lib.ExitErr
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(usage())
	}
	switch os.Args[1] {
	// Found in cmd/log.go
	case cmd.CmdName_GetLog:
		os.Args = os.Args[1:]
		os.Exit(cmd.GetLog_Main())
	case cmd.CmdName_PutLog:
		os.Args = os.Args[1:]
		os.Exit(cmd.PutLog_Main())
	// Found in cmd/env.go
	case cmd.CmdName_GetVar:
		os.Args = os.Args[1:]
		os.Exit(cmd.GetVar_Main())
	case cmd.CmdName_PutVar:
		os.Args = os.Args[1:]
		os.Exit(cmd.PutVar_Main())
	// Found in cmd/api.go
	case cmd.CmdName_GetRoute:
		os.Args = os.Args[1:]
		os.Exit(cmd.GetRoute_Main())
	case cmd.CmdName_PutRoute:
		os.Args = os.Args[1:]
		os.Exit(cmd.PutRoute_Main())
	default:
		os.Exit(usage())
	}
}
