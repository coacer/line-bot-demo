package main

import (
	"log"
	"os"

	"webhook/infrastructure/grpc"
	webhook "webhook/infrastructure/grpc/linewebhook"
)

func main() {
	t := webhook.New()
	s := grpc.NewRunner(os.Getenv("WEBHOOK_PORT"), t)
	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
