package app

import (
	"log"
	"sync"

	"github.com/piotrpersona/saga/config"
)

func Run() {
	redisHost := "localhost:6379"
	redisPassword := "topsecret"
	config := config.Config{
		OrdersConfig: config.OrdersConfig{
			Host: "localhost",
			Port: 8081,
		},
		RedisReserveConfig: config.RedisConfig{
			Channel:  "orders",
			URI:      redisHost,
			Password: redisPassword,
		},
		RedisOrdersConfig: config.RedisConfig{
			Channel:  "order.reserve",
			URI:      redisHost,
			Password: redisPassword,
		},
	}

	brokerName := "redis"

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	broker, err := createBroker(config, brokerName)
	if err != nil {
		log.Fatal(err)
	}

	go runOrderSubscriber(config, broker)
	runOrderService(config, broker)
}
