package main

import (
	"log"

	"trigger/grpc"
	"trigger/trigger"
)

const (
	port = ":5000"
)

func main() {
	t := trigger.New()
	s := grpc.NewRunner(port, t)
	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
