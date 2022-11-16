package cmd

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/TheDevtop/lgc/lgcadm/lib"
	"github.com/TheDevtop/lgc/lgcsrv/env"
)

const (
	CmdName_GetVar = "get-var"
	CmdName_PutVar = "put-var"
)

func GetVar_Main() int {
	var (
		res    *http.Response
		err    error
		buf    []byte
		srvUrl *url.URL
		vals   url.Values
	)

	// Assign and parse flags
	var flagHost = flag.String("h", "127.0.0.1:8080", "Specify host")
	var flagKey = flag.String("k", "", "Specify key")
	flag.Parse()

	// Construct url
	if srvUrl, err = url.Parse(fmt.Sprintf(lib.UrlFormat, *flagHost, env.RouteEnv)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	vals = srvUrl.Query()
	vals.Add(env.TagKey, *flagKey)
	srvUrl.RawQuery = vals.Encode()

	// Fetch response from server
	if res, err = http.Get(srvUrl.String()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	if buf, err = io.ReadAll(res.Body); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	fmt.Print(string(buf))
	return lib.ExitDef
}

func PutVar_Main() int {
	var (
		err    error
		srvUrl *url.URL
		vals   url.Values
	)

	// Assign and parse flags
	var flagHost = flag.String("h", "127.0.0.1:8080", "Specify host")
	var flagKey = flag.String("k", "", "Specify key")
	var flagVal = flag.String("v", "", "Specify value")
	flag.Parse()

	// Construct url
	if srvUrl, err = url.Parse(fmt.Sprintf(lib.UrlFormat, *flagHost, env.RouteEnv)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	vals = srvUrl.Query()
	vals.Add(env.TagKey, *flagKey)
	vals.Add(env.TagValue, *flagVal)
	srvUrl.RawQuery = vals.Encode()

	// Fetch response from server
	if _, err = http.Post(srvUrl.String(), lib.ContentType, strings.NewReader("")); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	return lib.ExitDef
}
