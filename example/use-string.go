package example

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func UseString(rdb *redis.Client, ctx context.Context) {
	// 获取旧值，设置新值
	rdb.GetSet(ctx, "name", "value")

	// 不存在设置值
	rdb.SetNX(ctx, "name", "value", time.Second*10)

	// 批量查询
	rdb.MGet(ctx, "key1", "key2")
	rdb.MSet(ctx, "key1", "value1", "key2", "value2")

	// 自增 自减
	rdb.IncrBy(ctx, "key", 1)
	rdb.DecrBy(ctx, "key", 1)

	// 删除
	rdb.Del(ctx, "key", "key1")
}

func UseOriginCmdDo(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.Do(ctx, "get", "name").Result()
	if err != nil {
		return
	}

	fmt.Printf("%v", result.(string))
}

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
