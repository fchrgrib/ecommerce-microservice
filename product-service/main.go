package main

import (
	"fmt"
	"github.com/product-service/common/config"
	"github.com/product-service/common/service"
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
	product := &config.ProductServiceServer{}

	fmt.Println("listening port:", port)

	service.RegisterProductServiceServer(srv, product)
	srv.Serve(lis)
}
