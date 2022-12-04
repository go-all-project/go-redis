package example

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func UseSetGet(rdb *redis.Client, ctx context.Context) {
	err := rdb.Set(ctx, "name", "zhangsan", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	result, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		return
	}
	fmt.Printf("redis get1 %v\n", result)
}
