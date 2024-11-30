package database

import (
	"context"
	"database/sql"
)

type Database struct {
	Ctx context.Context
	*sql.DB
}

func NewDatabase(ctx context.Context) (*Database, error) {
	return newPostgresDatabase(ctx)
}
