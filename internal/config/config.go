package config

import "time"

type Config struct {
	Env         string `yaml:"env" env-default:"development" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8081"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}
