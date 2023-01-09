// main.go
package main

import (
	"aozorarodoku-service/env"
	"aozorarodoku-service/usecase"
	"context"
)

func main() {
	ctx := context.Background()
	if err := env.RunOn(ctx, func(e *env.Env) error {
		return usecase.RegisterContentsToDB(e.Context, e.Db)
	}); err != nil {
		panic(err)
	}
}
