// main.go
package main

import (
	"aozorarodoku-service/env"
	"aozorarodoku-service/usecase"
	"context"
)

func do(e *env.Env) error {
	return usecase.RegisterContentsToDB(e.Context, e.Db)
}

func runOnEnv(dof func(e *env.Env) error) error {
	ctx := context.Background()
	env, err := env.New(ctx)
	if err != nil {
		return err
	}

	defer env.Teardown()
	if err := dof(env); env != nil {
		return err
	}
	return nil
}

func main() {
	if err := runOnEnv(func(e *env.Env) error {
		return do(e)
	}); err != nil {
		panic(err)
	}
}
