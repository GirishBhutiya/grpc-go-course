package main

import (
	"context"
	"io"
	"log"

	pb "github.com/GirishBhutiya/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {

	log.Println("listBlog was invoked")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling list blog: %v\n", err)

	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)

		}
		log.Println(res)
	}

}
