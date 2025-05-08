package dbconn

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

// WorkflowParams contains parameters for the workflow
type WorkflowParams struct {
	ActivityCount int
}

// Workflow executes a configurable number of activities
func Workflow(ctx workflow.Context, params WorkflowParams) (string, error) {
	logger := workflow.GetLogger(ctx)
	workflowInfo := workflow.GetInfo(ctx)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	var finalResult string

	// Execute the configured number of activities
	for i := 1; i <= params.ActivityCount; i++ {
		activityData := Data{
			WorkflowID: workflowInfo.WorkflowExecution.ID,
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
		workflowInfo.WorkflowExecution.ID, params.ActivityCount, finalResult), nil
}
