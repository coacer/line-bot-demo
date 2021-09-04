package main

import (
	"context"
	"log"
	"net"

	pb "github.com/coacer/line-bot-demo/services/bot-server/src/botserver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5000"
)

type server struct{}

func (s *server) Test(ctx context.Context, in *pb.TestRequest) (*pb.TestReply, error) {
	return &pb.TestReply{Data: "hoge"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// サーバ起動
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterBotServerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
