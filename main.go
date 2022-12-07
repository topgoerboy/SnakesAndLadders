package main

import (
	websocket "game/ws"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	m := websocket.InitMelody()
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	r.Run(":9090")
}
