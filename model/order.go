package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Order struct {
	ID         string  `json:"ID"`
	CustomerID string  `json:"CustomerID"`
	Amount     float64 `json:"Amount"`
	Status     string  `json:"Status"`
}

func (o Order) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}

const (
	OrderPending  = "PENDING"
	OrderApproved = "APPROVED"
	OrderRejected = "REJECTED"
)

func OrderStatus(statusName string) (status string, err error) {
	upperCaseStatus := strings.ToUpper(statusName)
	supportedStatuses := map[string]bool{
		OrderPending: true, OrderApproved: true, OrderRejected: true,
	}
	if _, ok := supportedStatuses[upperCaseStatus]; ok {
		status = upperCaseStatus
	} else {
		err = fmt.Errorf("Status: %s not supported", statusName)
	}
	return
}
