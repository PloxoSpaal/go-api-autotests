package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

func defaults() *Config {
	return &Config{
		HTTP: HTTP{
			URL:     "http://localhost:8000/api/v1",
			Timeout: 10 * time.Second,
		},
		GRPC: GRPC{
			Address: "localhost:9000",
			Timeout: 10 * time.Second,
		},
	}
}

func Load() (*Config, error) {
	cfg := defaults()

	file := os.Getenv("TEST_CONFIG_FILE")
	if file == "" {
		return cfg, nil
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", file, err)
	}

	if err = yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("decode config %q: %w", file, err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}
