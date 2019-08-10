package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type dbContextKey int

const dbCtxKey dbContextKey = iota

// ContextWithDB returns a child context with the supplied DB client attached.
func ContextWithDB(ctx context.Context, client *sqlx.DB) context.Context {
	return context.WithValue(ctx, dbCtxKey, client)
}

// ContextDB returns the DB client attached to the provided context.
func ContextDB(ctx context.Context) *sqlx.DB {
	return ctx.Value(dbCtxKey).(*sqlx.DB)
}
