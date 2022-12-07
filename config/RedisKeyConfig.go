package config

import "github.com/go-redis/redis/v8"

var (
	Redis *redis.Client
)

const (
	RedisKeyUserIdMax = "user:id:max"
	RedisKeyUserInfo  = "user:info:{userId}"
)
