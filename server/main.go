//Please go through ReadMe file before running this code.

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/varmaexe/go-grpc/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetUserServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failer to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterGreetUserServiceServer(srv, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GreetUser(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", request.GetName())
	// fakeDb = append(fakeDb, request.GetName())

	//Using slice as a fake database just to note the user inputs
	//although a database can be connected, Please check my "golang-portfolio"
	//to know how to connect postgresDB and perform CRUD operations.
	var fakeDb []string
	fakeDb = append(fakeDb, request.GetName())
	for range fakeDb {
		fmt.Print(fakeDb)
	}

	return &pb.Response{Greetings: request.GetName()}, nil

}
