package main

import (
	"context"
	"log"

	pb "github.com/GirishBhutiya/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("CreateBlog was invoked")

	blog := pb.Blog{
		AuthorId: "Girish",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), &blog)

	if err != nil {
		log.Fatalf("Unexpected erro : %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}
