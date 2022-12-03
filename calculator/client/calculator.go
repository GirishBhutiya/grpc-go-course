package main

import (
	"context"
	"log"

	pb "github.com/GirishBhutiya/grpc-go-course/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Println("doCalculate was invocked")
	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		Vlue1: 5,
		Vlue2: 10,
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %d\n", res.Result)
}
