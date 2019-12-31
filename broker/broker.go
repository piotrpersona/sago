package broker

import (
	"github.com/piotrpersona/saga/model"
)

type Broker interface {
	Save(order model.Order) (orderID string, err error)
}
