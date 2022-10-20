package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
)

type JobDesc struct {
	Name    string
	Enabled bool
	CmdName string
	CmdArgs []string
}

type Job struct {
	Desc JobDesc
	Proc *exec.Cmd
}

// Read and convert json to JobDesc
func ReadJobDesc(rc io.ReadCloser) (JobDesc, error) {
	var (
		buf []byte
		err error
		jd  = new(JobDesc)
	)

	buf, _ = io.ReadAll(rc)
	err = json.Unmarshal(buf, jd)
	return *jd, err
}

// Write JobDesc to json
func WriteJobDesc(w io.Writer, jdPtr *JobDesc) error {
	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(jdPtr); err != nil {
		return err
	}
	fmt.Fprint(w, string(buf))
	return nil
}
