package env

import (
	"context"

	"aozorarodoku-service/database"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Env struct {
	Context context.Context
	Db      *sqlx.DB
}

func setup(ctx context.Context) (*Env, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	db, err := database.Open(ctx)
	if err != nil {
		return nil, err
	}
	return &Env{ctx, db}, nil
}

func (e *Env) teardown() error {
	if err := e.Db.Close(); err != nil {
		return err
	}
	return nil
}

func (e *Env) run(do func(e *Env) error) error {
	defer e.teardown()
	if err := do(e); err != nil {
		return err
	}
	return nil
}

func RunOn(ctx context.Context, do func(e *Env) error) error {
	env, err := setup(ctx)
	if err != nil {
		return err
	}
	return env.run(do)
}
