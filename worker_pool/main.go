package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Output int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		output := job.Value * job.Value // simulate CPU work
		fmt.Printf("worker %d → job %d\n", id, job.ID)
		results <- Result{JobID: job.ID, Output: output}
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j}
	}
	close(jobs) // signal workers there are no more jobs

	// Close results once all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("job %d → %d\n", r.JobID, r.Output)
	}
}
