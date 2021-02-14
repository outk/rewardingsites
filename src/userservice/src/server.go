package main

import (
	"log"
	"net"

	"github.com/outk/rewardingsites/src/userservice/src/service"
	"google.golang.org/grpc"

	pb "github.com/outk/rewardingsites/src/userservice/src/pb/userservice"
)

const (
	port = ":50051"
)

func main() {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	userService := &services.UserService{}

	pb.RegisterUserServiceServer(server, userService)
	server.Serve(listenPort)
}
