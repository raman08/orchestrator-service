package main

import (
	"context"
	"errors"
	"log"
	"net"
	pb "orchestrator-service/proto"
	"strconv"

	"google.golang.org/grpc"
)

const (
	port = ":10000"
)

type UserNameServer struct {
	pb.UnimplementedOrchestratorServer
}

func (s *UserNameServer) GetMockUserData(ctx context.Context, name *pb.UserRequest) (*pb.User, error) {

	var nme string = name.GetName()

	if(len(nme)) < 6 {
		return nil, errors.New("user name is less then 6 chracters")

	}

	var class string =strconv.Itoa( len(nme))
	var roll int = len(nme) * 10

	return &pb.User{Name: nme, Class: class, Roll: int64(roll)}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Connection Falid! \n ERROR: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterOrchestratorServer(server, &UserNameServer{})
	log.Printf("Mock Data Server listening at: %v", lis.Addr().String())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Mock Data Server Crashed. \n Error: %v", err)
	}
}