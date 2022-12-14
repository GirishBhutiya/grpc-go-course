package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/GirishBhutiya/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v\n", in)

	oId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can not parse ID")
	}
	res, err := collection.DeleteOne(ctx, bson.M{"_id": oId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Can not delete object in MongoDB: %v\n", err))
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln("Blog was not found"))
	}

	return &emptypb.Empty{}, nil

}
