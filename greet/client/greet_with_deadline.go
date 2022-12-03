package main

import (
	"context"
	"log"
	"time"

	pb "github.com/GirishBhutiya/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadlines was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Girish",
	}
	res, err := c.GreetWithDeadLine(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline Exceeded")
				return
			} else {
				log.Fatalf("Unexpected gRPC error:%v\n", err)
			}
		} else {
			log.Fatalf("Non gRPC errors: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline %s\n", res.Result)
}
