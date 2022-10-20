package lib

import "testing"

// Test enqueue and dequeue
func JobScheduler_Queue(t *testing.T) {
	var err error
	var desc JobDesc
	var sched = make(JobScheduler)

	desc.Name = "test"

	if err = sched.Enqueue(desc); err != nil {
		t.Fatal(err)
	}
	if err = sched.Dequeue(desc.Name); err != nil {
		t.Fatal(err)
	}
}
