package example

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func UseSet(db *redis.Client, ctx context.Context) {
	db.SAdd(ctx, "set", 100, 200) // 将 100,200 插入set

	db.SMembers(ctx, "set") // 获取所有元素

	db.SRem(ctx, "set")     // 删除所有元素
	db.SPop(ctx, "set")     // 随机获取 set 中的元素，并删除
	db.SPopN(ctx, "set", 3) // 随机获取 set 中的 3 个元素，并删除

	db.SCard(ctx, "set") // 获取 key 中元素数量

	db.SIsMember(ctx, "set", 100) // 元素是否在集合中
}
