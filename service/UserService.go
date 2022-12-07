package service

import (
	"errors"

	"game/config"
	"game/dao"
	"game/model"
	"github.com/spf13/cast"
)

type userService struct{}

var UserService userService

func (srv *userService) NewConnect() (*model.User, error) {
	userId, err := dao.UserDao.GetUserId(config.Context)
	if err != nil {
		return nil, err
	}

	userInfo := &model.User{
		UserId:   userId,
		UserName: "Player-" + cast.ToString(userId),
		RoomId:   1,
		DicePos:  0,
	}

	userInfoMap := make(map[string]string)
	userInfoMap["user_id"] = cast.ToString(userInfo.UserId)
	userInfoMap["user_name"] = cast.ToString(userInfo.UserName)
	userInfoMap["room_id"] = cast.ToString(userInfo.RoomId)
	userInfoMap["dice_pos"] = cast.ToString(userInfo.DicePos)
	result, err := dao.UserDao.SetUserInfo(config.Context, userInfoMap)
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, errors.New("set user info fail")
	}

	return userInfo, nil
}
