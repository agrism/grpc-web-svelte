package category

import (
	"context"
	"github.com/agrism/grpc-web-svelte/backend/proto"
	"log"
)

type Server struct {
}

func (s *Server) Index(ctx context.Context,  request *proto.IndexRequest) (*proto.Categories, error){
	log.Printf("Receoved request from client %v", request)
	return &proto.Categories{Id:"1",Name:"someName"}, nil
}
