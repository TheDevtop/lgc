package cmd

import (
	"encoding/json"
	"os"
	"sort"
)

const (
	filePath   = "/etc/cluster.json"
	workingDir = "."
	exitDef    = 0
	exitErr    = 1
)

func readPool(path string) (map[string]string, error) {
	var (
		buf  []byte
		err  error
		pool map[string]string
	)

	if buf, err = os.ReadFile(path); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf, &pool); err != nil {
		return nil, err
	}
	return pool, nil
}

// Sort and reduce function
func sortNodes(keys []string, vals map[string]int) []string {
	for key := range vals {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return vals[keys[i]] < vals[keys[j]]
	})
	return keys
}
