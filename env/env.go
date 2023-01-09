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

func New(ctx context.Context) (*Env, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	db, err := database.Open(ctx)
	if err != nil {
		return nil, err
	}
	return &Env{ctx, db}, nil
}

func (e *Env) Teardown() error {
	if err := e.Db.Close(); err != nil {
		return err
	}
	return nil
}
