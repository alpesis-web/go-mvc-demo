package models

import (
    "github.com/go-redis/redis"
)

var client *redis.Client


func Init() {
    client = redis.NewClient(&redis.Options{
        Addr: "mandelbrot-redis:6379",
    })
}
