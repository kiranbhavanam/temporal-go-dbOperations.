package activity

import (
	"context"
	"fmt"
	"sync"

	"go-poc/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	connPool *pgxpool.Pool //It's like a container that holds multiple db connections.
	connOnce sync.Once //ensures no matter how many times InitDB is called, it will only initialize once.
	initErr  error
)

// InitDB initializes the database connection pool
func InitDB(cfg *config.DBConfig) error {
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


// CloseDB closes the connection pool
func CloseDB() {
	if connPool != nil {
		connPool.Close()
	}
}

