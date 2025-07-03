package main

import (
	"context"
	pb "example/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, _ := c.SayHello(context.Background(), &pb.HelloRequest{Name: "gopher"})
	log.Println(r.Message)
}
