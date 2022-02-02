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
	port = ":9000"
	address = "localhost:9001"
)

type UserNameServer struct {
	pb.UnimplementedOrchestratorServer
}

func (s *UserNameServer) GetUserByName(ctx context.Context, name *pb.UserRequest) (*pb.User, error) {

	log.Println("Orchestrator 1 GetUserByName Called!")

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed To connect to orchestrator 2. \n Error: ", err)
	}
	defer conn.Close()

	client := pb.NewOrchestratorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	name_to_send := name.GetName()

	res, err := client.GetUser(ctx, &pb.UserRequest{Name: name_to_send})


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

	log.Printf("Mock Data Server listening at: %v", lis.Addr().String())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Mock Data Server Crashed. \n Error: %v", err)
	}
}