package main

import (
	"context"
	"log"
	"net"
	pb "orchestrator-service/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":9001"
	address = "localhost:10000"
)

type UserNameServer struct {
	pb.UnimplementedOrchestratorServer
}

func (s *UserNameServer) GetUser(ctx context.Context, name *pb.UserRequest) (*pb.User, error) {

	log.Println("Orchestrator 2 GetUser Called!")
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed To connect to mock data server. \n Error: ", err)
	}
	defer conn.Close()

	client := pb.NewOrchestratorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	name_to_send := name.GetName()

	res, err := client.GetMockUserData(ctx, &pb.UserRequest{Name: name_to_send})


	if err != nil {
		return nil, err
	}

	var nm string = res.GetName()
	var class string = res.GetClass()
	var roll int = int(res.GetRoll())

	return &pb.User{Name: nm, Class: class, Roll: int64(roll)}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Connection Falid! \n ERROR: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterOrchestratorServer(server, &UserNameServer{})
	log.Printf("Orchestrator 2 listening at: %v", lis.Addr().String())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Orchestrator 2 Crashed. \n Error: %v", err)
	}
}