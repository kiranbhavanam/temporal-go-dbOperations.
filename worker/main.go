package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"go-poc/activity"
	"go-poc/config"
	"go-poc/workflow"
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

	// Create Temporal client
	c, err := client.Dial(client.Options{
		HostPort:  cfg.Temporal.HostPort,
		Namespace: cfg.Temporal.Namespace,
	})
	if err != nil {
		log.Fatalf("Failed to create Temporal client: %v", err)
	}
	defer c.Close()

	// Create worker with configured worker counts
	w := worker.New(c, cfg.TaskQueue, worker.Options{
		MaxConcurrentActivityExecutionSize:     cfg.Workers.ActivityWorkerCount,
		MaxConcurrentWorkflowTaskExecutionSize: cfg.Workers.WorkflowWorkerCount,
	})

	// Register workflows and activities
	w.RegisterWorkflow(workflow.Workflow)
	w.RegisterActivity(activity.Activity)

	// Start worker
	log.Printf("Starting worker on task queue: %s", cfg.TaskQueue)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
}
