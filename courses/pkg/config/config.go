package config

import (
	"database/sql"
)

type Config struct {
	DB *sql.DB
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
	c := &Config{}
	// set options
	for _, opt := range opts {
		opt(c)
	}
	return c
}
