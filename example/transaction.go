package example

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func UseTransaction(db *redis.Client, ctx context.Context) {
}

func pipe(db *redis.Client, ctx context.Context) {
	pipe := db.TxPipeline() // 开启事务

	// 设置事务进行的原子操作
	// 相当于执行命令：
	//		MULTI
	//			INCR count
	//			EXPIRE count 10
	//		EXEC
	incr := pipe.Incr(ctx, "count")
	pipe.Expire(ctx, "count", time.Second*10)

	// 执行事务
	_, err := pipe.Exec(ctx)
	if err != nil {
		return
	}

	fmt.Printf("%v\n", incr.Val())
}

// redis 乐观锁支持，可以通过 watch 监听一些 key,如果没有被其他人修改的话，才可以提交事务
// 乐观锁：总是假设最好的情况，每次去拿数据的时候都认为别人不会修改，所以不会上锁，只在更新的时候会判断一下在此期间别人有没有去更新这个数据
// 悲观锁：总是假设最坏的情况，每次去拿数据的时候都认为别人会修改，所以每次在拿数据的时候都会上锁，这样别人想拿这个数据就会阻塞
// 直到它拿到锁（共享资源每次只给一个线程使用，其它线程阻塞，用完后再把资源转让给其它线程）。
// 传统的关系型数据库里边就用到了很多这种锁机制，比如行锁，表锁等，读锁，写锁等，都是在做操作之前先上锁。
// Java中synchronized和ReentrantLock等独占锁就是悲观锁思想的实现。
func watch(db *redis.Client, ctx context.Context) {
	fn := func(tx *redis.Tx) error {
		// 查询 key
		v, err := tx.Get(ctx, "key").Int()
		if err != nil {
			return err
		}

		// 业务处理
		v++

		//如果 key 值操作过程中没有修改(没有被别人修改的操作打断) 事务才会成功
		_, err = tx.Pipelined(ctx, func(pipeLiner redis.Pipeliner) error {
			pipeLiner.Set(ctx, "key", v, 0)
			return nil
		})

		if err != nil {
			return err
		}

		return err
	}

	err := db.Watch(ctx, fn, "key", "key1")
	if err != nil {
		return
	}
}
