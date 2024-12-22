package database

import (
	"context"
	"database/sql"
)

type Database struct {
	Ctx context.Context
	db  *sql.DB
}

// Exec the main method used for Insertion, it could be used for deletion or update as well
func (db *Database) Exec(command string, args ...string) (sql.Result, error) {
	return db.db.Exec(command, args)
}

// Query is for demonstration only to show how you can adapt to sql interfaces given by
// Go SQL built-in library. Your language of choice and its built-in library are core layer
// We could easily replace `db  *sql.DB` with any other database adapting to sql.Rows and sql.Result
func (db *Database) Query(command string, args ...string) (*sql.Rows, error) {
	return db.db.Query(command, args)
}

// NewDatabase creates a new instance of Database to be used inside each repository
func NewDatabase() (*Database, error) {
	return newPostgresDatabase()
}
