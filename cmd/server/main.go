package main

import (
	"context"
	pb "example/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type srv struct{ pb.UnimplementedGreeterServer }

func (s *srv) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + r.Name}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &srv{})
	log.Fatal(s.Serve(lis))
}
