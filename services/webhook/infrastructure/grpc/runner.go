package grpc

import (
	"fmt"
	"log"
	"net"

	webhook "webhook/infrastructure/grpc/linewebhook"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Runner struct {
	port   string
	server webhook.LineWebhookServer
}

func NewRunner(port string, server webhook.LineWebhookServer) *Runner {
	return &Runner{port, server}
}

func (r *Runner) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", r.port))
	if err != nil {
		return err
	}

	i := new(Interceptor)
	s := grpc.NewServer(grpc.UnaryInterceptor(i.logging()))
	reflection.Register(s)
	webhook.RegisterLineWebhookServer(s, r.server)
	log.Println("Starting gRPC server")
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
