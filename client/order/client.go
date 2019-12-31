package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/piotrpersona/saga/random"

	"github.com/piotrpersona/saga/order"

	"google.golang.org/grpc"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second * 2),
	}
	conn, err := grpc.Dial("localhost:8081", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := order.NewOrderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := order.CreateOrderRequest{UserID: random.ID(32), Amount: float64(200)}
	resp, err := client.CreateOrder(ctx, &request)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}
