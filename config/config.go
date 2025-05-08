package config

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type TemporalConfig struct {
	HostPort  string `yaml:"host_port"`
	Namespace string `yaml:"namespace"`
}

type AppConfig struct {
	DB            DBConfig      `yaml:"database"`
	Temporal      TemporalConfig `yaml:"temporal"`
	WorkflowCount int            `yaml:"workflow_count"`
	ActivityCount int            `yaml:"activity_count"`
	TaskQueue     string         `yaml:"task_queue"`
}

func GetDefaultConfig() *AppConfig {
	return &AppConfig{
		DB: DBConfig{
			Host:     "localhost",
			Port:     5433,
			User:     "postgres",
			Password: "password",
			DBName:   "test",
			SSLMode:  "disable",
		},
		Temporal: TemporalConfig{
			HostPort:  "localhost:7233",
			Namespace: "default",
		},
		WorkflowCount: 5,
		ActivityCount: 3,
		TaskQueue:     "go-poc",
	}
}
