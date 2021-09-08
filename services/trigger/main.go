package main

import (
	"log"
	"os"

	"trigger/grpc"
	"trigger/trigger"
)

func main() {
	t := trigger.New()
	s := grpc.NewRunner(os.Getenv("TRIGGER_PORT"), t)
	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
