package main

import (
    "log"
    "net/http"

    "vote_bot/config"
    "vote_bot/db"
    "vote_bot/handlers"
)

func main() {
    err := db.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    p := config.AppPort
    if p == "" {
        p = "8081"
    }
    http.HandleFunc("/command", handlers.HandleCommand)
    log.Println("Starting server on :" + p)
    err = http.ListenAndServe(":"+p, nil)
    if err != nil {
        log.Fatal(err)
    }
}
