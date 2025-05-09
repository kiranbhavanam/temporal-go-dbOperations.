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
}

type AppConfig struct {
	DB            DBConfig       `yaml:"database"`
	Temporal      TemporalConfig `yaml:"temporal"`
	WorkflowCount int            `yaml:"workflow_count"`
	ActivityCount int            `yaml:"activity_count"`
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


