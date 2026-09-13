package config

import "time"

type Config struct {
	HTTP HTTP `yaml:"http"`
	GRPC GRPC `yaml:"grpc"`
}

type HTTP struct {
	URL     string        `yaml:"url"`
	Timeout time.Duration `yaml:"timeout"`
}

type GRPC struct {
	Address string        `yaml:"address"`
	Timeout time.Duration `yaml:"timeout"`
}
