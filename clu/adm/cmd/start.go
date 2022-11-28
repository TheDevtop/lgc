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
	"strconv"
	"strings"
)

// Start process
func StartMain() int {
	// Assign and parse flags
	var (
		flagName = flag.String("n", "", "Specify job name")
		flagHost = flag.String("h", "", "Specify cluster node")
		flagProg = flag.String("p", "", "Specify program path")
		flagWdir = flag.String("d", ".", "Specify working directory")
		flagArg  = flag.String("a", "", "Specify arguments")
		flagEnv  = flag.String("e", "", "Specify environment")
	)
	flag.Parse()

	// This function needs a lot of stuff
	var (
		buf  []byte
		res  *http.Response
		body *bytes.Reader
		err  error
		jf   = new(lib.JobForm)
		pool map[string]string // Map of nodes to hostnames
	)

	// Read the pool
	if pool, err = readPool(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	// Construct job form
	jf.Name = *flagName
	jf.Prog = *flagProg
	jf.Wdir = *flagWdir
	jf.Args = strings.Fields(*flagArg)
	jf.Envs = strings.Fields(*flagEnv)

	// Convert form to json buffer
	buf, _ = json.Marshal(jf)
	body = bytes.NewReader(buf)

	// If host is specified, schedule on host
	// Else schedule on pool
	if *flagHost != "" {
		// Post to cluster node and receive response
		if res, err = http.Post(fmt.Sprintf(lib.PathFormat, pool[*flagHost], lib.Port, lib.PathStart), lib.JsonMime, body); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return exitErr
		}
	} else {
		// Allocate pool schedule specific stuff
		var (
			stat     = make(map[string]int, len(pool))
			keys     = make([]string, 0, len(stat))
			srvLabel string
			srvAddr  string
			count    int
		)

		// Fetch the number of tasks of each node
		for srvLabel, srvAddr = range pool {
			if res, err = http.Get(fmt.Sprintf(lib.PathFormat, srvAddr, lib.Port, lib.PathCount)); err != nil {
				fmt.Fprintf(os.Stderr, "For %s error: %s", srvLabel, err)
				continue
			}
			if buf, err = io.ReadAll(res.Body); err != nil {
				fmt.Fprintf(os.Stderr, "For %s error: %s", srvLabel, err)
				continue
			}
			if count, err = strconv.Atoi(string(buf)); err != nil {
				fmt.Fprintf(os.Stderr, "For %s error: %s", srvLabel, err)
				continue
			}
			stat[srvAddr] = count
		}

		// Sort the nodes
		keys = sortNodes(keys, stat)

		// Post to cluster node and receive response
		if res, err = http.Post(fmt.Sprintf(lib.PathFormat, keys[0], lib.Port, lib.PathStart), lib.JsonMime, body); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return exitErr
		}

	}

	// Read response
	if buf, err = io.ReadAll(res.Body); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	fmt.Print(string(buf))
	return exitDef
}
