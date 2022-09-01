package config

import (
	"database/sql"
	"os"
)

type Config struct {
	DB         *sql.DB
	AuthSecret string
}

type configOptions func(*Config)

// WithDBConnection is used to set the db connection of the application
func WithDBConnection(db *sql.DB) configOptions {
	return func(c *Config) {
		c.DB = db
	}
}

// NewConfig used to setup configuration
func NewConfig(opts ...configOptions) *Config {
	c := &Config{
		AuthSecret: os.Getenv("AUTH_SECRET"),
	}
	// set options
	for _, opt := range opts {
		opt(c)
	}

	return c
}
