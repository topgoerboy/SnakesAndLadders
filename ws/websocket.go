package websocket

import (
	"encoding/json"
	"sync"

	"game/config"
	"game/ctrl"
	"game/initialize"
	"game/service"
	"game/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gopkg.in/olahol/melody.v1"
)

var (
	idSessionMap sync.Map
	m            *melody.Melody
	logger       *logrus.Logger
)

func InitMelody() *melody.Melody {
	logger = logrus.New()
	config.Redis = initialize.InitRedis()

	m = melody.New()
	m.HandleMessage(Receive)
	m.HandleConnect(Connect)
	m.HandleDisconnect(Disconnect)
	return m
}

func Connect(s *melody.Session) {
	userInfo, err := service.UserService.NewConnect()
	if err != nil {
		logger.Error(err.Error())
	}
	idSessionMap.Store(userInfo.UserId, s)
	s.Set("id", userInfo.UserId)
	Send(
		s, &utils.Message{
			Code: utils.Success,
			Data: userInfo,
		},
	)
}

func Disconnect(s *melody.Session) {
	idObject, ok := s.Get("id")
	if !ok {
		logger.Error("session with no 'id' key")
		return
	}
	idSessionMap.Delete(cast.ToString(idObject))
}

func Send(s *melody.Session, msg *utils.Message) {
	msgByte, _ := json.Marshal(msg)
	if err := s.Write(msgByte); err != nil {
		logger.Error(err)
	}
}

func SendErr(s *melody.Session, err error) {
	Send(s, utils.NewErrMsg(err))
}

func Broadcast(msg *utils.Message) {
	msgByte, _ := json.Marshal(*msg)
	if err := m.Broadcast(msgByte); err != nil {
		logger.Error(err)
	}
}

func Receive(s *melody.Session, msgByte []byte) {
	msg := &utils.Message{}
	if err := json.Unmarshal(msgByte, msg); err != nil {
		Send(s, utils.NewErrMsg(err))
	}

	switch msg.Code {
	case utils.GameStart:
		resp, err := ctrl.SnakesCtrl.Start(s)
		if err != nil {
			SendErr(s, err)
		}
		Broadcast(resp)
	}
}
