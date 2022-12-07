package service

import (
	"context"
	"math"
	"math/rand"
	"time"

	"game/dao"
	"game/model"
	"github.com/spf13/cast"
	"gopkg.in/olahol/melody.v1"
)

type snakesService struct{}

var SnakesService snakesService

func (srv *snakesService) GetDiceRollNumber() int64 {
	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())
	return cast.ToInt64(rand.Intn(max-min+1) + min)
}

func (srv *snakesService) Start(ctx context.Context, s *melody.Session) (
	*model.SnakesAndLaddersResp, error,
) {
	id, _ := s.Get("id")
	userId := cast.ToInt64(id)
	userInfo, err := dao.UserDao.GetUserInfo(ctx, userId)
	if err != nil {
		return nil, err
	}
	rollNumber := srv.GetDiceRollNumber()
	resp := &model.SnakesAndLaddersResp{
		DicePoints: rollNumber,
	}
	dicePos, ok := userInfo["dice_pos"]
	if !ok {
		dicePos = "0"
	}
	if cast.ToInt64(dicePos)+rollNumber == 100 {
		resp.IsFinish = true
		resp.GeneralPos = 100
		return resp, nil
	}

	distance := cast.ToInt64(dicePos) + rollNumber - 100
	resp.GeneralPos = 100 - cast.ToInt64(math.Abs(cast.ToFloat64(distance)))
	userInfo["dice_pos"] = cast.ToString(resp.GeneralPos)
	if _, ok := userInfo["operation_cnt"]; !ok {
		userInfo["operation_cnt"] = "1"
	} else {
		count := cast.ToInt64(userInfo["operation_cnt"])
		userInfo["operation_cnt"] = cast.ToString(count + 1)
	}
	_, err = dao.UserDao.SetUserInfo(ctx, userInfo)
	_, _ = dao.UserDao.SetRoomInfo(ctx, userId, 1, cast.ToInt64(userInfo["operation_cnt"]), resp)

	return resp, err
}
