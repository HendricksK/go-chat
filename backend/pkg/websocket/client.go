package websocket

import (
	"fmt"
	"log"
	// "sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int `json:"type"`
	Body string `json:"body`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c // Need to check in on this operator
		// Has to do with channels https://tour.golang.org/concurrency/2
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage() 
		if err != nil {
			log.Println(err) // write to file
			return // return nothing, after write
		}
		message := Message {
			Type: messageType,
			Body: string(p),
		}
		c.Pool.Broadcast <- message // Adding message to channel
		fmt.Printf("Message Received: %+v\n", message)
	}

}