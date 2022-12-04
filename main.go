package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-redis/example"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	ctx := context.Background()

	example.UseSetGet(rdb, ctx)
}
