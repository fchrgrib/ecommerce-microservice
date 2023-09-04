package main

import (
	"fmt"
	"github.com/user-service/common/config"
	"github.com/user-service/common/service"
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
	product := &config.UserService{}

	fmt.Println("listening port:", port)

	service.RegisterUserServiceServer(srv, product)
	srv.Serve(lis)
}
