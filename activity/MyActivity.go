package activity
import (
	"context"
	"fmt"
	"strconv"
	"time"
)
type Data struct {
	WorkflowID string
	ActivityID int
}

// Activity performs a database insert operation
func Activity(ctx context.Context, data Data) (string, error) {
	if connPool == nil {
		return "", fmt.Errorf("database not initialized")
	}

	query := `INSERT INTO workflow_activities (workflow_id, activity_id, status, created_at) 
		VALUES ($1, $2, $3, $4)`

	// Get a connection from the pool
	conn, err := connPool.Acquire(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to acquire connection: %v", err)
	}
	defer conn.Release()

	// Execute the query
	_, err = conn.Exec(ctx, query, data.WorkflowID, strconv.Itoa(data.ActivityID), "COMPLETED", time.Now().Format(time.RFC3339))
	if err != nil {
		return "", fmt.Errorf("failed to insert activity: %v", err)
	}

	return fmt.Sprintf("Activity %d completed for workflow %s", data.ActivityID, data.WorkflowID), nil
}