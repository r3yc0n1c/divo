package config

import "os"

var (
	DefaultTarget = "http://localhost:3000"
	DefaultPort   = ":8080"
)

type Config struct {
	TargetURL  string
	ListenPort string
}

func LoadFromEnv() *Config {
	target := os.Getenv("TARGET_URL")
	if target == "" {
		target = DefaultTarget
	}

	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = DefaultPort
	}

	return &Config{
		TargetURL:  target,
		ListenPort: port,
	}
}
