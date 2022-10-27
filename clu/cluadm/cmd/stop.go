package cmd

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/TheDevtop/lgc/clu/lib"
)

func StopMain() int {
	var (
		err  error
		desc = new(lib.JobDesc)
		buf  []byte
		body *bytes.Reader
		resp *http.Response
	)

	// Specify and parse flags
	var (
		flagHost = flag.String("h", "127.0.0.1", "Specify host")
		flagName = flag.String("n", "", "Specify job name")
	)
	flag.Parse()

	// Construct job descriptor
	desc.Name = *flagName
	desc.Enabled = false

	// Convert desc to json buffer
	buf, _ = json.Marshal(desc)
	body = bytes.NewReader(buf)

	// Post to cluster engine and receive response
	if resp, err = http.Post(fmt.Sprintf(lib.UrlFormat, *flagHost, lib.RouteStop), lib.JsonMime, body); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	if buf, err = io.ReadAll(resp.Body); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}

	fmt.Println(string(buf))
	return lib.ExitDef
}
