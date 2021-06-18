package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client
var Ctx = context.Background()

const Counter = "counter"

func ConnectRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := RedisDB.Ping(context.Background()).Result()
	if err == nil {
		log.Println("redis 回應成功，", pong)
	} else {
		log.Fatal("redis 無法連線，錯誤為", err)
	}

	err = RedisDB.Set(Ctx, Counter, 0, 0).Err()
	if err != nil {
		panic(err)
	}
}
