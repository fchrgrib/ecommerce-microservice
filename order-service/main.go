package main

import (
	"fmt"
	"github.com/order-service/common/config"
	"github.com/order-service/common/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	product := &config.PaymentService{}

	fmt.Println("listening port:", port)

	service.RegisterPaymentServiceServer(srv, product)
	srv.Serve(lis)
}
