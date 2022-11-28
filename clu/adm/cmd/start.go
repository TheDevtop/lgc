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

// Schedule based on the least busy node
func StartMain() int {
	// Assign and parse flags
	var (
		flagName = flag.String("n", "job", "Specify job name")
		flagProg = flag.String("p", "/usr/bin/true", "Specify program path")
		flagWdir = flag.String("d", ".", "Specify working directory")
		flagArg  = flag.String("a", "", "Specify arguments")
	)
	flag.Parse()

	// This function needs a lot of stuff
	var (
		buf      []byte
		res      *http.Response
		keys     []string
		body     *bytes.Reader
		err      error
		jf       = new(lib.JobForm)
		pool     map[string]string // Map of nodes to hostnames
		stat     map[string]int    // Map of hostnames to jobcount
		srvLabel string            // Label or nodename
		srvAddr  string            // Address or hostname
	)

	// Read the pool and allocate based on the pool size
	if pool, err = readPool(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}
	stat = make(map[string]int, len(pool))
	keys = make([]string, 0, len(stat))

	// Fetch the number of tasks of each node
	for srvLabel, srvAddr = range pool {
		var (
			count int
		)

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

	// Construct job form
	jf.Name = *flagName
	jf.Prog = *flagProg
	jf.Wdir = *flagWdir
	jf.Args = strings.Fields(*flagArg)

	// Convert form to json buffer
	buf, _ = json.Marshal(jf)
	body = bytes.NewReader(buf)

	// Post to cluster node and receive response
	if res, err = http.Post(fmt.Sprintf(lib.PathFormat, keys[0], lib.Port, lib.PathStart), lib.JsonMime, body); err != nil {
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
