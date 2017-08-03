package api

import (
	"encoding/json"
	fmt "fmt"
	"log"
	"net"
	"os"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

func (s *wrapper) Read(ctx context.Context, p *Point) (*WriteRequest, error) {
	return &WriteRequest{Latlng: p, Blob: "blob"}, nil
}

func (s *wrapper) Write(ctx context.Context, in *WriteRequest) (*Empty, error) {
	return &Empty{}, nil
}

func (s *Server) ParseConfig() error {

	configFile, err := os.Open(s.ConfigFile)

	if err != nil {
		log.Fatalf("error opening config file")
		return err
	}

	jsonParser := json.NewDecoder(configFile)

	if err = jsonParser.Decode(&s.config); err != nil {
		return err
	}

	return nil

}

func (s *Server) Start() error {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))

	if err != nil {
		return err
	}

	g := grpc.NewServer()
	RegisterAPIServer(g, &wrapper{})

	if err := g.Serve(lis); err != nil {
		return err
	}

	return nil

}
