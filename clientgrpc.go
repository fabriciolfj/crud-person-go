package main

import (
	"context"
	"fmt"
	pb "github.com/person/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewPersonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	e, err := c.GetPerson(ctx, &pb.PersonRequest{Id: "1"})
	if err != nil {
		log.Fatalf("could not  get person: %v", err)
		st, ok := status.FromError(err)
		if ok {
			fmt.Printf("error code: %v, message %s\n", st.Code(), st.Message())
		}
	}

	log.Printf("Person: %s", e.Name)

	stream, err := c.WatchPersons(ctx, &pb.WatchRequest{})

	if err != nil {
		log.Fatalf("fail stream %s", err)
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error receive %v", err)
		}

		fmt.Println(resp)
	}

}
