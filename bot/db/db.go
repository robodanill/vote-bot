package db

import (
    "fmt"
    "log"

    "github.com/tarantool/go-tarantool"
    "vote_bot/config"
)

var Conn *tarantool.Connection

func Connect() error {
    h := config.TarantoolHost
    if h == "" {
        h = "localhost"
    }

    p := config.TarantoolPort
    if p == "" {
        p = "3301"
    }
	
    u := config.TarantoolUser
    pass := config.TarantoolPass
    addr := fmt.Sprintf("%s:%s", h, p)
    c, err := tarantool.Connect(addr, tarantool.Opts{
        User: u,
        Pass: pass,
    })
    if err != nil {
        return err
    }
    Conn = c
    log.Println("Connected to Tarantool on", addr)
    return nil
}

func Close() {
    if Conn != nil {
        Conn.Close()
    }
}
