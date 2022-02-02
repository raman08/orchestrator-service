package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	pb "orchestrator-service/proto"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type UserNameServer struct {
	pb.UnimplementedOrchestratorServer
}

// GetUserByName Method response
func (s *UserNameServer) GetUserByName(ctx context.Context, name *pb.UserRequest) (*pb.User, error) {
	fmt.Println("Get Func Called")
	return nil, errors.New("not implemented yet. Raman Preet Singh will implement me")
}

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Connection Falid! \n ERROR: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterOrchestratorServer(server, &UserNameServer{})

	log.Printf("Server Started at %v", lis.Addr().String())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server. Error: %v", err)
	}

}