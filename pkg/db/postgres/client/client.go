package client

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/mynamesishuman/go_platform/pkg/db/postgres"
	"github.com/pkg/errors"
)

type pgClient struct {
	masterDBC db.Db
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (c *pgClient) Db() db.Db {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
