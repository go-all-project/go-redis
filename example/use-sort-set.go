package example

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func UseSortSet(db *redis.Client, ctx context.Context) {
	// 添加集合元素到集合中
	db.ZAdd(ctx, "sort-set", &redis.Z{
		Score:  2.5,
		Member: "张三",
	})

	db.ZRem(ctx, "sort-set", "张三", "李四")            // 删除
	db.ZRemRangeByRank(ctx, "sort-set", 0, 5)       // 根据返回删除,按最低分最高分排列，-1为最高分
	db.ZRemRangeByScore(ctx, "sort-set", "2", "(5") // 根据分数范围删除

	db.ZScore(ctx, "sort-set", "张三") // 查询对应的分数

	db.ZRank(ctx, "sort-set", "张三")    // 查询对应的排名,排序从大到小排
	db.ZRevRank(ctx, "sort-set", "张三") // 查询对应的排名,排序从大到小排

	db.ZCard(ctx, "sort-set") // 返回元素个数

	db.ZCount(ctx, "sort-set", "(1", "5") // 返回 1 < 分数 <= 5 , (1：代表大于1，相当于去掉了等于

	db.ZIncrBy(ctx, "sort-set", 2, "张三") // 给张三增加两分

	db.ZRange(ctx, "sort-set", 2, 5)    // 获取索引返回的元素 0,-1 为所有元素,从小到大
	db.ZRevRange(ctx, "sort-set", 2, 5) // 反向
	// 根据分数范围返回集合，根据分数从小到达排序，支持分页
	db.ZRangeByScore(ctx, "sort-set", &redis.ZRangeBy{
		Min:    "2",  // 最小分
		Max:    "10", // 最大分
		Offset: 0,    // 类似 sql limit, 开始偏移量
		Count:  5,    // 一次性返回多少
	})
	// 反向
	db.ZRevRangeByScore(ctx, "sort-set", &redis.ZRangeBy{
		Min:    "2",  // 最小分
		Max:    "10", // 最大分
		Offset: 0,    // 类似 sql limit, 开始偏移量
		Count:  5,    // 一次性返回多少
	})
	// 除了返回集合元素，还返回分数 从小到大
	db.ZRangeByScoreWithScores(ctx, "sort-set", &redis.ZRangeBy{
		Min:    "2",  // 最小分
		Max:    "10", // 最大分
		Offset: 0,    // 类似 sql limit, 开始偏移量
		Count:  5,    // 一次性返回多少
	})
	// 反向
	db.ZRevRangeByScoreWithScores(ctx, "sort-set", &redis.ZRangeBy{
		Min:    "2",  // 最小分
		Max:    "10", // 最大分
		Offset: 0,    // 类似 sql limit, 开始偏移量
		Count:  5,    // 一次性返回多少
	})

}
