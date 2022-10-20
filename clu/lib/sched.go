package lib

import (
	"fmt"
	"os/exec"
)

// JobScheduler data structure
type JobScheduler map[string]Job

// Add job to scheduler
func (js JobScheduler) Enqueue(jd JobDesc) error {
	if _, ok := js[jd.Name]; ok {
		return fmt.Errorf("Error job %s already queued", jd.Name)
	}
	js[jd.Name] = Job{Desc: jd}
	return nil
}

// Remove disabled job from scheduler
func (js JobScheduler) Dequeue(name string) error {
	if _, ok := js[name]; !ok {
		return fmt.Errorf("Error job %s not in queue", name)
	}
	if js[name].Desc.Enabled {
		return fmt.Errorf("Error job %s not disabled", name)
	}
	delete(js, name)
	return nil
}

// Start and enable job
func (js JobScheduler) Start(name string) error {
	var (
		job Job
		err error
		ok  bool
	)

	if job, ok = js[name]; !ok {
		return fmt.Errorf("Error job %s not in queue", name)
	}
	if job.Desc.Enabled {
		return fmt.Errorf("Error job %s already enabled", name)
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
		job Job
		err error
		ok  bool
	)

	if job, ok = js[name]; !ok {
		return fmt.Errorf("Error job %s not in queue", name)
	}
	if !job.Desc.Enabled {
		return fmt.Errorf("Error job %s already disabled", name)
	}

	// Kill free process resources
	job.Proc.Process.Kill()
	err = job.Proc.Wait()

	// Disable job and remove proc structure
	job.Desc.Enabled = false
	job.Proc = nil

	js[name] = job
	if err != nil {
		return err
	}

	return nil
}

// Restart job
func (js JobScheduler) Restart(name string) error {
	var (
		job Job
		err error
		ok  bool
	)

	if job, ok = js[name]; !ok {
		return fmt.Errorf("Error job %s not in queue", name)
	}
	if !js[name].Desc.Enabled {
		return fmt.Errorf("Error job %s not enabled", name)
	}

	job.Proc = exec.Command(job.Desc.CmdName, job.Desc.CmdArgs...)
	if err = job.Proc.Start(); err != nil {
		return err
	}

	js[name] = job
	return nil
}
