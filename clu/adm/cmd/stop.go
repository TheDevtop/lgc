package cmd

import (
	"bytes"
	"clu/adm/lib"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Stop and remove job from node
func StopMain() int {
	// Assign and parse flags
	var (
		flagName = flag.String("n", "job", "Specify job name")
		flagHost = flag.String("h", "localhost", "Specify cluster node")
	)
	flag.Parse()

	// This function needs a lot of stuff
	var (
		buf  []byte
		res  *http.Response
		body *bytes.Reader
		err  error
		jf   = new(lib.JobForm)
		pool map[string]string
	)

	if pool, err = readPool(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	// Convert form to json buffer
	jf.Name = *flagName
	buf, _ = json.Marshal(jf)
	body = bytes.NewReader(buf)

	// Post to cluster node and receive response
	if res, err = http.Post(fmt.Sprintf(lib.PathFormat, pool[*flagHost], lib.Port, lib.PathStop), lib.JsonMime, body); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}
	if buf, err = io.ReadAll(res.Body); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	fmt.Println(string(buf))
	return exitDef
}
