package config

// TODO: Define application configuration shared by API and worker.
// TODO: Start small: port, environment, storage backend, queue backend, database DSN.
// TODO: Later, decide whether config comes from env vars, flags, or config files.

type Config struct {
	// TODO: Add fields for app and infrastructure configuration.
}

// Suggested declarations to implement later:
// func Load() (Config, error)
// func MustLoad() Config
