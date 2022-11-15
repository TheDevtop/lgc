package env

import (
	"fmt"
	"net/http"
)

const (
	RouteEnv = "/env"
	TagKey   = "key"
	TagValue = "val"
)

var (
	EnvMap map[string]string
)

// Manage the environment key/value store
func HandleEnv(w http.ResponseWriter, r *http.Request) {
	var (
		k = r.URL.Query().Get(TagKey)
		v = r.URL.Query().Get(TagValue)
	)

	if r.Method == http.MethodPut || r.Method == http.MethodPost {
		if v == "" {
			delete(EnvMap, k)
			return
		}
		EnvMap[k] = v
		return
	}
	if k == "" {
		for k, v = range EnvMap {
			fmt.Fprintf(w, "%s=%s\n", k, v)
		}
	} else {
		fmt.Fprintln(w, EnvMap[k])
	}
}
