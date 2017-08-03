package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	configFilename string
	config         Config
}

type Config struct {
	Port int
}

func (s *Server) ParseConfig() error {

	configFile, err := os.Open(s.Config)

	if err != nil {
		log.Fatalf("error opening config file")
		return err
	}

	jsonParser := json.NewDecoder(configFile)

	if err = jsonParser.Decode(&s.Config); err != nil {
		return err
	}

}

func (s *Server) Start() error {

	lis, err := net.Listen("tcp", s.Config.Port)

	if err != nil {
		return err
	}

	g := grpc.NewServer()
	pb.RegisterGreeterServer(g, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		return err
	}

}
