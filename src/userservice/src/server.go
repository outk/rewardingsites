package main

import (
	"log"
	"net"

	"github.com/outk/rewardingsites/src/userservice/src/services"
	"google.golang.org/grpc"

	pb "github.com/outk/rewardingsites/src/userservice/src/pb"
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

	pb.RegisterUserServiceServer(server, &services.UserService{})
	log.Println("gRPC server start!")
	server.Serve(listenPort)
}
