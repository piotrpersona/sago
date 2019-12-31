package service

import (
	"github.com/piotrpersona/saga/broker"
)

func NewOrderService(b broker.Broker) *OrderService {
	return &OrderService{
		Broker: b,
	}
}
