package main

import (
    "fmt"
    "net/http"

    "github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websocket.Client{
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()
}

func setupRoutes() {
    // http.HandleFunc("/ws", serveWs)

    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Simple Server")
    // })

    pool := websocket.NewPool()
    go pool.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })
}

func main() {

    fmt.Println("chat app v0.0.1")
	setupRoutes()
    http.ListenAndServe(":8080", nil)
    
}
