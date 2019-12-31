package broker

import (
	"fmt"

	"github.com/piotrpersona/saga/model"

	"github.com/go-redis/redis"
)

type redisBroker struct {
	channel string
	client  *redis.Client
}

func (rb redisBroker) Save(order model.Order) (orderID string, err error) {
	orderID = order.ID
	channelName := fmt.Sprintf("%s.%s", rb.channel, orderID)
	_, err = rb.client.Publish(channelName, order).Result()
	return
}
