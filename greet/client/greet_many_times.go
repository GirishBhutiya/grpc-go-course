package main

import (
	"context"
	"io"
	"log"

	pb "github.com/GirishBhutiya/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("DoGreetManyTimes was invoked")

	req := &pb.GreetRequest{FirstName: "Girish"}

	strem, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTime: %v\n", err)
	}

	for {
		msg, err := strem.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes %s\n", msg.Result)

	}
}
