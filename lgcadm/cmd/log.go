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
	"github.com/TheDevtop/lgc/lgcsrv/emu"
)

const (
	CmdName_GetLog = "get-log"
	CmdName_PutLog = "put-log"
)

func GetLog_Main() int {
	var (
		res *http.Response
		err error
		buf []byte
	)

	// Assign and parse flags
	var flagHost = flag.String("h", "127.0.0.1:8080", "Specify host")
	flag.Parse()

	// Fetch response from server
	if res, err = http.Get(fmt.Sprintf(lib.UrlFormat, *flagHost, emu.RouteLog)); err != nil {
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

func PutLog_Main() int {
	var (
		err    error
		srvUrl *url.URL
		vals   url.Values
	)

	// Assign and parse flags
	var flagHost = flag.String("h", "127.0.0.1:8080", "Specify host")
	var flagMesg = flag.String("m", "Lorem ipsum...", "Specify log message")
	flag.Parse()

	// Construct url
	if srvUrl, err = url.Parse(fmt.Sprintf(lib.UrlFormat, *flagHost, emu.RouteLog)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	vals = srvUrl.Query()
	vals.Add(emu.TagModule, lib.ModName)
	srvUrl.RawQuery = vals.Encode()

	// Fetch response from server
	if _, err = http.Post(srvUrl.String(), lib.ContentType, strings.NewReader(*flagMesg)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return lib.ExitErr
	}
	return lib.ExitDef
}
