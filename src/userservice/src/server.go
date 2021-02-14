package main

import (
	"log"
	"net"

	"github.com/c-tag/timemachine/backend/reservecom/service"
	"github.com/c-tag/timemachine/backend/util"
	"github.com/c-tag/timemachine/backend/util/tx"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	server := util.GetGRPCServer()

	userservice.RegisterUserServiceServiceServer(server, &service.UserService{})
	reflection.Register(server)

	go tx.Watch()

	log.Printf("Listening on %v", port)
	if err = server.Serve(listenPort); err != nil {
		log.Fatalln(err)
	}
}
