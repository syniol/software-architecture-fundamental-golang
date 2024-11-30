package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	Ctx context.Context
	*sql.DB
}

func NewClient(ctx context.Context) (*Database, error) {
	return newPostgresClient(ctx)
}
