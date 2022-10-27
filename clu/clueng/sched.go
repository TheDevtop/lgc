package main

import (
	"fmt"
	"os/exec"

	"github.com/TheDevtop/lgc/clu/lib"
)

// JobScheduler data structure
type JobScheduler map[string]lib.Job

// Add job to scheduler
func (js JobScheduler) Enqueue(jd lib.JobDesc) error {
	if _, ok := js[jd.Name]; ok {
		return fmt.Errorf("error job %s already queued", jd.Name)
	}
	js[jd.Name] = lib.Job{Desc: jd}
	return nil
}

// Remove disabled job from scheduler
func (js JobScheduler) Dequeue(name string) error {
	if _, ok := js[name]; !ok {
		return fmt.Errorf("error job %s not in queue", name)
	}
	if js[name].Desc.Enabled {
		return fmt.Errorf("error job %s not disabled", name)
	}
	delete(js, name)
	return nil
}

// Start and enable job
func (js JobScheduler) Start(name string) error {
	var (
		job lib.Job
		err error
		ok  bool
	)

	if job, ok = js[name]; !ok {
		return fmt.Errorf("error job %s not in queue", name)
	}
	if job.Desc.Enabled {
		return fmt.Errorf("error job %s already enabled", name)
	}

	job.Proc = exec.Command(job.Desc.CmdName, job.Desc.CmdArgs...)
	if err = job.Proc.Start(); err != nil {
		return err
	}

	job.Desc.Enabled = true
	js[name] = job
	return nil
}

// Stop and disable job
func (js JobScheduler) Stop(name string) error {
	var (
		job lib.Job
		ok  bool
	)

	if job, ok = js[name]; !ok {
		return fmt.Errorf("error job %s not in queue", name)
	}
	if !job.Desc.Enabled {
		return fmt.Errorf("error job %s already disabled", name)
	}
	if job.Proc != nil {
		// Kill free process resources
		job.Proc.Process.Kill()
		job.Proc.Wait()
	}

	// Disable job and remove proc structure
	job.Desc.Enabled = false
	job.Proc = nil

	// Write job back
	js[name] = job
	return nil
}

// Restart job
func (js JobScheduler) Restart(name string) error {
	var (
		job lib.Job
		err error
		ok  bool
	)

	if job, ok = js[name]; !ok {
		return fmt.Errorf("error job %s not in queue", name)
	}
	if !js[name].Desc.Enabled {
		return fmt.Errorf("error job %s not enabled", name)
	}

	job.Proc = exec.Command(job.Desc.CmdName, job.Desc.CmdArgs...)
	if err = job.Proc.Start(); err != nil {
		return err
	}

	js[name] = job
	return nil
}
