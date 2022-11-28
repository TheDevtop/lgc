package cmd

import (
	"clu/adm/lib"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Print list of jobs of all the nodes
func ListMain() int {
	var (
		buf      []byte
		jl       []lib.JobForm
		jf       lib.JobForm
		res      *http.Response
		err      error
		pool     map[string]string
		srvLabel string
		srvAddr  string
	)

	if pool, err = readPool(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	for srvLabel, srvAddr = range pool {
		if res, err = http.Get(fmt.Sprintf(lib.PathFormat, srvAddr, lib.Port, lib.PathList)); err != nil {
			fmt.Fprintf(os.Stderr, "For %s error: %s", srvLabel, err)
			continue
		}
		if buf, err = io.ReadAll(res.Body); err != nil {
			fmt.Fprintf(os.Stderr, "For %s error: %s", srvLabel, err)
			continue
		}
		if err = json.Unmarshal(buf, &jl); err != nil {
			fmt.Fprintf(os.Stderr, "For %s error: %s", srvLabel, err)
			continue
		}
		for _, jf = range jl {
			fmt.Printf("Name: %s/%s Dir: %s\nProgram: %s Args: %v\n", srvLabel, jf.Name, jf.Wdir, jf.Prog, jf.Args)
		}
	}
	return exitDef
}
