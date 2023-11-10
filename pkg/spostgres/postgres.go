package spostgres

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

const (
	_defaultMaxPoolSize = 10
	_defaultConnTimeout = time.Second
)

var _defaultRootFS = os.DirFS("./")

type Postgres struct {
	*sqlx.DB

	maxPoolSize int
	connTimeout time.Duration
	rootFS      fs.FS

	Builder squirrel.StatementBuilderType
}

// New -.
func New(ctx context.Context, url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize: _defaultMaxPoolSize,
		connTimeout: _defaultConnTimeout,
		rootFS:      _defaultRootFS,

		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	connCtx, connCancel := context.WithTimeout(ctx, pg.connTimeout)
	defer connCancel()

	connConfig, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("can't parse connection: %w", err)
	}
	if connConfig == nil {
		return nil, errors.New("connection config is nil")
	}

	pg.DB = sqlx.NewDb(stdlib.OpenDB(*connConfig), "pgx")
	if err = pg.DB.PingContext(connCtx); err != nil {
		return nil, fmt.Errorf("can't ping server: %w", err)
	}

	return pg, nil
}

func NewWithMigration(ctx context.Context, url string, path string, opts ...Option) (*Postgres, error) {
	pg, err := New(ctx, url, opts...)
	if err != nil {
		return nil, err
	}

	if err = goose.Up(pg.DB.DB, path); err != nil {
		return nil, fmt.Errorf("failed up migrations to postrgres: : %w", err)
	}

	return pg, nil
}

// Close -.
func (p *Postgres) Close() error {
	if err := p.DB.Close(); err != nil {
		return fmt.Errorf("failed to disconnect postrgres: : %w", err)
	}

	return nil
}
