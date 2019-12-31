package service

import (
	"context"

	"github.com/piotrpersona/saga/broker"
	"github.com/piotrpersona/saga/model"
	"github.com/piotrpersona/saga/random"

	"github.com/piotrpersona/saga/order"
)

type OrderService struct {
	Broker broker.Broker
}

func (o *OrderService) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	const IDLength = 32
	orderID := random.ID(IDLength)

	newOrder := model.Order{
		ID:         orderID,
		CustomerID: request.UserID,
		Amount:     request.Amount,
		Status:     model.OrderPending,
	}

	orderID, err := o.Broker.Save(newOrder)

	return &order.CreateOrderResponse{
		OrderID: orderID,
		Message: "Order pending",
	}, err
}
