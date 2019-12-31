package events

import (
	"encoding/json"

	"github.com/piotrpersona/saga/broker"
	"github.com/piotrpersona/saga/model"

	"github.com/go-redis/redis"
)

type RedisOrderEvent struct {
	OrderID    string  `json:"OrderID"`
	CustomerID string  `json:"CustomerID"`
	Amount     float64 `json:"Amount"`
	Status     string  `json:"Status"`
}

func (roe RedisOrderEvent) MarshalBinary() ([]byte, error) {
	return json.Marshal(roe)
}

type RedisOrderHandler struct {
	Channel string
	Client  *redis.Client
	Broker  broker.Broker
}

func (roh RedisOrderHandler) Handle(payload string) (err error) {
	var orderEvent RedisOrderEvent
	err = json.Unmarshal([]byte(payload), &orderEvent)
	if err != nil {
		return
	}

	orderStatus, err := model.OrderStatus(orderEvent.Status)
	if err != nil {
		return
	}

	order := model.Order{
		ID:         orderEvent.OrderID,
		CustomerID: orderEvent.CustomerID,
		Status:     orderStatus,
	}

	_, err = roh.Broker.Save(order)

	return
}
