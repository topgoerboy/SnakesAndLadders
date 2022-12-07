package dao

import (
	"context"
	"encoding/json"
	"strings"

	"game/config"
	"game/model"
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

func (dao *userDao) SetRoomInfo(
	ctx context.Context, userId, roomId, opCnt int64, resp *model.SnakesAndLaddersResp,
) (int64, error) {
	cacheKey := strings.ReplaceAll(
		config.RedisKeyRoomInfo, "{userId}", cast.ToString(userId),
	)
	cacheKey = strings.ReplaceAll(cacheKey, "{roomId}", cast.ToString(roomId))

	str, err := json.Marshal(resp)
	if err != nil {
		return 0, err
	}
	return config.Redis.HSet(ctx, cacheKey, opCnt, str).Result()
}
