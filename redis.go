package cstools

import (
	"github.com/go-redis/redis"
)

func InitRedis(host string, port string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb

}
