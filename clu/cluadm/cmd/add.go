package cmd

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/TheDevtop/lgc/clu/lib"
)

func AddMain() int {
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
		flagCmd  = flag.String("c", "", "Specify command")
		flagArg  = flag.String("a", "", "Specify arguments")
	)
	flag.Parse()

	// Construct job descriptor
	desc.Name = *flagName
	desc.Enabled = false
	desc.CmdName = *flagCmd
	desc.CmdArgs = strings.Fields(*flagArg)

	// Convert desc to json buffer
	buf, _ = json.Marshal(desc)
	body = bytes.NewReader(buf)

	// Post to cluster engine and receive response
	if resp, err = http.Post(fmt.Sprintf(lib.UrlFormat, *flagHost, lib.RouteAdd), lib.JsonMime, body); err != nil {
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
