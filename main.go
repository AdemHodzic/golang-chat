package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"log"
	"fmt"
)

func determineListenAdress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("CAN'T FINT $PORT\n")
		return "", fmt.Errorf("$PORT not set!")
	
	}

	return ":" + port, nil

}

func main() {
	
	addr, err := determineListenAdress()
	
	if err != nil {
		log.Fatal(err)
	}
	
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

	r.Run(addr)
}
