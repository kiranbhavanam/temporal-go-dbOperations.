package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type PoolConfig struct {
	MaxConns        int           `yaml:"max_conns"`
	MinConns        int           `yaml:"min_conns"`
	MaxConnLifetime time.Duration `yaml:"max_conn_lifetime"`
	MaxConnIdleTime time.Duration `yaml:"max_conn_idle_time"`
}

type DBConfig struct {
	Host     string     `yaml:"host"`
	Port     int        `yaml:"port"`
	User     string     `yaml:"user"`
	Password string     `yaml:"password"`
	DBName   string     `yaml:"dbname"`
	SSLMode  string     `yaml:"sslmode"`
	Pool     PoolConfig `yaml:"pool"`
}

type TemporalConfig struct {
	HostPort string `yaml:"host_port"`
	Namespace string `yaml:"namespace"`
	APIKey string `yaml:"api_key"`
}

type WorkflowConfig struct {
	Count            int `yaml:"count"`
	BatchSize        int `yaml:"batchSize"`
	Parallelism      int `yaml:"parallelism"`
	ThrottleDelayMs  int `yaml:"throttleDelayMs"`
	RetryCount       int `yaml:"RetryCount"`
	ExecutionTimeout int `yaml:"executionTimeout"`
}

type WorkerConfig struct {
	WorkflowWorkerCount int `yaml:"workflowWorkerCount"`
	ActivityWorkerCount int `yaml:"activityWorkerCount"`
}

type ActivityConfig struct {
	StartToCloseTimeout    int     `yaml:"StartToCloseTimeout"`
	ScheduledToCloseTimeout int     `yaml:"ScheduledToCloseTimeout"`
	StartToScheduleTimeout int     `yaml:"StartToScheduleTimeout"`
	RetryCount            int     `yaml:"RetryCount"`
	HeartbeatIntervalSeconds int   `yaml:"heartbeatIntervalSeconds"`
	InitialInterval       int     `yaml:"initialInterval"`
	MaximumInterval       int     `yaml:"maximumInterval"`
	BackoffCoefficient    float64 `yaml:"backoffCoefficient"`
	Count                int     `yaml:"count"`
}

type AppConfig struct {
	DB            DBConfig       `yaml:"database"`
	Temporal      TemporalConfig `yaml:"temporal"`
	Workflows     WorkflowConfig `yaml:"workflows"`
	Workers       WorkerConfig   `yaml:"workers"`
	Activities    ActivityConfig `yaml:"activities"`
	TaskQueue     string         `yaml:"task_queue"`
}

// LoadConfig loads configuration from YAML file
func LoadConfig() (*AppConfig, error) {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	return &cfg, nil
}


