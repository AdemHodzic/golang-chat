package main

import (
  "github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"net/http"
)

func main() {

  r := gin.Default()
  m := melody.New()

  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  r.GET("/", func(c *gin.Context) {
    http.ServeFile(c.Writer, c.Request, "static/index.html")
  })

  r.GET("/ws", func(c *gin.Context) {
    m.HandleRequest(c.Writer, c.Request)
  })

  m.HandleMessage(func(melody *melody.Session, msg []byte) {
    m.Broadcast(msg)
  })

  r.Run(":5000")

}
