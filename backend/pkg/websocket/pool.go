package websocket

import "fmt"

type Pool struct {
	Register chan *Client
	Unregister chan *Client
	Clients map[*Client]bool
	Broadcast chan Message
}

func NewPool() *Pool {
	// https://tutorialedge.net/projects/chat-system-in-go-and-react/part-4-handling-multiple-clients/
	return &Pool {
		Register: 	make(chan *Client),
		Unregister: make(chan *Client),
		Clients: 	make(map[*Client]bool),
		Broadcast:	make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
			case client := <-pool.Register:
				pool.Clients[client] = true
				fmt.Println("Size of connection pool: ", len(pool.Clients))
				for client, _ := range pool.Clients {
					fmt.Println(client)
						client.Conn.WriteJSON(Message {Type: 1,Body: "New user joined..."})
				}
				break
			case client := <-pool.Unregister:
				delete(pool.Clients, client)
				fmt.Println("Size of connection pool: ", len(pool.Clients))
				for client, _ := range pool.Clients {
					fmt.Println(client)
						client.Conn.WriteJSON(Message {Type: 1,Body: "User disconnected..."})
				}
				break
			case message := <-pool.Broadcast:
				fmt.Println("Sending message to all clients in Pool")
				for client, _ := range pool.Clients {
					if err := client.Conn.WriteJSON(message); err != nil {
						fmt.Println(err)
						return
					}
				}
		}
	}
}