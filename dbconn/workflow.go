package dbconn

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Data struct {
	Name  string
	Email string
}

func Workflow(ctx workflow.Context, values Data) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	var result string
	var finalResult string

	err := workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 1 failed: %v", err)
	}
	finalResult += "Activity 1: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 2 failed: %v", err)
	}
	finalResult += "Activity 2: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 3 failed: %v", err)
	}
	finalResult += "Activity 3: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 4 failed: %v", err)
	}
	finalResult += "Activity 4: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 5 failed: %v", err)
	}
	finalResult += "Activity 5: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 6 failed: %v", err)
	}
	finalResult += "Activity 6: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 7 failed: %v", err)
	}
	finalResult += "Activity 7: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 8 failed: %v", err)
	}
	finalResult += "Activity 8: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 9 failed: %v", err)
	}
	finalResult += "Activity 9: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 10 failed: %v", err)
	}
	finalResult += "Activity 10: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 11 failed: %v", err)
	}
	finalResult += "Activity 11: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 12 failed: %v", err)
	}
	finalResult += "Activity 12: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 13 failed: %v", err)
	}
	finalResult += "Activity 13: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 14 failed: %v", err)
	}
	finalResult += "Activity 14: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 15 failed: %v", err)
	}
	finalResult += "Activity 15: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 16 failed: %v", err)
	}
	finalResult += "Activity 16: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 17 failed: %v", err)
	}
	finalResult += "Activity 17: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 18 failed: %v", err)
	}
	finalResult += "Activity 18: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 19 failed: %v", err)
	}
	finalResult += "Activity 19: " + result + "\n"

	err = workflow.ExecuteActivity(ctx, Activity, values).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("Activity 20 failed: %v", err)
	}
	finalResult += "Activity 20: " + result

	return fmt.Sprintf("Workflow completed successfully:\n%s", finalResult), nil
}
