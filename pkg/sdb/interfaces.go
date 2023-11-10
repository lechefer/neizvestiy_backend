package sdb

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type BeginTxer interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}
