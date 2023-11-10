package spostgres

import (
	"io/fs"
	"time"
)

// Option -.
type Option func(*Postgres)

// MaxPoolSize -.
func MaxPoolSize(size int) Option {
	return func(c *Postgres) {
		c.maxPoolSize = size
	}
}

// ConnTimeout -.
func ConnTimeout(timeout time.Duration) Option {
	return func(c *Postgres) {
		c.connTimeout = timeout
	}
}

// RootFS -.
func RootFS(rootFS fs.FS) Option {
	return func(c *Postgres) {
		c.rootFS = rootFS
	}
}
