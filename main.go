package main

import (
	"fmt"
	"geodb/server"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("no config file provided")
		os.Exit(1)
	}

	s := &server.Server{
		Config: os.Args[1],
	}

	if err := s.ParseConfig(); err != nil {
		log.Fatalf("unable to parse config file: %n", err)
		os.Exit(2)
	}

	if err := s.Start(); err != nil {
		log.Fatalf("failed to listen: %n", err)
		os.Exit(2)
	}

}
