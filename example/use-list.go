package example

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func UseList(rdb *redis.Client, ctx context.Context) {
	rdb.LPush(ctx, "list", 1, 2, 3, 4)
	rdb.LPushX(ctx, "list", 1, 2, 3, 4) // key 存在时插入

	rdb.RPush(ctx, "list", 11, 22, 33) // 从右边插入
	rdb.LPush(ctx, "list", 11, 22, 33)
	rdb.LInsert(ctx, "list", "before", 5, 4) // 在 5 前面插入 4
	rdb.LInsert(ctx, "list", "after", 5, 3)  // 在 5 后边插入 4

	rdb.LRem(ctx, "list", 1, 100) // 从左边开始删除 count 个 100 ,count 为负数代表从另一边删除几个，count=0 删除所有

	rdb.RPop(ctx, "list")         // 右侧 弹出栈并删除一个数据
	rdb.LPop(ctx, "list")         // 右侧 弹出栈并删除一个数据
	rdb.LRange(ctx, "list", 0, 1) // 获取一个范围的数据

	rdb.LLen(ctx, "list")      // 返回列表大小
	rdb.LIndex(ctx, "list", 1) // 返回数据的位置
}
