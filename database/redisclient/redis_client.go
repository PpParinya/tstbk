package redisclient

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // แก้ตามที่ใช้จริง
		Password: "",               // ถ้ามี password ใส่ตรงนี้
		DB:       0,
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
}
