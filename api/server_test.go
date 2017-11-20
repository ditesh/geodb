package api

import (
	"geodb/config"
	"geodb/logger"
	"testing"

	context "golang.org/x/net/context"
)

func TestListen(t *testing.T) {

	logger.Configure(config.LoggerConfig{
		Type: "discard",
	})

	if lis, err := listen(-1); err == nil {
		lis.Close()
		t.Error("expected an error but didn't receive one")
	}

	if lis, err := listen(1); err == nil {
		lis.Close()
		t.Error("expected an error but didn't receive one")
	}

	// Random high port
	lis, err := listen(12311)

	if err != nil {
		t.Error("did not expect an error but received one")
	} else {
		lis.Close()
	}
}

func TestWrite(t *testing.T) {

	p := &Point{
		Lat: 0,
		Lng: 0,
		Elv: 0,
	}

	wr := &WriteRequest{
		P:    p,
		Blob: "test",
	}

	s := &wrapper{}

	if _, err := s.Write(context.Background(), wr); err != nil {
		t.Error("unable to write point")
	}
}
