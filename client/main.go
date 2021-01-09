package main

import (
	"context"
	"github.com/memochou1993/grpc-go-example/gen"
	pb "github.com/memochou1993/grpc-go-example/gen"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addr := "127.0.0.1:8080"
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	c := hello.NewHelloServiceClient(conn)
	r, err := c.SayHello(ctx, &pb.HelloRequest{Greeting: "World!"})
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("Response received: %s", r.GetReply())
}
