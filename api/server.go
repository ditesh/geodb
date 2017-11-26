package api

import (
	fmt "fmt"
	"geodb/geometry"
	"geodb/logger"
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

func (s *wrapper) WritePoint(ctx context.Context, in *WritePointRequest) (*Empty, error) {

	point := &geometry.Point{
		Lat:  in.P.Lat,
		Lng:  in.P.Lng,
		Elv:  in.P.Elv,
		Blob: in.Blob,
	}

	err := point.Write()

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

// Start starts the grpc server
func (s *Server) Start() error {

	lis, err := listen(s.Config.Port)

	if err != nil {
		return err
	}

	g := grpc.NewServer()
	RegisterAPIServer(g, &wrapper{})

	return g.Serve(lis)

}
