package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Client interface {
	Db() Db
	Close() error
}

type Handler func(ctx context.Context) error

type TransactionManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

type Query struct {
	Name     string
	QueryRaw string
}

type Transaction interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type SqlContract interface {
	NamedQuery
	RawQuery
}

type NamedQuery interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type RawQuery interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type Db interface {
	SqlContract
	Transaction
	Pinger
	Close()
}
