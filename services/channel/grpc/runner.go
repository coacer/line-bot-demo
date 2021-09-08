package grpc

import (
	"fmt"
	"log"
	"net"

	"channel/channel"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Runner struct {
	port   string
	server channel.ChannelServer
}

func NewRunner(port string, server channel.ChannelServer) *Runner {
	return &Runner{port, server}
}

func (r *Runner) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", r.port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)
	channel.RegisterChannelServer(s, r.server)
	log.Println("Starting gRPC server")
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
