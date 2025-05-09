package dbconn

import (
	"fmt"
	"log"
	"time"

	"go-poc/config"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// WorkflowData contains parameters for the workflow, for type safety it is better to use struct rather than sending data directly.
type WorkflowData struct {
	ID          string
	ActivityCount int
}

// Workflow executes a configurable number of activities
func Workflow(ctx workflow.Context, data WorkflowData) (string, error) {
	logger := workflow.GetLogger(ctx)
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// Configure activity options from config
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Duration(data.ActivityCount) * time.Second,
		HeartbeatTimeout: time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval: time.Duration(cfg.Activities.InitialInterval) * time.Second,
			MaximumInterval: time.Duration(cfg.Activities.MaximumInterval) * time.Millisecond,
			MaximumAttempts: int32(cfg.Activities.RetryCount),
			BackoffCoefficient: cfg.Activities.BackoffCoefficient,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	var finalResult string

	// Execute the configured number of activities
	for i := 1; i <= data.ActivityCount; i++ {
		activityData := Data{
			WorkflowID: data.ID,
			ActivityID: i,
		}

		err := workflow.ExecuteActivity(ctx, Activity, activityData).Get(ctx, &result)
		if err != nil {
			logger.Error("Activity failed", "ActivityID", i, "error", err)
			return "", fmt.Errorf("activity %d failed: %v", i, err)
		}

		finalResult += fmt.Sprintf("Activity %d: %s\n", i, result)
		logger.Info("Activity completed", "ActivityID", i, "result", result)
	}

	return fmt.Sprintf("Workflow %s completed successfully with %d activities. Results:\n%s",
		data.ID, data.ActivityCount, finalResult), nil
}
