package main

import (
	"context"
	_ "embed"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"sync"
	"time"
)

//go:embed job_template.yaml
var jobTemplate string

const (
	charset     = "abcdefghijklmnopqrstuvwxyz0123456789"
	totalJobs   = 15000
	concurrency = 64
)

var (
	successCount = 0
	failureCount = 0
	mu           sync.Mutex
)

func generateRandomName() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func createJob(wg *sync.WaitGroup, jobChan chan string) {
	defer wg.Done()
	for jobName := range jobChan {
		jobYAML := strings.Replace(jobTemplate, "stress-test-job", jobName, 1)
		cmd := exec.Command("kubectl", "apply", "-f", "-")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			mu.Lock()
			failureCount++
			mu.Unlock()
			continue
		}

		go func() {
			defer stdin.Close()
			stdin.Write([]byte(jobYAML))
		}()

		if err := cmd.Run(); err != nil {
			mu.Lock()
			failureCount++
			mu.Unlock()
		} else {
			mu.Lock()
			successCount++
			mu.Unlock()
		}
	}
}

func monitorJobs(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Monitoring stopped")
			return
		default:
			fmt.Printf("Success: %d, Failure: %d\r", successCount, failureCount)
			time.Sleep(1 * time.Second)
		}
	}
}

func cleanupJobs() {
	cmd := exec.Command("kubectl", "delete", "jobs", "-l", "app=stress-test")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error cleaning up jobs: %v\n", err)
	} else {
		fmt.Println("Successfully cleaned up jobs")
	}
}

func main() {
	startTime := time.Now()
	jobChan := make(chan string, totalJobs)
	var wg *sync.WaitGroup = new(sync.WaitGroup)

	ctx, ctxCancel := context.WithCancel(context.Background())
	// Start monitoring
	go monitorJobs(ctx)

	// Start worker goroutines
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go createJob(wg, jobChan)
	}

	// Generate job names
	for i := 0; i < totalJobs; i++ {
		jobChan <- "stress-job-" + generateRandomName()
	}
	close(jobChan)

	wg.Wait()

	ctxCancel()

	fmt.Printf("Job creation completed in %v\n", time.Since(startTime))
	fmt.Printf("Success: %d, Failure: %d\n", successCount, failureCount)

	// Wait for all jobs to complete
	// time.Sleep(10 * time.Second) // Give monitor time to detect completion
	// cleanupJobs()
}
