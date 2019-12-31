package store

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/piotrpersona/saga/config"
)

func NewRedisClient(config config.RedisConfig) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     config.URI,
		Password: config.Password,
		DB:       0,
	})

	_, err = client.Ping().Result()
	if err == nil {
		log.Println("Connected to redis")
	}
	return
}
