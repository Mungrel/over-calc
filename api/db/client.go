package db

import (
	"context"
	"io/ioutil"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"

	// Need this import for its side effects with sqlx.
	_ "github.com/go-sql-driver/mysql"
)

const (
	driver = "mysql"
	dsn    = "app:password@tcp(localhost:3306)/db?parseTime=true&clientFoundRows=true"
)

var (
	dbClient     *sqlx.DB
	dbClientSync sync.Once
)

// Client returns a DB client.
// It will initialise it on first call.
func Client(ctx context.Context) *sqlx.DB {
	dbClientSync.Do(func() {
		client, err := sqlx.ConnectContext(ctx, driver, dsn)
		if err != nil {
			panic(err)
		}

		client.Mapper = reflectx.NewMapper("json")

		err = initTables(ctx, client)
		if err != nil {
			panic(err)
		}

		dbClient = client
	})

	return dbClient
}

func initTables(ctx context.Context, client *sqlx.DB) error {
	const tablesDir = "./db/tables/"
	files, err := ioutil.ReadDir(tablesDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		contents, err := ioutil.ReadFile(tablesDir + file.Name())
		if err != nil {
			return err
		}

		_, err = client.ExecContext(ctx, string(contents))
		if err != nil {
			return err
		}
	}

	return nil
}
