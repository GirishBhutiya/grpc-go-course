package main

import (
	"context"
	"log"

	pb "github.com/GirishBhutiya/grpc-go-course/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Greet Function was invocked with %v\n", in)
	return &pb.CalculatorResponse{
		Result: in.Vlue1 + in.Vlue2,
	}, nil
}
