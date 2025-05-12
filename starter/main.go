package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go-poc/activity"
	"go-poc/config"
	"go-poc/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	if err := activity.InitDB(&cfg.DB); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer activity.CloseDB()

	// Create Temporal client with API key
	c, err := client.Dial(client.Options{
		HostPort:  cfg.Temporal.HostPort,
		Namespace: cfg.Temporal.Namespace,
	})
	if err != nil {
		log.Fatalf("Failed to create Temporal client: %v", err)
	}
	defer c.Close()

	// Start workflows in batches
	workflowCount := cfg.Workflows.Count
	batchSize := cfg.Workflows.BatchSize
	for i := 0; i < workflowCount; i += batchSize {
		end := i + batchSize
		if end > workflowCount {
			end = workflowCount
		}

		// Create and start workflows in parallel
		var wg sync.WaitGroup
		for j := i; j < end; j++ {
			workflowID := fmt.Sprintf("workflow-%d", j)
			workflowOptions := client.StartWorkflowOptions{
				ID:                       workflowID,
				TaskQueue:                cfg.TaskQueue,
				WorkflowExecutionTimeout: time.Duration(cfg.Workflows.ExecutionTimeout) * time.Second,
				RetryPolicy: &temporal.RetryPolicy{
					InitialInterval: time.Duration(cfg.Workflows.ThrottleDelayMs) * time.Millisecond,
					MaximumAttempts: int32(cfg.Workflows.RetryCount),
				},
			}

			workflowData := workflow.WorkflowData{
				ID:            workflowID,
				ActivityCount: cfg.Activities.Count,
			}

			wg.Add(1)
			go func(id string, data workflow.WorkflowData, options client.StartWorkflowOptions) {
				defer wg.Done()
				// Start workflow
				we, err := c.ExecuteWorkflow(context.Background(), options, workflow.Workflow, data)
				if err != nil {
					log.Printf("Failed to start workflow %s: %v", id, err)
					return
				}
				log.Printf("Started workflow %s", id)

				// Process workflow result in a separate goroutine do not wait for workflow to complete each workflow is processed on separate go routine.
				go func(we client.WorkflowRun) {
					var result string
					if err := we.Get(context.Background(), &result); err != nil {
						log.Printf("Workflow %s failed: %v", we.GetID(), err)
						return
					}
					log.Printf("Workflow %s completed: %s", we.GetID(), result)
				}(we)
			}(workflowID, workflowData, workflowOptions)
		}

		// Wait for all workflows in this batch to start
		wg.Wait()

		// Add delay between batches
		time.Sleep(time.Duration(cfg.Workflows.ThrottleDelayMs) * time.Millisecond)
	}

	// Keep the main goroutine alive
	select {}
}
