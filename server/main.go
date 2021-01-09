package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/memochou1993/grpc-go-example/gen"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const (
	grpcServerEndpoint = ":8080"
	httpServerEndpoint = ":8890"
)

type service struct {
	gw.UnimplementedHelloServiceServer
}

func (s *service) SayHello(ctx context.Context, r *gw.HelloRequest) (*gw.HelloResponse, error) {
	log.Printf("Request received: %s", r.GetGreeting())
	return &gw.HelloResponse{Reply: "Hello, " + r.GetGreeting()}, nil
}

func httpServer() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		log.Fatalln(err.Error())
	}
	log.Fatalln(http.ListenAndServe(httpServerEndpoint, mux))
}

func grpcServer() {
	ln, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		log.Fatalln(err.Error())
	}
	s := grpc.NewServer()
	gw.RegisterHelloServiceServer(s, new(service))
	log.Fatalln(s.Serve(ln))
}

func main() {
	go grpcServer()
	httpServer()
}
