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
	"github.com/TheDevtop/lgc/lgcsrv/api"
)

const (
	CmdName_GetRoute = "get-route"
	CmdName_PutRoute = "put-route"
)

func GetRoute_Main() int {
	var (
		res    *http.Response
		err    error
		buf    []byte
		srvUrl *url.URL
	)

	// Assign and parse flags
	var flagHost = flag.String("h", "127.0.0.1:8080", "Specify host")
	flag.Parse()

	// Construct url
	if srvUrl, err = url.Parse(fmt.Sprintf(lib.UrlFormat, *flagHost, api.RouteMgr)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}

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

func PutRoute_Main() int {
	var (
		err    error
		srvUrl *url.URL
		req    *http.Request
	)

	// Assign and parse flags
	var flagHost = flag.String("h", "127.0.0.1:8080", "Specify host")
	var flagRoute = flag.String("s", "/foo", "Specify source route")
	var flagDest = flag.String("d", "", "Specify destination route")
	flag.Parse()

	// Construct url
	if srvUrl, err = url.Parse(fmt.Sprintf(lib.UrlFormat, *flagHost, api.RouteMgr)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}

	// Construct request
	if req, err = http.NewRequest(http.MethodPost, srvUrl.String(), strings.NewReader("")); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	req.Header.Add(api.TagRoute, *flagRoute)
	req.Header.Add(api.TagDestination, *flagDest)

	// Fetch response from server
	if _, err = http.DefaultClient.Do(req); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	return lib.ExitDef
}
