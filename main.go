package main

import "net/http"
import "time"

func main() {
    server := http.Server{
        Addr: ":8080",
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  20 * time.Second,
    }

    http.HandleFunc("/user/register", HandleUserRegister)
    server.ListenAndServe()
}
