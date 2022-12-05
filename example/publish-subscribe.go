package example

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func UsePublishSubscribe(db *redis.Client, ctx context.Context) {

	go subscribe(db, ctx)

	publish(db, ctx)
}
func publish(db *redis.Client, ctx context.Context) {
	fmt.Printf("发布 %v\n", 1)
	db.Publish(ctx, "channel1", "消息")
}

func subscribe(db *redis.Client, ctx context.Context) {
	fmt.Printf("订阅 %v\n", 1)
	sub := db.Subscribe(ctx, "channel1")
	//db.PSubscribe(ctx, "channel*") // 以 channel 开头的都进行订阅

	for msg := range sub.Channel() {
		fmt.Printf("%v\n", msg.Channel)
		fmt.Printf("%v\n", msg.Payload)
	}

	// 第二种获取消息方式
	//for {
	//	message, err := sub.ReceiveMessage(ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Printf("%v\n", message.Channel)
	//	fmt.Printf("%v\n", message.Payload)
	//}

	//sub.Unsubscribe(ctx,"channel1") // 取消订阅

	// 查询有多少订阅者
	chs, _ := db.PubSubNumSub(ctx, "channel1").Result()
	for ch, count := range chs {
		fmt.Printf("%v\n", ch)    // 名字
		fmt.Printf("%v\n", count) // 数量
	}
}
