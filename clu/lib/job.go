package lib

import (
	"encoding/json"
	"fmt"
	"io"
)

type JobForm struct {
	Name string
	Wdir string
	Prog string
	Args []string
	Envs []string
}

func ReadForm(rc io.ReadCloser) (JobForm, error) {
	var jf = new(JobForm)
	buf, _ := io.ReadAll(rc)
	err := json.Unmarshal(buf, jf)
	return *jf, err
}

func WriteForm(w io.Writer, jf *JobForm) error {
	if buf, err := json.Marshal(jf); err != nil {
		return err
	} else {
		fmt.Fprint(w, string(buf))
		return nil
	}
}

func ReadList(rc io.ReadCloser) ([]JobForm, error) {
	var jl = make([]JobForm, 2)
	buf, _ := io.ReadAll(rc)
	err := json.Unmarshal(buf, &jl)
	return jl, err
}

func WriteList(w io.Writer, jl []JobForm) error {
	if buf, err := json.Marshal(&jl); err != nil {
		return err
	} else {
		fmt.Fprint(w, string(buf))
		return nil
	}
}
