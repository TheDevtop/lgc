package cmd

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/TheDevtop/lgc/clu/lib"
)

func ListMain() int {
	var (
		//res *http.Response
		err error
	)

	// Assign and parse flags
	var flagHost = flag.String("h", "127.0.0.1", "Specify host")
	flag.Parse()

	// Fetch response from server
	if _, err = http.Get(fmt.Sprintf(urlFormat, *flagHost, lib.RouteList)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	return exitDef
}
