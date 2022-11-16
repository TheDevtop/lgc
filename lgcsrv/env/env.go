package env

import (
	"fmt"
	"log"
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
			log.Printf("Removed key \"%s\"\n", k)
			return
		}
		EnvMap[k] = v
		log.Printf("Added key \"%s\" with value \"%s\"\n", k, v)
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
