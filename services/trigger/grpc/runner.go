package grpc

import (
	"log"
	"net"

	"trigger/trigger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Runner struct {
	port   string
	server trigger.TriggerServer
}

func NewRunner(port string, server trigger.TriggerServer) *Runner {
	return &Runner{port, server}
}

func (r *Runner) Start() error {
	lis, err := net.Listen("tcp", r.port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)
	trigger.RegisterTriggerServer(s, r.server)
	log.Println("Starting gRPC server")
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
