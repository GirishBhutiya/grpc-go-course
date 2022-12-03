package main

import (
	"context"
	"log"
	"net"

	pb "github.com/GirishBhutiya/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"
var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalf("Can not listen from mongo: %v\n", err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Can not listen to address: %v\n", err)
	}

	log.Println("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Server %v\n", err)
	}
}
