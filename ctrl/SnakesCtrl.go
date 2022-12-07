package ctrl

import (
	"game/config"
	"game/service"
	"game/utils"
	"gopkg.in/olahol/melody.v1"
)

type snakesCtrl struct{}

var SnakesCtrl snakesCtrl

func (ctrl snakesCtrl) Start(s *melody.Session) (*utils.Message, error) {
	resp, err := service.SnakesService.Start(config.Context, s)
	if err != nil {
		return nil, err
	}

	msg := &utils.Message{
		Code: utils.Success,
		Data: *resp,
	}

	return msg, nil
}
