package dao

import (
	"context"
	"strings"

	"game/config"
	"github.com/spf13/cast"
)

type userDao struct{}

var UserDao userDao

func (dao *userDao) GetUserId(ctx context.Context) (int64, error) {
	return config.Redis.Incr(ctx, config.RedisKeyUserIdMax).Result()
}

func (dao *userDao) GetUserInfo(ctx context.Context, userId int64) (map[string]string, error) {
	cacheKey := strings.ReplaceAll(config.RedisKeyUserInfo, "{userId}", cast.ToString(userId))

	userInfo, err := config.Redis.HGetAll(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (dao *userDao) SetUserInfo(ctx context.Context, userInfo map[string]string) (bool, error) {
	cacheKey := strings.ReplaceAll(
		config.RedisKeyUserInfo, "{userId}", cast.ToString(userInfo["user_id"]),
	)

	return config.Redis.HMSet(ctx, cacheKey, userInfo).Result()
}
