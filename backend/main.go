package main

import (
    "fmt"
    "net/http"
    "https://github.com/HendricksK/go-chatapp/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
    ws, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprint(w, "%+V\n", err)
    }

    go websocket.Writer(ws)
    websocket.Reader(ws)
}

func setupRoutes() {
    http.HandleFunc("/ws", serveWs)
}

func main() {

    fmt.Println("chat app v0.0.1")
	setupRoutes()
    http.ListenAndServe(":8080", nil)
    
}
