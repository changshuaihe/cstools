package cstools

import (
	"github.com/go-redis/redis"
)

func InitRedis(host string, port string, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       0,  // use default DB
	})
	return rdb

}
