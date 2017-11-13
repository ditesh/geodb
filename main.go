package main

import (
	"fmt"
	"geodb/api"
	"geodb/config"
	"geodb/logger"
	"geodb/storage"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("ERROR: no config file provided")
		os.Exit(1)
	}

	c := &config.Config{}

	if err := c.Parse(os.Args[1]); err != nil {
		fmt.Println("ERROR: unable to parse config file:", err)
		os.Exit(1)
	}

	if err := logger.Configure(c.Logger); err != nil {
		fmt.Println("ERROR: unable to configure logger:", err)
		os.Exit(1)
	}

	if err := storage.Init(c.Datapath); err != nil {
		fmt.Println("ERROR: unable to access datadir:", err)
		os.Exit(1)
	}

	s := &api.Server{Config: c.APIServer}

	if err := s.Start(); err != nil {
		logger.Error("failed to listen:", err)
		os.Exit(1)
	}

}
