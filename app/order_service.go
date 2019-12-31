package app

import (
	"fmt"
	"log"
	"net"

	"github.com/piotrpersona/saga/broker"
	"github.com/piotrpersona/saga/config"
	"github.com/piotrpersona/saga/order"
	"github.com/piotrpersona/saga/service"
	"google.golang.org/grpc"
)

func createBroker(config config.Config, brokerName string) (b broker.Broker, err error) {
	b, err = broker.New(brokerName, config)
	return
}

func runOrderService(config config.Config, broker broker.Broker) {
	grpcListenAddress := fmt.Sprintf("%s:%d", config.OrdersConfig.Host, config.OrdersConfig.Port)
	lis, err := net.Listen("tcp", grpcListenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC listening on: %s\n", grpcListenAddress)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	order.RegisterOrderServer(grpcServer, service.NewOrderService(broker))
	grpcServer.Serve(lis)
}
