package main

import (
	"context"
	"log"
	pb "orchestrator-service/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	serverAddress = "0.0.0.0:10000"
	// serverAddress = "0.0.0.0:9000"
)

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("DialUp Failed! \n Error: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrchestratorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var new_users = []string {"Raman Preet Singh", "Preet"}

	for _, name := range new_users {

		// response, err := client.GetUserByName(ctx, &pb.UserRequest{Name: name})
		response, err := client.GetMockUserData(ctx, &pb.UserRequest{Name: name})

		if err != nil {
			log.Fatalf("Failed to get User! \n Error: %v", err)
		}


		log.Printf("Name: %v, Class: %v, RollNo: %d", response.Name, response.Class, response.Roll)
	}
}