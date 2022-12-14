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
		flagFile = flag.String("f", filePath, "Specify configuration file")
		flagName = flag.String("n", "", "Specify job name")
		flagHost = flag.String("h", "", "Specify cluster node")
	)
	flag.Parse()

	var (
		buf  []byte
		res  *http.Response
		body *bytes.Reader
		err  error
		jf   = new(lib.JobForm)
		pool map[string]string
	)

	// Read the pool
	if pool, err = readPool(*flagFile); err != nil {
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

	fmt.Print(string(buf))
	return exitDef
}
