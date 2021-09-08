package main

import (
	"log"
	"os"

	"channel/channel"
	"channel/grpc"
)

func main() {
	c := channel.New()
	s := grpc.NewRunner(os.Getenv("CHANNEL_PORT"), c)
	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
