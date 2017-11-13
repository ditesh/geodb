package api

import (
	fmt "fmt"
	"geodb/logger"
	"geodb/storage"
	"geodb/structs"
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

func (s *wrapper) Read(ctx context.Context, p *Point) (*WriteRequest, error) {
	return &WriteRequest{P: p, Blob: "blob"}, nil
}

func (s *wrapper) Write(ctx context.Context, in *WriteRequest) (*Empty, error) {

	p := &structs.Point{
		Lat: in.P.Lat,
		Lng: in.P.Lng,
		Elv: in.P.Elv,
	}

	storage.WritePoint(p)
	return &Empty{}, nil
}

func (s *Server) Start() error {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.Port))

	if err != nil {
		return err
	}

	logger.Info("listening on", s.Config.Port)

	g := grpc.NewServer()
	RegisterAPIServer(g, &wrapper{})

	err = g.Serve(lis)

	return err

}
