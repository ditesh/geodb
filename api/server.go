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

	p := structs.Point{
		Lat: in.P.Lat,
		Lng: in.P.Lng,
		Elv: in.P.Elv,
	}

	err := storage.WritePoint(p)

	if err != nil {
		return nil, err
	}

	return &Empty{}, nil

}

func listen(port int) (net.Listener, error) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return nil, err
	}

	logger.Info("listening on ", port)
	return lis, nil

}

func (s *Server) Start() error {

	lis, err := listen(s.Config.Port)

	if err != nil {
		return err
	}

	g := grpc.NewServer()
	RegisterAPIServer(g, &wrapper{})

	return g.Serve(lis)

}
