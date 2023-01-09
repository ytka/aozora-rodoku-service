package database

import (
	"context"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

//type closeFunc func() error
//func (cf closeFunc) Close() error { return cf() }

// func Open(ctx context.Context) (io.Closer, error) {

func Open(ctx context.Context) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	/*
		return closeFunc(func() error {
			return conn.Close(ctx)
		}), nil
	*/
	return db, nil
}

func DoTransaction(ctx context.Context, db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			// TODO: Rollback 失敗かつ、その原因となった error 両方を把握したい
			// go 1.20 の errors.join を使いたい
			return fmt.Errorf("rollback failed: cause %w", err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func ExecQuery(ctx context.Context, tx *sqlx.Tx, query string, arg interface{}) error {
	var err error
	if arg == nil {
		_, err = tx.ExecContext(ctx, query)
	} else {
		_, err = tx.NamedExecContext(ctx, query, arg)
	}
	if err != nil {
		return err
	}
	return nil
}
