package app

import (
	"log"

	"github.com/piotrpersona/saga/broker"

	"github.com/go-redis/redis"
	"github.com/piotrpersona/saga/config"
	"github.com/piotrpersona/saga/events"
	"github.com/piotrpersona/saga/store"
)

func runOrderSubscriber(config config.Config, broker broker.Broker) (err error) {
	var redisClient *redis.Client
	redisClient, err = store.NewRedisClient(config.RedisReserveConfig)
	pubSub := redisClient.Subscribe(config.RedisOrdersConfig.Channel)
	_, err = pubSub.Receive()
	if err != nil {
		return
	}

	orderCreateChannel := pubSub.Channel()

	orderHandler := events.RedisOrderHandler{
		Client:  redisClient,
		Channel: config.RedisOrdersConfig.Channel,
		Broker:  broker,
	}

	for order := range orderCreateChannel {
		err = orderHandler.Handle(order.Payload)
		if err != nil {
			log.Println(err)
			break
		}
	}

	return
}
