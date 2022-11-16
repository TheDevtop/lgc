package main

import (
	"fmt"
	"os"

	"github.com/TheDevtop/lgc/lgcadm/cmd"
	"github.com/TheDevtop/lgc/lgcadm/lib"
)

func usage() int {
	fmt.Println("Usage: lgcadm [command] [options...]")
	return lib.ExitErr
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(usage())
	}
	switch os.Args[1] {
	case cmd.CmdName_GetLog:
		os.Args = os.Args[1:]
		os.Exit(cmd.GetLog_Main())
	case cmd.CmdName_PutLog:
		os.Args = os.Args[1:]
		os.Exit(cmd.PutLog_Main())
	default:
		os.Exit(usage())
	}
}
