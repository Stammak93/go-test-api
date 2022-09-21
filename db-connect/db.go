package pg_test

import (
	"fmt"
	"os"

	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	db_pass = os.Getenv("POSTGRESPASS")
	db_user = os.Getenv("PGUSER")
	testing_db = os.Getenv("TESTING_PG")
)

func Connect_Db() *bun.DB {

	dsn := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", db_user, db_pass, testing_db)
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(pgdb, pgdialect.New())

	return db
}

var DB *bun.DB = Connect_Db()
