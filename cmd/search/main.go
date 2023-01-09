// main.go
package main

import (
	"aozorarodoku-service/env"
	"aozorarodoku-service/usecase"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	searchText := "宮沢賢治"
	if err := env.RunOn(ctx, func(e *env.Env) error {
		contents, err := usecase.SearchContents(e.Context, e.Db, searchText)
		if err != nil {
			return err
		}
		for _, v := range contents {
			fmt.Println(v)
		}
		return nil
	}); err != nil {
		panic(err)
	}
}
