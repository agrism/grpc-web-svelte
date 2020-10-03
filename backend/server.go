package main

import (
	"context"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/agrism/grpc-web-svelte/backend/proto"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type RpcServer struct {
	Grpc        *grpc.Server
	WrappedGrpc *grpcweb.WrappedGrpcServer
}

type CategoryServer struct{}

func (instance * CategoryServer) Index(ctx context.Context, in *proto.IndexRequest) (*proto.Categories, error) {

	cat := proto.Categories{
		Categories: []*proto.Category{},
	}

	for i:= 0; i < 1000000; i++ {
		cat.Categories = append(cat.Categories, &proto.Category{
			Id: fmt.Sprintf("%v", i),
			Name: randomdata.City(),
		})
	}

	return &cat, nil
}

func NewServer() *RpcServer {
	var opts []grpc.ServerOption
	//opts = append(opts, ServerInterceptor())
	opts = append(opts)

	// It's increase to 5MB the maximum size allowed for requests and responses
	opts = append(opts, grpc.MaxSendMsgSize(5*1024*1024*1024*1024))
	opts = append(opts, grpc.MaxRecvMsgSize(5*1024*1024*1024*1024))
	gs := grpc.NewServer(opts...)
	return &RpcServer{
		Grpc:        gs,
		WrappedGrpc: grpcweb.WrapServer(gs),
	}
}

func main() {
	rpcServer := NewServer()

	categoryServer := CategoryServer{}


	proto.RegisterCategoryServiceServer(rpcServer.Grpc, &categoryServer)

	Start(":9002", rpcServer)
}

func Start(httpPort string, rpcServer *RpcServer) {
	grpc := rpcServer.WrappedGrpc

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		allowCors(resp, req)
		if grpc.IsGrpcWebRequest(req) || grpc.IsAcceptableGrpcCorsRequest(req) {
			grpc.ServeHTTP(resp, req)
		}
	})

	fmt.Println("HTTP server listening on", httpPort)
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("Failed to start a HTTP server:", err)
	}
}

func allowCors(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}