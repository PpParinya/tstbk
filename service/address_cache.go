package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/tst/backend/database/redisclient"
	"github.com/tst/backend/utils"
)

var ctx = context.Background()

func GetCachedAddress(lat, lng float64) (string, bool) {
	key := fmt.Sprintf("addr:%.6f:%.6f", lat, lng)
	val, err := redisclient.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		log.Println("Redis get error:", err)
		return "", false
	}
	return val, true
}

func FetchAddressAsync(lat, lng float64) {
	go func() {
		addr := utils.ReverseGeocode(lat, lng)
		key := fmt.Sprintf("addr:%.6f:%.6f", lat, lng)
		err := redisclient.Client.Set(ctx, key, addr, 24*time.Hour).Err()
		if err != nil {
			log.Println("Redis set error:", err)
		}
	}()
}