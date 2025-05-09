package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-poc/config"
	"go-poc/dbconn"

	"go.temporal.io/sdk/client"
)


func main() {
	// Load configuration 
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	if err := dbconn.InitDB(&cfg.DB); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbconn.CloseDB()

	// Create Temporal client
	c, err := client.Dial(client.Options{
		HostPort:  cfg.Temporal.HostPort,
		Namespace: cfg.Temporal.Namespace,
	})
	if err != nil {
		log.Fatalf("Failed to create Temporal client: %v", err)
	}
	defer c.Close()

	// Start workflow for each workflow count
	for i := 0; i < cfg.WorkflowCount; i++ {
		workflowID := fmt.Sprintf("workflow-%d-%d", time.Now().UnixNano(), i)
		workflowOptions := client.StartWorkflowOptions{
			ID:        workflowID,
			TaskQueue: cfg.TaskQueue,
		}

		workflowParams := dbconn.WorkflowParams{
			ActivityCount: cfg.ActivityCount,
		}

		we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, dbconn.Workflow, workflowParams)
		if err != nil {
			log.Printf("Failed to start workflow %d: %v", i, err)
			continue
		}

		log.Printf("Started workflow %d: ID=%s, RunID=%s", i, we.GetID(), we.GetRunID())

		// Process workflow result in a separate goroutine do not wait for workflow to complete each workflow is processed on separate go routine.
		go func(we client.WorkflowRun) {
			var result string
			if err := we.Get(context.Background(), &result); err != nil {
				log.Printf("Workflow %s failed: %v", we.GetID(), err)
				return
			}
			log.Printf("Workflow %s completed: %s", we.GetID(), result)
		}(we)
	}

	// Keep the main goroutine alive
	select {}
}