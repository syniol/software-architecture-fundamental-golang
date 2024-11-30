package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var postgresInstanceOnce sync.Once

var postgresInstance *Database

const maximumConnections = 1

func newPostgresDatabase(ctx context.Context) (*Database, error) {
	if postgresInstance != nil {
		return postgresInstance, nil
	}

	var dbError error
	postgresInstanceOnce.Do(func() {
		// [Postgres SSL Support]: https://www.postgresql.org/docs/current/libpq-ssl.html
		connStr := fmt.Sprintf(
			"postgresql://%s:%s@%s/%s?sslmode=require",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			func() string {
				if len(os.Getenv("DEBUG")) > 0 {
					return "127.0.0.1"
				}

				return "db"
			}(),
		)
		cnn, err := sql.Open("postgres", connStr)
		if err != nil {
			dbError = err
		}

		cnn.SetMaxOpenConns(maximumConnections)

		postgresInstance = &Database{
			Ctx: ctx,
			DB:  cnn,
		}
	})
	if dbError != nil {
		return nil, dbError
	}

	return postgresInstance, nil
}
