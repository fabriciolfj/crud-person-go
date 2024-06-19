package main

import (
	"context"
	pb "github.com/person/grpc"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedPersonServiceServer
}

func (s *server) GetPerson(ctx context.Context, in *pb.PersonRequest) (*pb.PersonResponse, error) {
	return &pb.PersonResponse{Id: uuid.NewV4().String(), Name: "Teste", Age: 1}, nil
}

func (s *server) WatchPersons(req *pb.WatchRequest, srv pb.PersonService_WatchPersonsServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{})

	log.Println("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
