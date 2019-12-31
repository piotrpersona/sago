package main

import (
	"fmt"
	"log"

	"github.com/piotrpersona/saga/events"

	"github.com/piotrpersona/saga/config"
	"github.com/piotrpersona/saga/store"
)

func main() {
	config := config.RedisConfig{
		Channel:  "order.reserve",
		URI:      "localhost:6379",
		Password: "topsecret",
	}
	client, err := store.NewRedisClient(config)
	if err != nil {
		log.Fatal(err)
	}

	orderEvent := events.RedisOrderEvent{
		OrderID:    "111",
		CustomerID: "111",
		Amount:     200,
		Status:     "APPROVED",
	}

	resp, err := client.Publish(config.Channel, orderEvent).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
