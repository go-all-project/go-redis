package example

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func UseHash(rdb *redis.Client, ctx context.Context) {

	// 设置 hash
	rdb.HSet(ctx, "hash", "key", "value1")
	rdb.HSet(ctx, "hash", "key2", "value2")
	rdb.HMSet(ctx, "hash", map[string]string{
		"key": "HMset",
	})
	rdb.HSetNX(ctx, "hash", "id", "100")

	rdb.HGet(ctx, "hash", "value1")
	// 获取 hash 下所有键值对
	rdb.HGetAll(ctx, "hash")
	// 获取所有值
	rdb.HMGet(ctx, "hash", "key", "key2")

	// 删除
	rdb.HDel(ctx, "key", "key2")

	// 累加
	rdb.HIncrBy(ctx, "hash", "count", 2)

	// 获取所有 key
	rdb.HKeys(ctx, "hash")

	// 获取数量
	rdb.HLen(ctx, "hash")

	// 判断 id 字段是否存在
	rdb.HExists(ctx, "key", "id")
}
