package main

import (
    "fmt"
    "net/http"
    "io"
    "os"
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
    pool := websocket.NewPool()
    go pool.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        io.WriteString(w, "CHATAPP 1.0")
    })
}

func main() {
    port := os.Getenv("PORT")
	setupRoutes()
    http.ListenAndServe(":"+port, nil)
    
}
