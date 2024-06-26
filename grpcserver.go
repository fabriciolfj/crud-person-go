package main

import (
	"context"
	pb "github.com/person/grpc"
	"github.com/person/model"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strconv"
)

var persons = []model.Person{
	{ID: 1, Name: "John", Age: 39, Uuid: uuid.NewV4().String()},
	{ID: 2, Name: "Ricardo", Age: 39, Uuid: uuid.NewV4().String()},
	{ID: 3, Name: "Mateus", Age: 39, Uuid: uuid.NewV4().String()},
}

type server struct {
	pb.UnimplementedPersonServiceServer
}

func (s *server) GetPerson(ctx context.Context, in *pb.PersonRequest) (*pb.PersonResponse, error) {
	var result *pb.PersonResponse
	for _, person := range persons {
		idFormat := strconv.FormatUint(uint64(person.ID), 10)
		if idFormat == in.Id {
			result = &pb.PersonResponse{Id: idFormat, Name: person.Name, Age: int32(person.Age)}
		}
	}

	if result == nil {
		return nil, status.Errorf(codes.NotFound, "not found", in.GetId())
	}

	return result, nil
}

func (s *server) WatchPersons(req *pb.WatchRequest, srv pb.PersonService_WatchPersonsServer) error {
	for _, person := range persons {
		idFormat := strconv.FormatUint(uint64(person.ID), 10)
		if err := srv.Send(&pb.PersonResponse{Id: idFormat, Name: person.Name, Age: int32(person.Age)}); err != nil {
			return err
		}
	}

	return nil
}

func mainOld() {
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
