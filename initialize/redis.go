package initialize

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
			PoolSize: 10,
		},
	)
	result := rdb.Ping(context.Background())
	fmt.Println("redis ping:", result.Val())
	if result.Val() != "PONG" {
		// 连接有问题
		color.Red("[InitRedis] 链接redis异常:")
		return nil
	}

	return rdb
}
