package main

import (
    "fmt"
    "net/http"
)

func main() {

    fmt.Println("chat app v0.0.1")

	setupRoutes()
    http.ListenAndServe(":8080", nil)
    
}
