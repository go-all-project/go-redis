package example

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func UseOriginCmdDo(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.Do(ctx, "get", "name").Result()
	if err != nil {
		return
	}

	fmt.Printf("%v", result.(string))
}
