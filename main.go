package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {

	r := gin.Default()
	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(melody *melody.Session, msg []byte) {
		log.Printf("RECEIVED: %v\n", msg)
		m.Broadcast(msg)
	})

	r.Run(":5000")

}
