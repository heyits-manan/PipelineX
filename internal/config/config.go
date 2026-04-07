package config

import "fmt"

// TODO: Define application configuration shared by API and worker.
// TODO: Start small: port, environment, storage backend, queue backend, database DSN.
// TODO: Later, decide whether config comes from env vars, flags, or config files.

type StorageBackend string
type QueueBackend string

type Config struct {
	// TODO: Add fields for app and infrastructure configuration.
	HTTPPort int
	Environment string
	StorageBackend string
	QueueBackend string

}

func Load() (Config, error) {
	return Config{
		HTTPPort: 8080,
		Environment: "development",
		StorageBackend: "local",
		QueueBackend: "local",
	}, nil
}

func MustLoad() Config {
	config, err := Load()
	if err != nil {
		panic(err)
	}
	return config
}

func (c *Config) Validate() error {
	if c.HTTPPort == 0 {
		return fmt.Errorf("HTTPPort is required")
	}
	return nil
}



// Suggested declarations to implement later:
// func Load() (Config, error)
// func MustLoad() Config
