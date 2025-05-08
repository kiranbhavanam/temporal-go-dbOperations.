package dbconn

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"go-poc/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	connPool *pgxpool.Pool
	connOnce sync.Once
	initErr  error
)

type Data struct {
	WorkflowID string
	ActivityID int
}

// InitDB initializes the database connection pool
func InitDB(cfg *config.DBConfig) error {
	var initErr error
	connOnce.Do(func() {
		config, err := pgxpool.ParseConfig(fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode))
		if err != nil {
			initErr = fmt.Errorf("failed to parse connection config: %v", err)
			return
		}

		// Set pool configuration
		config.MaxConns = int32(cfg.Pool.MaxConns)
		config.MinConns = int32(cfg.Pool.MinConns)
		config.MaxConnLifetime = cfg.Pool.MaxConnLifetime
		config.MaxConnIdleTime = cfg.Pool.MaxConnIdleTime

		// Create the connection pool
		connPool, initErr = pgxpool.NewWithConfig(context.Background(), config)
		if initErr != nil {
			return
		}

		// Test the connection
		err = connPool.Ping(context.Background())
		if err != nil {
			initErr = fmt.Errorf("failed to ping database: %v", err)
			return
		}
	})

	return initErr
}

// GetDB returns the connection pool
func GetDB() *pgxpool.Pool {
	return connPool
}

// CloseDB closes the connection pool
func CloseDB() {
	if connPool != nil {
		connPool.Close()
	}
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
