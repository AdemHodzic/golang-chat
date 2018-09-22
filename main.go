package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	
	port := os.Getenv("PORT")
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(melody *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(port)

}
