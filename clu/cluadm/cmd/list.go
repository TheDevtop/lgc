package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/TheDevtop/lgc/clu/lib"
)

func ListMain() int {
	var (
		res    *http.Response
		err    error
		buf    []byte
		jobMap = make(map[string]lib.JobDesc)
		job    lib.JobDesc
	)

	// Assign and parse flags
	var flagHost = flag.String("h", "127.0.0.1", "Specify host")
	flag.Parse()

	// Fetch response from server
	if res, err = http.Get(fmt.Sprintf(lib.UrlFormat, *flagHost, lib.RouteList)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	if buf, err = io.ReadAll(res.Body); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	if err = json.Unmarshal(buf, &jobMap); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	for _, job = range jobMap {
		fmt.Printf("Name: \"%s\" Enabled: \"%t\" Command: \"%s\"\n", job.Name, job.Enabled, job.CmdName)
	}

	return lib.ExitDef
}
