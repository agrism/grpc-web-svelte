package main

import (
	"github.com/agrism/grpc-web-svelte/backend/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	categories := proto.NewCategoryServiceClient(conn)

	request := proto.IndexRequest{}

	response, err := categories.Index(context.Background(), &request)

	if err != nil{
		log.Fatalf("Error when calling Index: %s", err)
	}

	log.Printf("Response from server: %s", response)
}
